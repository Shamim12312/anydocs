package ai

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
)

func AnthropicResponse(message string) (string, error) {
	client := anthropic.NewClient()
	response, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeSonnet4_0,
		MaxTokens: 1024,
		System: []anthropic.TextBlockParam{
			{Text: "Create an LLM-digestible summary of the following documentation. The summary has to prioritize mainly explanations on implementation, as well as to preserve code examples and the main core concepts of the documentation."},
		},
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(message)),
		},
	})
	if err != nil {
		return err.Error(), err
	}

	content := ""

	for _, block := range response.Content {
		if block.Text != "" {
			content += block.Text + "\n"
		}
	}

	return content, nil
}
