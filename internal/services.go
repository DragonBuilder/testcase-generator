package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testcase-generator/internal/config"
	"testcase-generator/internal/openai"
)

func GenerateTestcases(featureExplation string) ([]string, error) {
	client := http.Client{}
	messages := []openai.Message{
		{
			Role:    "user",
			Content: "Say this is a test!",
		},
	}
	data := openai.ChatRequest{
		Model:       "gpt-3.5-turbo",
		Messages:    messages,
		Temperature: 0.7,
	}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json: %v", err)
	}
	postReq, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	log.Printf("OPENAI_KEY: %s", config.Config.OpenAI_API_Key)

	postReq.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", config.Config.OpenAI_API_Key)},
		"Content-Type":  {"application/json"},
	}
	resp, err := client.Do(postReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	log.Println(resp)
	// body := make([]byte, 0)
	// _, err = resp.Body.Read(body)
	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	log.Println(string(body))
	log.Println("---------")
	log.Println(resp.StatusCode)
	log.Println("---------")
	return nil, nil
}
