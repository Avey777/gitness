// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package gitrpc

import (
	"context"
	"io"
)

type Interface interface {
	CreateRepository(ctx context.Context, params *CreateRepositoryParams) (*CreateRepositoryOutput, error)
	GetTreeNode(ctx context.Context, params *GetTreeNodeParams) (*GetTreeNodeOutput, error)
	ListTreeNodes(ctx context.Context, params *ListTreeNodeParams) (*ListTreeNodeOutput, error)
	GetSubmodule(ctx context.Context, params *GetSubmoduleParams) (*GetSubmoduleOutput, error)
	GetBlob(ctx context.Context, params *GetBlobParams) (*GetBlobOutput, error)
	CreateBranch(ctx context.Context, params *CreateBranchParams) (*CreateBranchOutput, error)
	GetBranch(ctx context.Context, params *GetBranchParams) (*GetBranchOutput, error)
	DeleteBranch(ctx context.Context, params *DeleteBranchParams) error
	ListBranches(ctx context.Context, params *ListBranchesParams) (*ListBranchesOutput, error)

	GetRef(ctx context.Context, params GetRefParams) (GetRefResponse, error)

	// UpdateRef creates, updates or deletes a git ref. If the OldValue is defined it must match the reference value
	// prior to the call. To remove a ref use the zero ref as the NewValue. To require the creation of a new one and
	// not update of an exiting one, set the zero ref as the OldValue.
	UpdateRef(ctx context.Context, params UpdateRefParams) error

	/*
	 * Commits service
	 */
	GetCommit(ctx context.Context, params *GetCommitParams) (*GetCommitOutput, error)
	ListCommits(ctx context.Context, params *ListCommitsParams) (*ListCommitsOutput, error)
	ListCommitTags(ctx context.Context, params *ListCommitTagsParams) (*ListCommitTagsOutput, error)
	GetCommitDivergences(ctx context.Context, params *GetCommitDivergencesParams) (*GetCommitDivergencesOutput, error)
	CommitFiles(ctx context.Context, params *CommitFilesParams) (CommitFilesResponse, error)

	/*
	 * Git Cli Service
	 */
	GetInfoRefs(ctx context.Context, w io.Writer, params *InfoRefsParams) error
	ServicePack(ctx context.Context, w io.Writer, params *ServicePackParams) error

	/*
	 * Diff services
	 */
	RawDiff(ctx context.Context, in *DiffParams, w io.Writer) error
	DiffShortStat(ctx context.Context, params *DiffParams) (DiffShortStatOutput, error)
	DiffStats(ctx context.Context, params *DiffParams) (DiffStatsOutput, error)

	/*
	 * Merge services
	 */
	Merge(ctx context.Context, in *MergeParams) (MergeOutput, error)

	/*
	 * Blame services
	 */
	Blame(ctx context.Context, params *BlameParams) (<-chan *BlamePart, <-chan error)
}
