package types_test

import (
	"testing"

	"github.com/fanfury-sports/fury/x/sportevent/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				SportEventList: []types.SportEvent{
					{
						UID: "0",
					},
					{
						UID: "1",
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicated sportEvent",
			genState: &types.GenesisState{
				SportEventList: []types.SportEvent{
					{
						UID: "0",
					},
					{
						UID: "0",
					},
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
