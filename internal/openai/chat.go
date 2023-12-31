package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testcase-generator/internal/config"
)

const (
	System    = "system"
	User      = "user"
	Assistant = "assistant"

	chatCompletionsAPI = "https://api.openai.com/v1/chat/completions"
	temperature        = 0.0
	gptModel           = "gpt-3.5-turbo"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	Stream      bool      `json:"stream"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// {"id":"chatcmpl-7r1I5PziAUG70qR8Fw5q00KqEpN8X",
// "object":"chat.completion.chunk",
// "created":1692870701,
// "model":"gpt-3.5-turbo-0613",
// "choices":[{"index":0,"delta":{},"finish_reason":"stop"}]
// }

func NewChatRequest(messages []Message, stream bool) ChatRequest {
	return ChatRequest{
		Model:       gptModel,
		Messages:    messages,
		Temperature: temperature,
		Stream:      stream,
	}
}

func Chat(messages []Message) (ChatResponse, error) {
	request := NewChatRequest(messages, false)
	client := http.Client{}
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error marshaling request json: %v", err)
	}
	httpReq, err := http.NewRequest("POST", chatCompletionsAPI, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error creating request: %v", err)
	}
	// log.Printf("OPENAI_KEY: %s", config.Config.OpenAI_API_Key)

	httpReq.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", config.Config.OpenAI_API_Key)},
		"Content-Type":  {"application/json"},
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	// log.Println(resp)

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	response := ChatResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error unmarshaling json: %v", err)
	}
	return response, nil
}
