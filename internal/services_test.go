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
			name: "c",
			args: args{
				featureExplation: "c",
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
