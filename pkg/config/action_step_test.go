package config_test

import (
	"testing"

	"github.com/jippi/scm-engine/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionStep_RequiredIntSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		step    config.ActionStep
		key     string
		want    []int
		wantErr string
	}{
		{
			name: "returns slice when key exists with valid value",
			step: config.ActionStep{
				"user_ids": []int{1, 2, 3},
			},
			key:     "user_ids",
			want:    []int{1, 2, 3},
			wantErr: "",
		},
		{
			name: "returns empty slice when key exists with empty slice",
			step: config.ActionStep{
				"user_ids": []int{},
			},
			key:     "user_ids",
			want:    []int{},
			wantErr: "",
		},
		{
			name:    "returns error when key is missing",
			step:    config.ActionStep{},
			key:     "user_ids",
			want:    nil,
			wantErr: "Required 'step' key 'user_ids' is missing",
		},
		{
			name: "returns error when value is wrong type (string)",
			step: config.ActionStep{
				"user_ids": "not a slice",
			},
			key:     "user_ids",
			want:    nil,
			wantErr: "Required 'step' key 'user_ids' must be of type []int, got string",
		},
		{
			name: "returns error when value is wrong type (int)",
			step: config.ActionStep{
				"user_ids": 123,
			},
			key:     "user_ids",
			want:    nil,
			wantErr: "Required 'step' key 'user_ids' must be of type []int, got int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.step.RequiredIntSlice(tt.key)

			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
