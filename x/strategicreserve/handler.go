package strategicreserve

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sge-network/sge/x/strategicreserve/keeper"
	"github.com/sge-network/sge/x/strategicreserve/types"
)

// NewHandler returns sdk.handler instance with configured message handler function
func NewHandler(k keeper.Keeper) sdk.Handler {
	keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)

		}
	}
}
