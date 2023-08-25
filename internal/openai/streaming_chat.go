package openai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testcase-generator/internal/config"
)

type StreamingChatDelta struct {
	Role         string            `json:"role,omitempty"`
	Content      string            `json:"content"`
	FunctionCall map[string]string `json:"function_call,omitempty"`
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
		"Cache-Control": {"no-cache"},
		"Connection":    {"keep-alive"},
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	chunkReader := bufio.NewReader(resp.Body)
	stream := make(chan StreamingChatResponseChunk)
	go func() {
		defer resp.Body.Close()
		for {
			body, err := chunkReader.ReadBytes('\n')
			if err != nil {
				log.Fatalln(err)
			}

			chunk := parseChunk(body)

			stream <- chunk
			if chunk.isLast() {
				close(stream)
				break
			}

			// discard the next `\n`, since it's part of the `\n\n` SSE termination.
			chunkReader.ReadBytes('\n')
		}
	}()
	return stream, nil
}

// TODO: give select, to receive error
func parseChunk(body []byte) StreamingChatResponseChunk {
	data := bytes.TrimPrefix(body, []byte("data: "))
	log.Printf("Cleaned %s\n", data)
	var result StreamingChatResponseChunk
	err := json.Unmarshal(data, &result)
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
