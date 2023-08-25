package openai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
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
	Delta        StreamingChatDelta `json:"delta"`
	FinishReason string             `json:"finish_reason"`
}

type StreamingChatResponseChunk struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created int64             `json:"created"`
	Model   string            `json:"model"`
	Choices []StreamingChoice `json:"choices"`
}

func (r StreamingChatResponseChunk) isLast() bool {
	for _, c := range r.Choices {
		if c.FinishReason == "stop" {
			return true
		}
	}
	return false
}

// TODO: give select, to receive error
func StreamingChat(messages []Message) (<-chan StreamingChatResponseChunk, error) {
	request := NewChatRequest(messages, true)
	client := http.Client{}
	reqJSON, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request json: %v", err)
	}
	reqBody := bytes.NewBuffer(reqJSON)
	// log.Printf("Request body: %s", reqBody.String())

	httpReq, err := http.NewRequest("POST", chatCompletionsAPI, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", config.Config.OpenAI_API_Key)},
		"Content-Type":  {"application/json"},
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	stream := make(chan StreamingChatResponseChunk)
	go func() {
		defer resp.Body.Close()
		for {
			chunkReader := bufio.NewReader(resp.Body)
			body, _ := chunkReader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error : %v\n", err)
				// return err
			}
			// log.Println("received a stream chunk...")
			// log.Println(body)

			chunk := parseChunk(body)

			stream <- chunk
			if chunk.isLast() {
				close(stream)
				break
			}
		}
	}()
	return stream, nil
}

// TODO: give select, to receive error
func parseChunk(body string) StreamingChatResponseChunk {
	trimmed := strings.Split(body, "data: ")
	chunk := trimmed[1]

	var result StreamingChatResponseChunk
	err := json.Unmarshal([]byte(chunk), &result)
	if err != nil {
		log.Fatalf("Error : %v\n", err)
		// return err
	}

	// asStrArr := strings.Split(string(body), "data: ")
	// for i := range asStrArr {
	// 	asStrArr[i] = strings.TrimSpace(asStrArr[i])
	// }
	// asStrArr = asStrArr[:len(asStrArr)-1]

	// asStr := strings.Join(asStrArr, ",")
	// asStr = strings.TrimLeft(asStr, ",")
	// asStr = fmt.Sprintf("[%s]", asStr)

	// result := StreamingChatResponseChunk{}

	// err := json.Unmarshal([]byte(asStr), &result)
	// if err != nil {
	// 	log.Fatalf("Error : %v\n", err)
	// 	// return err
	// }
	return result
}
