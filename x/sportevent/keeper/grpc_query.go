package keeper

import (
	"github.com/fanfury-sports/fury/x/sportevent/types"
)

var _ types.QueryServer = Keeper{}
