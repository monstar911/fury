package keeper

import (
	"github.com/fanfury-sports/fury/x/house/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the house MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}
