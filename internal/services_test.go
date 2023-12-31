package internal

import (
	"testcase-generator/internal/openai"
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

func TestStreamingGenerateTestcases(t *testing.T) {
	type args struct {
		featureExplation string
		stream           chan openai.StreamingChatResponseChunk
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sanity test",
			args: args{
				featureExplation: "A REST API to fetch a list of users.",
				stream:           make(chan openai.StreamingChatResponseChunk),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunks, err := StreamingGenerateTestcases(tt.args.featureExplation)
			assert.NoError(t, err)
			for chunk := range chunks {
				// t.Log(chunk.Choices[0].Delta.Content)
				assert.NotEmpty(t, chunk)
			}
		})
	}
}
