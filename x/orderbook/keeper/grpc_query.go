package keeper

import (
	"github.com/fanfury-sports/fury/x/orderbook/types"
)

var _ types.QueryServer = Keeper{}
