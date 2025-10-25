package openai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Client *http.Client
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		APIKey: apiKey,
	}
}

func (c *Client) GetCompletions(prompt string) (string, error) {
	requestBody := map[string]interface{}{
		"prompt": prompt,
		"model":  "gpt-4o-mini",
		"max_tokens": 150,
		"temperature": 0.7,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", strings.NewReader(string(jsonBody)))
	if err != nil {
		return "", err
	}
	
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("OpenAI API error: %s", string(body))
	}

	var completionsResponse CompletionsResponse
	if err := json.Unmarshal(body, &completionsResponse); err != nil {
		return "", err
	}
	
	if len(completionsResponse.Choices) == 0 {
		return "", fmt.Errorf("no completions returned")
	}
	
	return completionsResponse.Choices[0].Text, nil
}

type CompletionsResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
	Index int `json:"index"`
	FinishReason string `json:"finish_reason"`
	FinishReasonDetail FinishReasonDetail `json:"finish_reason_detail"`
}

type FinishReasonDetail struct {
	Type string `json:"type"`
}