package analysis

import (
	"educationallsp/openai"
	"log"
	"os"
)

func ExplainCode(logger *log.Logger, line, fullDocument string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		logger.Printf("No OpenAI API key found, returning line content")
		return line, nil
	}

	prompt := `Explain what this line of code does in simple terms. Be concise but helpful:

Line: ` + line + `

Context (full file):
` + fullDocument + `

Explanation:`

	explanation, err := openai.NewClient(apiKey).GetCompletions(prompt)
	if err != nil {
		logger.Printf("Error getting completions: %s", err)
		return "", err
	}
	logger.Printf("Explanation: %s", explanation)
	return explanation, nil
}