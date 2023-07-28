package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

type OpenaiResponse struct {
	Choices []OpenaiChoice `json:"choices"`
}

type OpenaiChoice struct {
	Message OpenaiMessage `json:"message"`
}

type OpenaiMessage struct {
	Content string `json:"content"`
}

func getStory(name string) []string {
	data, err := json.Marshal(gin.H{
		"model": "gpt-3.5-turbo",
		"messages": []gin.H{
			gin.H{
				"role":    "user",
				"content": "tell a joke about " + name,
			},
		},
	})
	if err != nil {
		return []string{err.Error()}
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(data))
	if err != nil {
		return []string{err.Error()}
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("OPENAI_KEY"))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []string{err.Error()}
	}
	bodyBytes, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return []string{err.Error()}
	}
	openaiResponse := &OpenaiResponse{}
	err = json.Unmarshal(bodyBytes, openaiResponse)
	if err != nil {
		return []string{err.Error()}
	}
	if len(openaiResponse.Choices) == 0 {
		return []string{"Something went wrong:("}
	}

	story := strings.FieldsFunc(openaiResponse.Choices[0].Message.Content, func(c rune) bool { return c == '\n' })
	return story
}
