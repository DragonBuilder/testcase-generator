package openai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamingChat(t *testing.T) {
	type args struct {
		request ChatRequest
		// chucks  chan StreamingChatResponseChunk
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sanity test",
			args: args{
				request: NewChatRequest([]Message{
					{
						Role:    User,
						Content: "Say this is a test!",
					},
				}),
				// chucks: make(chan StreamingChatResponseChunk),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := StreamingChat(tt.args.request)
			assert.NoError(t, err)
		})
	}
}
