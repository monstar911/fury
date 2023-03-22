package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fanfury-sports/fury/consts"
	"github.com/fanfury-sports/fury/x/strategicreserve/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Params returns the Params for the strategic reserve
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest)
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
