package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fanfury-sports/fury/utils"
	"github.com/fanfury-sports/fury/x/sportevent/types"
)

// SetSportEvent sets a specific sport-event in the store
func (k Keeper) SetSportEvent(ctx sdk.Context, sportEvent types.SportEvent) {
	store := k.getSportEventsStore(ctx)
	b := k.cdc.MustMarshal(&sportEvent)
	store.Set(utils.StrBytes(sportEvent.UID), b)
}

// GetSportEvent returns a specific sport-event by its UID
func (k Keeper) GetSportEvent(ctx sdk.Context, sportEventUID string) (val types.SportEvent, found bool) {
	sportEventsStore := k.getSportEventsStore(ctx)
	b := sportEventsStore.Get(utils.StrBytes(sportEventUID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// SportEventExists checks if a specific sport-event exists or not
func (k Keeper) SportEventExists(ctx sdk.Context, sportEventUID string) bool {
	_, found := k.GetSportEvent(ctx, sportEventUID)
	return found
}

// RemoveSportEvent removes a sport-event from the store
func (k Keeper) RemoveSportEvent(ctx sdk.Context, sportEventUID string) {
	store := k.getSportEventsStore(ctx)
	store.Delete(utils.StrBytes(sportEventUID))
}

// GetSportEventAll returns all sport-events
func (k Keeper) GetSportEventAll(ctx sdk.Context) (list []types.SportEvent, err error) {
	store := k.getSportEventsStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() {
		err = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SportEvent
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ResolveSportEvent updates a sport-event with its resolution
func (k Keeper) ResolveSportEvent(ctx sdk.Context, resolutionEvent *types.SportEventResolutionTicketPayload) (*types.SportEvent, error) {
	storedEvent, found := k.GetSportEvent(ctx, resolutionEvent.UID)
	if !found {
		return nil, types.ErrNoMatchingSportEvent
	}

	if storedEvent.Status != types.SportEventStatus_SPORT_EVENT_STATUS_ACTIVE &&
		storedEvent.Status != types.SportEventStatus_SPORT_EVENT_STATUS_INACTIVE {
		return nil, types.ErrCanNotBeAltered
	}

	storedEvent.ResolutionTS = resolutionEvent.ResolutionTS
	storedEvent.Status = resolutionEvent.Status

	// if the result is declared for the sport-event, we need to update the winner odds uids.
	if resolutionEvent.Status == types.SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED {
		storedEvent.WinnerOddsUIDs = resolutionEvent.WinnerOddsUIDs
	}

	// if the result is declared or the sport-event is canceled or aborted, it should be added
	// to the unsettled resolved sport-event list  in the state.
	if resolutionEvent.Status == types.SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED ||
		resolutionEvent.Status == types.SportEventStatus_SPORT_EVENT_STATUS_CANCELED ||
		resolutionEvent.Status == types.SportEventStatus_SPORT_EVENT_STATUS_ABORTED {
		// append sport-event id to the unsettled resolved in statistics.
		k.appendUnsettledResolvedSportEvent(ctx, storedEvent.UID)
	}

	k.SetSportEvent(ctx, storedEvent)

	return &storedEvent, nil
}

func emitTransactionEvent(ctx sdk.Context, emitType string, uid, bookUID, creator string) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			emitType,
			sdk.NewAttribute(types.AttributeKeySportEventsSuccessUID, uid),
			sdk.NewAttribute(types.AttributeKeyOrderBookUID, bookUID),
			sdk.NewAttribute(types.AttributeKeyEventsCreator, creator),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, emitType),
			sdk.NewAttribute(sdk.AttributeKeySender, creator),
		),
	})
}
