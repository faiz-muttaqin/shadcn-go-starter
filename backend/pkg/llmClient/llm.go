package llmClient

import (
	"log"
	"os"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

var OLLAMA *ollama.LLM
var GPTLLM *openai.LLM
var LLM_MODEL = "gpt-4.1-mini"

// InitLLM creates and returns a reusable OpenAI LLM instance.
func InitLLM() *openai.LLM {
	llm, err := openai.New(
		openai.WithModel(LLM_MODEL),
		// ganti model kalau perlu
	)
	if err != nil {
		log.Fatalf("failed to initialize LLM: %v", err)
	}
	return llm
}

func InitOLLAMA() *ollama.LLM {
	model := os.Getenv("OLLLAMA_MODEL")
	baseURL := os.Getenv("OLLLAMA_API_URL")

	llm, err := ollama.New(
		ollama.WithModel(model),
		ollama.WithServerURL(baseURL), // penting untuk remote host
	)
	if err != nil {
		logrus.Errorf("failed to initialize Ollama LLM: %v", err)
	}
	return llm
}

// CountTokens counts the number of tokens in a given string for a specific model.
func CountTokens(text string) int {
	enc, err := tiktoken.EncodingForModel(LLM_MODEL)
	if err != nil {
		log.Printf("fallback to cl100k_base encoder: %v", err)
		enc, err = tiktoken.GetEncoding("cl100k_base")
		if err != nil {
			log.Fatalf("failed to get encoder: %v", err)
		}
	}

	tokens := enc.Encode(text, nil, nil)
	return len(tokens)
}
