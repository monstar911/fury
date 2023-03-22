package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/fanfury-sports/fury/x/orderbook/types"
)

// SetBookOddsExposure sets a book odds exposure.
func (k Keeper) SetBookOddsExposure(ctx sdk.Context, boe types.BookOddsExposure) {
	bookKey := types.GetBookOddsExposureKey(boe.BookUID, boe.OddsUID)

	store := k.getBookOddsExposureStore(ctx)
	b := k.cdc.MustMarshal(&boe)
	store.Set(bookKey, b)
}

// GetBookOddsExposure returns a specific book odds exposure.
func (k Keeper) GetBookOddsExposure(ctx sdk.Context, bookUID, oddsUID string) (val types.BookOddsExposure, found bool) {
	sportEventsStore := k.getBookOddsExposureStore(ctx)
	exposureKey := types.GetBookOddsExposureKey(bookUID, oddsUID)
	b := sportEventsStore.Get(exposureKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// GetOddsExposuresByBook returns all exposures for a book
func (k Keeper) GetOddsExposuresByBook(ctx sdk.Context, bookUID string) (list []types.BookOddsExposure, err error) {
	store := k.getBookOddsExposureStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetBookOddsExposuresKey(bookUID))

	defer func() {
		err = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BookOddsExposure
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllBookExposures returns all book exposures used during genesis dump.
func (k Keeper) GetAllBookExposures(ctx sdk.Context) (list []types.BookOddsExposure, err error) {
	store := k.getBookOddsExposureStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() {
		err = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BookOddsExposure
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
