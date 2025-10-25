package analysis

import (
	"educationallsp/openai"
	"log"
	"os"
)

func ExplainCode(logger *log.Logger, line, fullDocument string) (string, error) {
	prompt := `Explain what this line of code does in simple terms. Be concise but helpful:

Line: ` + line + `

Context (full file):
` + fullDocument + `

Explanation:`

	explanation, err := openai.NewClient(os.Getenv("OPENAI_API_KEY")).GetCompletions(prompt)
	if err != nil {
		logger.Printf("Error getting completions: %s", err)
		return "", err
	}
	logger.Printf("Explanation: %s", explanation)
	return explanation, err
}