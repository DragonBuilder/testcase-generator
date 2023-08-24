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
			// if err := StreamingChat(tt.args.request, tt.args.stream); (err != nil) != tt.wantErr {
			// 	t.Errorf("StreamingChat() error = %v, wantErr %v", err, tt.wantErr)
			// }
			chunks, err := StreamingChat(tt.args.request)
			assert.NoError(t, err)
			for chunk := range chunks {
				for _, c := range chunk {
					// log.Println(c.Choices[0].Delta.Content)
					assert.NotNil(t, c)
				}
			}
		})
	}
}
