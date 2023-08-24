package internal

import (
	"fmt"
	"log"
	"testcase-generator/internal/openai"
)

func GenerateTestcases(featureExplation string) (string, error) {
	messages := []openai.Message{
		{
			Role:    openai.System,
			Content: "You are a software tester, your job is to generate testcase scenarios for the following feature.",
		},
		{
			Role:    openai.User,
			Content: featureExplation,
		},
	}

	resp, err := openai.Chat(openai.NewChatRequest(messages))
	if err != nil {
		return "", fmt.Errorf("error while asking chat completion api : %v", err)
	}
	log.Println(resp)
	return "", nil
}
