package openai

import (
	"reflect"
	"testing"
)

func TestChat(t *testing.T) {
	type args struct {
		request ChatRequest
	}
	tests := []struct {
		name    string
		args    args
		want    ChatResponse
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
			},
			want:    ChatResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Chat(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chat() = %v, want %v", got, tt.want)
			}
		})
	}
}
