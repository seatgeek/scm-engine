package gitlab_test

import (
	"testing"

	"github.com/jippi/scm-engine/pkg/scm"
	"github.com/jippi/scm-engine/pkg/scm/gitlab"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCodeOwners(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		mergeRequest   *gitlab.ContextMergeRequest
		expectedOwners scm.Actors
	}{
		{
			name: "No approval rules",
			mergeRequest: &gitlab.ContextMergeRequest{
				ApprovalState: &gitlab.ContextApprovalState{
					Rules: []gitlab.ContextApprovalRule{},
				},
			},
			expectedOwners: scm.Actors{},
		},
		{
			name: "Approval rules without code owners",
			mergeRequest: &gitlab.ContextMergeRequest{
				ApprovalState: &gitlab.ContextApprovalState{
					Rules: []gitlab.ContextApprovalRule{
						{
							Type: scm.Ptr(gitlab.ApprovalRuleTypeAnyApprover),
							EligibleApprovers: []gitlab.ContextUser{
								{Username: "user1"},
							},
						},
					},
				},
			},
			expectedOwners: scm.Actors{},
		},
		{
			name: "Approval rules with code owners",
			mergeRequest: &gitlab.ContextMergeRequest{
				ApprovalState: &gitlab.ContextApprovalState{
					Rules: []gitlab.ContextApprovalRule{
						{
							Type: scm.Ptr(gitlab.ApprovalRuleTypeCodeOwner),
							EligibleApprovers: []gitlab.ContextUser{
								{Username: "user1"},
								{Username: "user2", Bot: true}, // Should be ignored
								{Username: "user3"},
							},
						},
					},
				},
			},
			expectedOwners: scm.Actors{
				{Username: "user1"},
				{Username: "user3"},
			},
		},
		{
			name: "Duplicate code owners",
			mergeRequest: &gitlab.ContextMergeRequest{
				ApprovalState: &gitlab.ContextApprovalState{
					Rules: []gitlab.ContextApprovalRule{
						{
							Type: scm.Ptr(gitlab.ApprovalRuleTypeCodeOwner),
							EligibleApprovers: []gitlab.ContextUser{
								{Username: "user1"},
								{Username: "user1"}, // Duplicate, should be ignored
							},
						},
					},
				},
			},
			expectedOwners: scm.Actors{
				{Username: "user1"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := &gitlab.Context{
				MergeRequest: tt.mergeRequest,
			}
			owners := ctx.GetCodeOwners()
			assert.Equal(t, tt.expectedOwners, owners)
		})
	}
}

func TestTotalLinesAdded(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mr       *gitlab.ContextMergeRequest
		expected int
	}{
		{
			name:     "Empty diff stats",
			mr:       &gitlab.ContextMergeRequest{},
			expected: 0,
		},
		{
			name: "Single file with additions",
			mr: &gitlab.ContextMergeRequest{
				DiffStats: []gitlab.ContextDiffStat{
					{Path: "file1.go", Additions: 10, Deletions: 5},
				},
			},
			expected: 10,
		},
		{
			name: "Multiple files",
			mr: &gitlab.ContextMergeRequest{
				DiffStats: []gitlab.ContextDiffStat{
					{Path: "file1.go", Additions: 10, Deletions: 5},
					{Path: "file2.go", Additions: 20, Deletions: 3},
					{Path: "file3.go", Additions: 5, Deletions: 0},
				},
			},
			expected: 35,
		},
		{
			name: "Files with zero additions",
			mr: &gitlab.ContextMergeRequest{
				DiffStats: []gitlab.ContextDiffStat{
					{Path: "file1.go", Additions: 0, Deletions: 10},
					{Path: "file2.go", Additions: 0, Deletions: 5},
				},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.mr.TotalLinesAdded()
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestTotalLinesRemoved(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mr       *gitlab.ContextMergeRequest
		expected int
	}{
		{
			name:     "Empty diff stats",
			mr:       &gitlab.ContextMergeRequest{},
			expected: 0,
		},
		{
			name: "Single file with deletions",
			mr: &gitlab.ContextMergeRequest{
				DiffStats: []gitlab.ContextDiffStat{
					{Path: "file1.go", Additions: 10, Deletions: 5},
				},
			},
			expected: 5,
		},
		{
			name: "Multiple files",
			mr: &gitlab.ContextMergeRequest{
				DiffStats: []gitlab.ContextDiffStat{
					{Path: "file1.go", Additions: 10, Deletions: 5},
					{Path: "file2.go", Additions: 20, Deletions: 3},
					{Path: "file3.go", Additions: 5, Deletions: 12},
				},
			},
			expected: 20,
		},
		{
			name: "Files with zero deletions",
			mr: &gitlab.ContextMergeRequest{
				DiffStats: []gitlab.ContextDiffStat{
					{Path: "file1.go", Additions: 10, Deletions: 0},
					{Path: "file2.go", Additions: 5, Deletions: 0},
				},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.mr.TotalLinesRemoved()
			require.Equal(t, tt.expected, result)
		})
	}
}

