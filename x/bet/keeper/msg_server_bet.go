package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fanfury-sports/fury/x/bet/types"
)

func (k msgServer) PlaceBet(goCtx context.Context, msg *types.MsgPlaceBet) (*types.MsgPlaceBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetBetID(ctx, msg.Bet.UID)
	if isFound {
		return &types.MsgPlaceBetResponse{
				Error: types.ErrDuplicateUID.Error(),
				Bet:   msg.Bet,
			},
			types.ErrDuplicateUID
	}

	ticketData := &types.BetPlacementTicketPayload{}
	err := k.dvmKeeper.VerifyTicketUnmarshal(sdk.WrapSDKContext(ctx), msg.Bet.Ticket, &ticketData)
	if err != nil {
		return &types.MsgPlaceBetResponse{
				Error: sdkerrors.Wrapf(types.ErrInVerification, "%s", err).Error(),
				Bet:   msg.Bet,
			},
			sdkerrors.Wrapf(types.ErrInVerification, "%s", err)
	}

	if err = ticketData.Validate(msg.Creator); err != nil {
		return &types.MsgPlaceBetResponse{
				Error: err.Error(),
				Bet:   msg.Bet,
			},
			err
	}

	bet, err := types.NewBet(msg.Creator, msg.Bet, ticketData.OddsType, ticketData.SelectedOdds)
	if err != nil {
		return &types.MsgPlaceBetResponse{
				Error: err.Error(),
				Bet:   msg.Bet,
			},
			err
	}

	if err := k.Keeper.PlaceBet(ctx, bet); err != nil {
		return &types.MsgPlaceBetResponse{
				Error: err.Error(),
				Bet:   msg.Bet,
			},
			err
	}

	emitBetEvent(ctx, types.TypeMsgPlaceBet, msg.Bet.UID, msg.Creator)
	return &types.MsgPlaceBetResponse{
			Error: "",
			Bet:   msg.Bet,
		},
		nil
}

func emitBetEvent(ctx sdk.Context, msgType string, betUID string, betCreator string) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			msgType,
			sdk.NewAttribute(types.AttributeKeyBetCreator, betCreator),
			sdk.NewAttribute(types.AttributeKeyBetUID, betUID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, msgType),
			sdk.NewAttribute(sdk.AttributeKeySender, betCreator),
		),
	})
}
