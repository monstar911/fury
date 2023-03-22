package simulation

import (
	//#nosec
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/fanfury-sports/fury/x/sportevent/keeper"
	"github.com/fanfury-sports/fury/x/sportevent/types"
)

// SimulateMsgAddEvent simulates the add event flow
func SimulateMsgAddEvent(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddSportEvent{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddSportEvent simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddSportEvent simulation not implemented"), nil, nil
	}
}
