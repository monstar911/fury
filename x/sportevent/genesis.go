package sportevent

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fanfury-sports/fury/x/sportevent/keeper"
	"github.com/fanfury-sports/fury/x/sportevent/types"
)

// InitGenesis initializes the module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the sport-events
	for _, elem := range genState.SportEventList {
		k.SetSportEvent(ctx, elem)
	}

	k.SetSportEventStats(ctx, genState.Stats)

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	var err error
	genesis.SportEventList, err = k.GetSportEventAll(ctx)
	if err != nil {
		panic(err)
	}

	genesis.Stats = k.GetSportEventStats(ctx)

	return genesis
}
