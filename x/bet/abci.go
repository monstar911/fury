package bet

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fanfury-sports/fury/x/bet/keeper"
)

// EndBlocker settles the active bets of resolved sport events
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.BatchSportEventSettlements(ctx)
	if err != nil {
		panic(fmt.Sprintf("end block no %d failed : %s", ctx.BlockHeight(), err.Error()))
	}
}
