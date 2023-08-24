package internal

import (
	"fmt"
	"testcase-generator/internal/openai"
)

func testRoleMessages(featureExplation string) []openai.Message {
	return []openai.Message{
		{
			Role:    openai.System,
			Content: "You are a software tester, your job is to generate testcase scenarios for the following feature.",
		},
		{
			Role:    openai.User,
			Content: featureExplation,
		},
	}
}

func GenerateTestcases(featureExplation string) (string, error) {
	messages := testRoleMessages(featureExplation)
	resp, err := openai.Chat(openai.NewChatRequest(messages))
	if err != nil {
		return "", fmt.Errorf("error while asking chat completion api : %v", err)
	}
	// log.Println(resp)
	return resp.Choices[0].Message.Content, nil
}

func StreamingGenerateTestcases(featureExplation string) (<-chan openai.StreamingChatResponseChunk, error) {
	messages := testRoleMessages(featureExplation)
	stream, err := openai.StreamingChat(openai.NewChatRequest(messages))
	if err != nil {
		return nil, fmt.Errorf("error while asking chat completion api streaming : %v", err)
	}
	return stream, nil
}
