package openai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamingChat(t *testing.T) {
	type args struct {
		messages []Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sanity test",
			args: args{
				messages: []Message{
					{
						Role:    User,
						Content: "Say this is a test!",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream, err := StreamingChat(tt.args.messages)
			assert.NoError(t, err)
			for c := range stream {
				assert.NotEmpty(t, c.Choices)
			}
		})
	}
}
