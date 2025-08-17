package ai

import (
	"os"
	"testing"
)

func TestAnthropic(t *testing.T) {
	_, ok := os.LookupEnv("ANTHROPIC_API_KEY")
	if !ok {
		t.Skip("Skipping test because ANTHROPIC_API_KEY is not in the environment")
	}
	message, err := AnthropicResponse("There is no documentation, you are done for today!")
	if err != nil {
		t.Errorf("Expecting no error, got %s", message)
	}
}
