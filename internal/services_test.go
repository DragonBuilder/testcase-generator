package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTestcases(t *testing.T) {
	type args struct {
		featureExplation string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sanity test",
			args: args{
				featureExplation: "A REST API to fetch a list of users.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, e := GenerateTestcases(tt.args.featureExplation)
			assert.NoError(t, e)
		})
	}
}
