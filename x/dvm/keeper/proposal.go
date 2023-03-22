package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fanfury-sports/fury/utils"
	"github.com/fanfury-sports/fury/x/dvm/types"
)

// SetActivePubkeysChangeProposal sets a pubkey list change proposal in the store.
func (k Keeper) SetActivePubkeysChangeProposal(ctx sdk.Context, proposal types.PublicKeysChangeProposal) {
	store := k.getActivePubKeysChangeProposalStore(ctx)
	b := k.cdc.MustMarshal(&proposal)
	store.Set(utils.Uint64ToBytes(proposal.Id), b)
}

// GetActivePubkeysChangeProposal returns a pubkeys change proposat by its id
func (k Keeper) GetActivePubkeysChangeProposal(ctx sdk.Context, id uint64) (val types.PublicKeysChangeProposal, found bool) {
	store := k.getActivePubKeysChangeProposalStore(ctx)

	b := store.Get(utils.Uint64ToBytes(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllActivePubkeysChangeProposals returns list of all active pubkeys change proposals
func (k Keeper) GetAllActivePubkeysChangeProposals(ctx sdk.Context) (list []types.PublicKeysChangeProposal, err error) {
	store := k.getActivePubKeysChangeProposalStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() {
		err = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PublicKeysChangeProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveActiveProposal removes an active pubkeys change proposal.
func (k Keeper) RemoveActiveProposal(ctx sdk.Context, id uint64) {
	store := k.getActivePubKeysChangeProposalStore(ctx)
	store.Delete(utils.Uint64ToBytes(id))
}

// SetFinishedPubkeysChangeProposal sets a finished pubkey list change proposal in the store.
func (k Keeper) SetFinishedPubkeysChangeProposal(ctx sdk.Context, finishedProposal types.PublicKeysChangeFinishedProposal) {
	store := k.getFinishedPubKeysChangeProposalStore(ctx)
	b := k.cdc.MustMarshal(&finishedProposal)
	store.Set(utils.Uint64ToBytes(finishedProposal.Proposal.Id), b)
}

// GetFinishedPubkeysChangeProposal returns a finished pubkeys change proposat by its id
func (k Keeper) GetFinishedPubkeysChangeProposal(ctx sdk.Context, id uint64) (val types.PublicKeysChangeFinishedProposal, found bool) {
	store := k.getFinishedPubKeysChangeProposalStore(ctx)

	b := store.Get(utils.Uint64ToBytes(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllFinishedPubkeysChangeProposals returns list of all finished pubkeys change proposals
func (k Keeper) GetAllFinishedPubkeysChangeProposals(ctx sdk.Context) (list []types.PublicKeysChangeFinishedProposal, err error) {
	store := k.getFinishedPubKeysChangeProposalStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() {
		err = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PublicKeysChangeFinishedProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) FinishProposals(ctx sdk.Context) error {
	return k.finishPubkeysChangeProposals(ctx)
}

func (k Keeper) finishPubkeysChangeProposals(ctx sdk.Context) error {
	activeProposals, err := k.GetAllActivePubkeysChangeProposals(ctx)
	if err != nil {
		return fmt.Errorf("error while retrieving the active pubkeys change proposals: %s", err)
	}

	if len(activeProposals) == 0 {
		return nil
	}

	blockTime := ctx.BlockTime().Unix()

	for _, proposal := range activeProposals {
		if proposal.IsExpired(blockTime) {
			// proposal is expired
			if err = k.finishPubkeysChangeProposal(
				ctx,
				proposal.Id,
				types.ProposalResult_PROPOSAL_RESULT_EXPIRED,
				fmt.Sprintf("current block time is %d, more than %d minutes is passed since start time.", blockTime, types.MaxValidProposalMinutes),
			); err != nil {
				return fmt.Errorf("error while setting the proposal as expired: %s", err)
			}

			// this proposal is expired go for next proposal
			continue
		}

		var approvedCount, rejectedCount int
		for _, v := range proposal.Votes {
			switch v.Vote {
			case types.ProposalVote_PROPOSAL_VOTE_YES:
				approvedCount++
			case types.ProposalVote_PROPOSAL_VOTE_NO:
				rejectedCount++
			}
		}

		if rejectedCount > types.MinVoteCountForDecision {
			if err = k.finishPubkeysChangeProposal(
				ctx,
				proposal.Id,
				types.ProposalResult_PROPOSAL_RESULT_REJECTED,
				fmt.Sprintf("rejected with %d number of 'no' votes.", rejectedCount),
			); err != nil {
				return fmt.Errorf("error while setting the proposal as rejected: %s", err)
			}

			// this proposal is rejected go for next proposal
			continue
		}

		if approvedCount > types.MinVoteCountForDecision {
			keyVault, found := k.GetKeyVault(ctx)
			if !found {
				fmt.Printf("there is no publick keys record")
			}

			for _, deleted := range proposal.Modifications.Deletions {
				keyVault.PublicKeys = utils.RemoveStr(keyVault.PublicKeys, deleted)
			}

			keyVault.PublicKeys = append(keyVault.PublicKeys, proposal.Modifications.Additions...)

			if err = k.finishPubkeysChangeProposal(
				ctx,
				proposal.Id,
				types.ProposalResult_PROPOSAL_RESULT_APPROVED,
				"",
			); err != nil {
				return fmt.Errorf("error while setting the proposal as approved: %s", err)
			}

			k.SetKeyVault(ctx, keyVault)
		}
	}

	return nil
}

func (k Keeper) finishPubkeysChangeProposal(
	ctx sdk.Context,
	proposalID uint64,
	result types.ProposalResult,
	resultMetadata string,
) error {
	proposal, found := k.GetActivePubkeysChangeProposal(ctx, proposalID)
	if !found {
		return fmt.Errorf("proposal not found with id %d", proposalID)
	}

	k.SetFinishedPubkeysChangeProposal(ctx,
		types.NewFinishedPublicKeysChangeProposal(
			proposal,
			result,
			resultMetadata,
			ctx.BlockTime().Unix(),
		),
	)
	k.RemoveActiveProposal(ctx, proposalID)

	return nil
}
