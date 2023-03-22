package bet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fanfury-sports/fury/x/bet/keeper"
	"github.com/fanfury-sports/fury/x/bet/types"
)

// InitGenesis initializes the module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetBetStats(ctx, genState.Stats)

	// Set all the bet
	for _, bet := range genState.BetList {
		var id uint64
		for _, uid2ID := range genState.Uid2IdList {
			if uid2ID.UID == bet.UID {
				id = uid2ID.ID
			}
		}

		// Set all the active bet
		for i := range genState.ActiveBetList {
			active := genState.ActiveBetList[i]
			if genState.ActiveBetList[i].UID == bet.UID {
				k.SetActiveBet(ctx, &active, id, bet.SportEventUID)
			}
		}

		// Set all the settled bet
		for i := range genState.SettledBetList {
			settled := genState.SettledBetList[i]
			if settled.UID == bet.UID {
				k.SetSettledBet(ctx, &settled, id, bet.SettlementHeight)
			}
		}

		k.SetBet(ctx, bet, id)
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	var err error
	genesis.BetList, err = k.GetBets(ctx)
	if err != nil {
		panic(err)
	}

	genesis.ActiveBetList, err = k.GetActiveBets(ctx)
	if err != nil {
		panic(err)
	}

	genesis.SettledBetList, err = k.GetSettledBets(ctx)
	if err != nil {
		panic(err)
	}

	genesis.Uid2IdList, err = k.GetBetIDs(ctx)
	if err != nil {
		panic(err)
	}

	genesis.Stats = k.GetBetStats(ctx)

	return genesis
}
