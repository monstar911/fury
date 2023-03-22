package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fanfury-sports/fury/x/sportevent/types"
)

// AddSportEvent accepts ticket containing multiple creation events and return batch response after processing
func (k msgServer) AddSportEvent(goCtx context.Context, msg *types.MsgAddSportEvent) (*types.MsgAddSportEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var addPayload types.SportEventAddTicketPayload
	if err := k.dvmKeeper.VerifyTicketUnmarshal(goCtx, msg.Ticket, &addPayload); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInVerification, "%s", err)
	}

	params := k.GetParams(ctx)

	if err := addPayload.Validate(ctx, &params); err != nil {
		return nil, sdkerrors.Wrap(err, "validate add event")
	}

	_, found := k.Keeper.GetSportEvent(ctx, addPayload.UID)
	if found {
		return nil, types.ErrEventAlreadyExist
	}

	var oddsUIDs []string
	for _, odds := range addPayload.Odds {
		oddsUIDs = append(oddsUIDs, odds.UID)
	}
	err := k.orderBookKeeper.InitiateBook(ctx, addPayload.UID, addPayload.SrContributionForHouse, oddsUIDs)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInOrderBookInitiation, "%s", err)
	}

	sportEvent := types.NewSportEvent(
		addPayload.UID,
		msg.Creator,
		addPayload.StartTS,
		addPayload.EndTS,
		addPayload.Odds,
		params.NewEventBetConstraints(addPayload.MinBetAmount, addPayload.BetFee),
		addPayload.Meta,
		addPayload.UID,
		addPayload.SrContributionForHouse,
		addPayload.Status,
	)

	k.Keeper.SetSportEvent(ctx, sportEvent)

	response := &types.MsgAddSportEventResponse{
		Error: "",
		Data:  &sportEvent,
	}
	emitTransactionEvent(ctx, types.TypeMsgCreateSportEvents, response.Data.UID, addPayload.UID, msg.Creator)

	return response, nil
}
