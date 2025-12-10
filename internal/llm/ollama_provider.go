package llm

import (
    "log"

    "net/http"
    "encoding/json"
    "io"
    "bytes"

    "shaimaa/chatbot/internal/domain"
)

type OllamaProvider struct {
    url string
    model string
}

func NewOllamaProvider() *OllamaProvider {
    return &OllamaProvider{
        url: "http://localhost:11434/v1/completions",
        model: "llama2",
    }
}

func (p *OllamaProvider) Chat(history []domain.ChatMessage) string {
    prompt := ""
    for _, m := range history {
        if m.Role == domain.UserRole {
            prompt += "User: " + m.Content + "\n"
        } else {
            prompt += "Assistant: " + m.Content + "\n"
        }
    }
    prompt += "Assistant: " // asking model to continue

    messages := []map[string]string{}
    for _, m := range history {
        messages = append(messages, map[string]string{
            "role":    m.Role,
            "content": m.Content,
        })
    }

    body := map[string]interface{}{
        "model":  p.model,
        "prompt": prompt,
        "messages": messages,
    }

    bodyJson, _ := json.Marshal(body)
    resp, err := http.Post(
        p.url,
        "application/json",
        bytes.NewBuffer(bodyJson),
    )

    if err != nil {
        log.Fatal("Error making request: ", err)
    }

    defer resp.Body.Close()

    data, _ := io.ReadAll(resp.Body)
    type OllamaResponse struct {
        Choices []struct {
            Text string `json:"text"`
        } `json:"choices"`
    }
    var result OllamaResponse
    if err := json.Unmarshal(data, &result); err != nil {
        log.Fatal("Error unmarshaling response:", err)
    }

    return result.Choices[0].Text
}