package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	fmt.Println("OpenAI API Key:", openaiAPIKey)

	client := openai.NewClient(openaiAPIKey)

	// Set proxy if needed  就是非常needed
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "你将会被提供一些博客下的评论，你的任务是将这些评论中的脏话过滤掉，并以文明和易于接受的方式重述出来"},
				{Role: "user", Content: "真tm有意思"},
			},
			Temperature: 0.7,
			MaxTokens:   64,
			TopP:        1,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
