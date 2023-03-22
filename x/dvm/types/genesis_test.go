package types_test

import (
	"testing"
	"time"

	"github.com/fanfury-sports/fury/x/dvm/types"
	"github.com/stretchr/testify/require"
)

const (
	testAddress = "cosmos1s4ycalgh3gjemd4hmqcvcgmnf647rnd0tpg2w9"
)

func TestGenesisState_Validate(t *testing.T) {
	pubKey1 := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+9wlxVu9a8lzUO2kcFLu\nUBIuV0+DpUdgEmsyQXr4y65sPSx/XjbK3GSZS1fB4irYPPG8EPHa6Z9KwWJLrTBr\nHayQcUBV5GQPf7nDktCkljYEBRmJZ+x3tlTf2kyKf3JMPAYgSFcs792dMpx8EiuE\n683QzUyeCutmiSWj1e7/IR9tjD4X/XFGkLES6wtqpQpOsL10z3hZllQEqZif8pDZ\nZcDvF97dg0l+JIWW3jBINL/UzuBRmdtDMuS1d57bpaMNb7L9HLUDBiwlZTGhs1+v\n9eTMY6IEdIzQ6M1KTFDeLYdnpGWP0ttBpt7SesLNpsKStbZ7QkbNtzlkTN8eJ6qu\nJQIDAQAB\n-----END PUBLIC KEY-----"
	vote := types.Vote{
		PublicKey: pubKey1,
		Vote:      types.ProposalVote_PROPOSAL_VOTE_YES,
	}
	validState := types.GenesisState{
		KeyVault: types.KeyVault{
			PublicKeys: []string{pubKey1},
		},
		ActivePubkeysChangeProposals: []types.PublicKeysChangeProposal{
			{
				Id:      1,
				Creator: testAddress,
				StartTS: time.Now().Unix(),
				Modifications: types.PubkeysChangeProposalPayload{
					Additions: []string{pubKey1},
					Deletions: []string{},
				},
				Votes: []*types.Vote{&vote},
			},
		},
		ProposalStats: types.ProposalStats{
			PubkeysChangeCount: 1,
		},
		Params: types.DefaultParams(),
	}

	notEqualCount := validState
	notEqualCount.ActivePubkeysChangeProposals = []types.PublicKeysChangeProposal{}

	duplicateProposal := validState
	duplicateProposal.FinishedPubkeysChangeProposals = []types.PublicKeysChangeFinishedProposal{
		{Proposal: duplicateProposal.ActivePubkeysChangeProposals[0]},
	}

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
			desc:     "not equal proposal count",
			genState: &notEqualCount,
			valid:    false,
		},
		{
			desc:     "duplicate proposal",
			genState: &duplicateProposal,
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
