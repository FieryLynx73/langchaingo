package fake

import (
	"testing"

	"github.com/FieryLynx73/langchaingo/testing/llmtest"
)

func TestLLM(t *testing.T) {
	// Fake LLM doesn't need API keys
	llm := &LLM{}

	// Test basic functionality only
	llmtest.TestLLM(t, llm)
}
