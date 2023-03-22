package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	simappUtil "github.com/fanfury-sports/fury/testutil/simapp"
	"github.com/fanfury-sports/fury/x/bet/types"
	sporteventtypes "github.com/fanfury-sports/fury/x/sportevent/types"
	"github.com/stretchr/testify/require"
)

func TestSettleBet(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)
	testCreator = simappUtil.TestParamUsers["user1"].Address.String()
	addTestSportEvent(t, tApp, ctx)

	tcs := []struct {
		desc             string
		betUID           string
		updateSportEvent *sporteventtypes.SportEvent
		bet              *types.Bet
		err              error
	}{
		{
			desc:   "invalid betUID",
			betUID: "invalidUID",
			err:    types.ErrInvalidBetUID,
		},
		{
			desc:   "not found bet",
			betUID: "0db09053-2901-4110-8fb5-c14e21f8d400",
			err:    types.ErrNoMatchingBet,
		},
		{
			desc: "invalid creator",
			bet: &types.Bet{
				UID:           "0db09053-2901-4110-8fb5-c14e21f8d400",
				SportEventUID: "nonExistSportEvent",
				Creator:       "invalid creator",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "failed in checking status",
			bet: &types.Bet{
				SportEventUID: testSportEventUID,
				OddsValue:     "10",
				OddsType:      types.OddsType_ODDS_TYPE_DECIMAL,
				Amount:        sdk.NewInt(500),
				Creator:       testCreator,
				OddsUID:       testOddsUID1,
				Status:        types.Bet_STATUS_SETTLED,
			},
			err: types.ErrBetIsSettled,
		},
		{
			desc: "not found sport-event",
			bet: &types.Bet{
				UID:           "0db09053-2901-4110-8fb5-c14e21f8d400",
				SportEventUID: "nonExistSportEvent",
				Creator:       testCreator,
			},
			err: types.ErrNoMatchingSportEvent,
		},
		{
			desc: "sport-event is aborted",
			bet: &types.Bet{
				SportEventUID: testSportEventUID,
				OddsValue:     "10",
				OddsType:      types.OddsType_ODDS_TYPE_DECIMAL,
				Amount:        sdk.NewInt(500),
				Creator:       testCreator,
				OddsUID:       testOddsUID1,
			},
			updateSportEvent: &sporteventtypes.SportEvent{
				UID:                    testSportEventUID,
				Creator:                testCreator,
				StartTS:                1111111111,
				EndTS:                  uint64(ctx.BlockTime().Unix()) + 1000,
				Odds:                   testEventOdds,
				SrContributionForHouse: sdk.NewInt(500000),

				Status: sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_ABORTED,
			},
		},
		{
			desc: "sport-event is canceled",
			bet: &types.Bet{
				SportEventUID: testSportEventUID,
				OddsValue:     "10",
				OddsType:      types.OddsType_ODDS_TYPE_DECIMAL,
				Amount:        sdk.NewInt(300),
				Creator:       testCreator,
				OddsUID:       testOddsUID1,
			},
			updateSportEvent: &sporteventtypes.SportEvent{
				UID:                    testSportEventUID,
				Creator:                testCreator,
				StartTS:                1111111111,
				EndTS:                  uint64(ctx.BlockTime().Unix()) + 1000,
				Odds:                   testEventOdds,
				SrContributionForHouse: sdk.NewInt(500000),

				Status: sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_CANCELED,
			},
		},
		{
			desc: "result is not declared",
			bet: &types.Bet{
				SportEventUID: testSportEventUID,
				OddsValue:     "10",
				OddsType:      types.OddsType_ODDS_TYPE_DECIMAL,
				Amount:        sdk.NewInt(500),
				Creator:       testCreator,
				OddsUID:       testOddsUID1,
			},
			updateSportEvent: &sporteventtypes.SportEvent{
				UID:                    testSportEventUID,
				Creator:                testCreator,
				StartTS:                1111111111,
				EndTS:                  uint64(ctx.BlockTime().Unix()) + 1000,
				Odds:                   testEventOdds,
				SrContributionForHouse: sdk.NewInt(500000),

				Status: sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_ACTIVE,
			},
			err: types.ErrResultNotDeclared,
		},
		{
			desc: "success",
			bet: &types.Bet{
				SportEventUID: testSportEventUID,
				OddsValue:     "10",
				OddsType:      types.OddsType_ODDS_TYPE_DECIMAL,
				Amount:        sdk.NewInt(500),
				Creator:       testCreator,
				OddsUID:       testOddsUID1,

				Result: types.Bet_RESULT_WON,
			},
			updateSportEvent: &sporteventtypes.SportEvent{
				UID:                    testSportEventUID,
				Creator:                testCreator,
				StartTS:                1111111111,
				EndTS:                  uint64(ctx.BlockTime().Unix()) + 1000,
				Odds:                   testEventOdds,
				SrContributionForHouse: sdk.NewInt(500000),

				Status: sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			betUID := uuid.NewString()

			if tc.bet != nil {
				resetSportEvent := sporteventtypes.SportEvent{
					UID:     testSportEventUID,
					Creator: testCreator,
					StartTS: 1111111111,
					EndTS:   uint64(ctx.BlockTime().Unix()) + 1000,
					Odds:    testEventOdds,
					Status:  sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_ACTIVE,
					BetConstraints: &sporteventtypes.EventBetConstraints{
						MinAmount: sdk.NewInt(1),
						BetFee:    sdk.NewInt(1),
					},
					SrContributionForHouse: sdk.NewInt(500000),
				}
				tApp.SportEventKeeper.SetSportEvent(ctx, resetSportEvent)
				tc.bet.UID = betUID
				placeTestBet(ctx, t, tApp, betUID, nil)
				k.SetBet(ctx, *tc.bet, 1)
			}

			if tc.updateSportEvent != nil {
				tApp.SportEventKeeper.SetSportEvent(ctx, *tc.updateSportEvent)
			}

			if tc.betUID != "" {
				betUID = tc.betUID
			}

			if tc.bet == nil {
				tc.bet = &types.Bet{
					Creator: "",
				}
			}

			err := k.SettleBet(ctx, tc.bet.Creator, betUID)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestBatchSettleBet(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)

	p := k.GetParams(ctx)
	p.BatchSettlementCount = 7
	k.SetParams(ctx, p)

	sportEventCount := 5
	sportEventBetCount := 10
	allBetCount := sportEventCount * sportEventBetCount
	blockCount := allBetCount/int(p.BatchSettlementCount) + 1

	sportEventUIDs := addTestSportEventBatch(t, tApp, ctx, sportEventCount)
	for _, sportEventUID := range sportEventUIDs {
		sportEvent, found := tApp.SportEventKeeper.GetSportEvent(ctx, sportEventUID)
		require.True(t, found)

		sportEvent.Status = sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_ACTIVE
		sportEvent.BetConstraints = &sporteventtypes.EventBetConstraints{
			MinAmount: sdk.NewInt(1),
			BetFee:    sdk.NewInt(1),
		}
		tApp.SportEventKeeper.SetSportEvent(ctx, sportEvent)

		for i := 0; i < sportEventBetCount; i++ {
			placeTestBet(ctx, t, tApp,
				uuid.NewString(),
				&types.BetOdds{
					UID:               testOddsUID1,
					SportEventUID:     sportEventUID,
					Value:             "4.20",
					MaxLossMultiplier: sdk.MustNewDecFromStr("0.1"),
				},
			)
		}
	}

	allActiveBets, err := k.GetActiveBets(ctx)
	require.NoError(t, err)
	require.Equal(t, allBetCount, len(allActiveBets))

	for _, sportEventUID := range sportEventUIDs[:len(sportEventUIDs)-2] {
		_, err := tApp.SportEventKeeper.ResolveSportEvent(ctx, &sporteventtypes.SportEventResolutionTicketPayload{
			UID:            sportEventUID,
			ResolutionTS:   uint64(ctx.BlockTime().Unix()) + 10000,
			WinnerOddsUIDs: []string{testOddsUID1, testOddsUID2, testOddsUID3},
			Status:         sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED,
		})
		require.NoError(t, err)
	}

	_, err = tApp.SportEventKeeper.ResolveSportEvent(ctx, &sporteventtypes.SportEventResolutionTicketPayload{
		UID:          sportEventUIDs[len(sportEventUIDs)-2],
		ResolutionTS: uint64(ctx.BlockTime().Unix()) + 10000,
		Status:       sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_CANCELED,
	})
	require.NoError(t, err)

	_, err = tApp.SportEventKeeper.ResolveSportEvent(ctx, &sporteventtypes.SportEventResolutionTicketPayload{
		UID:          sportEventUIDs[len(sportEventUIDs)-1],
		ResolutionTS: uint64(ctx.BlockTime().Unix()) + 10000,
		Status:       sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_ABORTED,
	})
	require.NoError(t, err)

	for i := 1; i <= blockCount; i++ {
		ctx = ctx.WithBlockHeight(int64(i))
		err := k.BatchSportEventSettlements(ctx)
		require.NoError(t, err)

		activeBets, err := k.GetActiveBets(ctx)
		require.NoError(t, err)

		settledBets, err := k.GetSettledBets(ctx)
		require.NoError(t, err)

		sportEventStats := tApp.SportEventKeeper.GetSportEventStats(ctx)

		t.Logf("block: %d, active bets: %d, settled bets: %d, resolved events: %v\n", i, len(activeBets), len(settledBets), sportEventStats.ResolvedUnsettled)
		require.GreaterOrEqual(t, int(p.BatchSettlementCount)*i, len(settledBets))
	}

	allActiveBets, err = k.GetActiveBets(ctx)
	require.NoError(t, err)
	require.Equal(t, 0, len(allActiveBets))

	allSettledBets, err := k.GetSettledBets(ctx)
	require.NoError(t, err)
	require.Equal(t, allBetCount, len(allSettledBets))

	allBets, err := k.GetBets(ctx)
	require.NoError(t, err)
	for _, bet := range allBets {
		require.NotEqual(t, 0, bet.SettlementHeight)
	}
}

func TestCheckBetStatus(t *testing.T) {
	k, _ := setupKeeper(t)
	tcs := []struct {
		desc string
		bet  *types.Bet
		err  error
	}{
		{
			desc: "pending bet",
			bet: &types.Bet{
				Status: types.Bet_STATUS_PENDING,
			},
		},
		{
			desc: "canceled bet",
			bet: &types.Bet{
				Status: types.Bet_STATUS_CANCELED,
			},
			err: types.ErrBetIsCanceled,
		},
		{
			desc: "settled bet",
			bet: &types.Bet{
				Status: types.Bet_STATUS_SETTLED,
			},
			err: types.ErrBetIsSettled,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			err := k.CheckBetStatus(tc.bet)
			if tc.err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestProcessBetResultAndStatus(t *testing.T) {
	k, _ := setupKeeper(t)
	tcs := []struct {
		desc       string
		bet        *types.Bet
		sportEvent sporteventtypes.SportEvent
		err        error
		result     types.Bet_Result
	}{
		{
			desc: "not declared",
			sportEvent: sporteventtypes.SportEvent{
				Status: sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED,
			},
			bet: &types.Bet{},
			err: types.ErrResultNotDeclared,
		},
		{
			desc: "won",
			sportEvent: sporteventtypes.SportEvent{
				Status:         sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED,
				WinnerOddsUIDs: []string{"oddsUID"},
			},
			bet: &types.Bet{
				OddsUID: "oddsUID",
			},
			result: types.Bet_RESULT_WON,
		},
		{
			desc: "lost",
			sportEvent: sporteventtypes.SportEvent{
				Status:         sporteventtypes.SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED,
				WinnerOddsUIDs: []string{"oddsUID"},
			},
			bet:    &types.Bet{},
			result: types.Bet_RESULT_LOST,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			err := k.ProcessBetResultAndStatus(tc.bet, tc.sportEvent)
			if tc.err != nil {
				require.Equal(t, tc.err, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tc.bet.Result, tc.result)
				require.Equal(t, tc.bet.Status, types.Bet_STATUS_RESULT_DECLARED)
			}
		})
	}
}
