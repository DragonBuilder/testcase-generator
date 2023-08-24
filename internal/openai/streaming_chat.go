package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testcase-generator/internal/config"
)

type StreamingChatDelta struct {
	Content string `json:"content"`
}

type StreamingChoice struct {
	Index        int                `json:"index"`
	Message      Message            `json:"message"`
	Delta        StreamingChatDelta `json:"delta"`
	FinishReason string             `json:"finish_reason"`
}

type StreamingChatResponse struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created int64             `json:"created"`
	Model   string            `json:"model"`
	Choices []StreamingChoice `json:"choices"`
}

func (r StreamingChatResponse) done() bool {
	for _, c := range r.Choices {
		if c.FinishReason == "stop" {
			return true
		}
	}
	return false
}

type StreamingChatResponseChunk []StreamingChatResponse

func (chunk StreamingChatResponseChunk) done() bool {
	for _, c := range chunk {
		if c.done() {
			return true
		}
	}
	return false
}

// TODO: give select, to receive error
func StreamingChat(request ChatRequest, chunks chan<- StreamingChatResponseChunk) error {
	client := http.Client{}
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling request json: %v", err)
	}
	httpReq, err := http.NewRequest("POST", chatCompletionsAPI, bytes.NewBuffer(jsonStr))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", config.Config.OpenAI_API_Key)},
		"Content-Type":  {"application/json"},
		"Accept":        {"text/event-stream"},
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	for {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		resultChunk := parseChunk(body)

		chunks <- resultChunk
		if resultChunk.done() {
			close(chunks)
			return nil
		}
	}
}

// TODO: give select, to receive error
func parseChunk(body []byte) StreamingChatResponseChunk {
	asStrArr := strings.Split(string(body), "data: ")
	for i := range asStrArr {
		asStrArr[i] = strings.TrimSpace(asStrArr[i])
	}
	asStrArr = asStrArr[:len(asStrArr)-1]

	asStr := strings.Join(asStrArr, ",")
	asStr = strings.TrimLeft(asStr, ",")
	asStr = fmt.Sprintf("[%s]", asStr)

	result := StreamingChatResponseChunk{}

	err := json.Unmarshal([]byte(asStr), &result)
	if err != nil {
		log.Fatalf("Error : %v\n", err)
		// return err
	}
	return result
}

// func parseChunkIntoJsonString() string {
// 	chuckReceivedStrArr := strings.Split(string(body), "data: ")
// 	for i := range chuckReceivedStrArr {
// 		chuckReceivedStrArr[i] = strings.TrimSpace(chuckReceivedStrArr[i])
// 	}

// 	chuckReceivedStrArr = chuckReceivedStrArr[:len(chuckReceivedStrArr)-1]

// 	chuckReceivedStr := strings.Join(chuckReceivedStrArr, ",")
// 	chuckReceivedStr = strings.TrimLeft(chuckReceivedStr, ",")
// 	chuckReceivedStr = fmt.Sprintf("[%s]", chuckReceivedStr)

// }
