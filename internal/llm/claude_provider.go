package llm

import (
    "log"
    "context"

    anthropic "github.com/anthropics/anthropic-sdk-go"

    "shaimaa/chatbot/internal/domain"
)

type ClaudeProvider struct {
    client anthropic.Client
    model anthropic.Model
}

func NewClaudeProvider() *ClaudeProvider {
    return &ClaudeProvider{
        client: anthropic.NewClient(),
        model: anthropic.Model("claude-sonnet-4-20250514"),
    }
}

func (p * ClaudeProvider) Chat(history []domain.ChatMessage) string {
    messages := []anthropic.MessageParam{}
     for _, m := range history {
        if m.Role == domain.UserRole {
            messages = append(messages, anthropic.NewUserMessage(anthropic.NewTextBlock(m.Content)))
        } else {
            messages = append(messages, anthropic.NewAssistantMessage(anthropic.NewTextBlock(m.Content)))
        }
     }

    resp, err := p.client.Messages.New(
        context.Background(),
        anthropic.MessageNewParams{
            Model: p.model,
            Messages: messages,
            MaxTokens: 1024,
        },
    )

    if err != nil {
        log.Fatal("Error:", err)
    }

    return resp.Content[0].Text
}