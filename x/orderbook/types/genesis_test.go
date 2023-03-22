package types_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/fanfury-sports/fury/x/orderbook/types"
	"github.com/stretchr/testify/require"
)

const (
	testAddress = "cosmos1s4ycalgh3gjemd4hmqcvcgmnf647rnd0tpg2w9"
)

func TestGenesisState_Validate(t *testing.T) {
	sportEventUID := uuid.NewString()
	oddsUID := uuid.NewString()
	validState := types.GenesisState{
		BookList: []types.OrderBook{
			{
				ID:                 sportEventUID,
				ParticipationCount: 1,
				OddsCount:          1,
				Status:             types.OrderBookStatus_ORDER_BOOK_STATUS_STATUS_ACTIVE,
			},
		},
		BookExposureList: []types.BookOddsExposure{
			{
				BookUID:          sportEventUID,
				OddsUID:          oddsUID,
				FulfillmentQueue: []uint64{},
			},
		},
		BookParticipationList: []types.BookParticipation{
			{
				BookUID:            sportEventUID,
				Index:              1,
				ParticipantAddress: testAddress,
				IsModuleAccount:    true,
			},
		},
		ParticipationExposureList: []types.ParticipationExposure{
			{
				BookUID:            sportEventUID,
				OddsUID:            oddsUID,
				ParticipationIndex: 1,
			},
		},
		ParticipationExposureByIndexList: []types.ParticipationExposure{
			{
				BookUID:            sportEventUID,
				OddsUID:            oddsUID,
				ParticipationIndex: 1,
			},
		},
		HistoricalParticipationExposureList: []types.ParticipationExposure{
			{
				BookUID:            sportEventUID,
				OddsUID:            oddsUID,
				ParticipationIndex: 1,
			},
		},
		ParticipationBetPairExposureList: []types.ParticipationBetPair{
			{
				BookUID:            sportEventUID,
				ParticipationIndex: 1,
			},
		},
		PayoutLock: [][]byte{},
		Stats: types.OrderBookStats{
			ResolvedUnsettled: []string{sportEventUID},
		},
		Params: types.DefaultParams(),
	}

	invalidParticipantAddress := validState
	invalidParticipantAddress.BookParticipationList = []types.BookParticipation{validState.BookParticipationList[0]}
	invalidParticipantAddress.BookParticipationList[0].ParticipantAddress = "wrong"

	notEqualBookCount := validState
	notEqualBookCount.BookList = []types.OrderBook{}

	notEqualBookExposureCount := validState
	notEqualBookExposureCount.BookExposureList = []types.BookOddsExposure{}

	notEqualParticipationExposureCount := validState
	notEqualParticipationExposureCount.ParticipationExposureList = []types.ParticipationExposure{}

	notEqualParticipationExposureIndexCount := validState
	notEqualParticipationExposureIndexCount.ParticipationExposureByIndexList = []types.ParticipationExposure{}

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
			desc:     "valid genesis state",
			genState: &validState,
			valid:    true,
		},
		{
			desc:     "invalid participant address",
			genState: &invalidParticipantAddress,
			valid:    false,
		},
		{
			desc:     "not equal book",
			genState: &notEqualBookCount,
			valid:    false,
		},
		{
			desc:     "not equal book exposure",
			genState: &notEqualBookExposureCount,
			valid:    false,
		},
		{
			desc:     "not equal participation exposure",
			genState: &notEqualParticipationExposureCount,
			valid:    false,
		},
		{
			desc:     "not equal participation exposure index",
			genState: &notEqualParticipationExposureIndexCount,
			valid:    false,
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
