package llm

import (
    "context"

    openai "github.com/openai/openai-go"

    "shaimaa/chatbot/internal/domain"
)

type OpenAIProvider struct {
    client openai.Client
    model string
}

func NewOpenAIProvider() *OpenAIProvider {
    return &OpenAIProvider{
        client: openai.NewClient(),
        model: "gpt-4.1-mini", // or another model you prefer
    }
}

func (p *OpenAIProvider) Chat(history []domain.ChatMessage) string {
    messages := []openai.ChatCompletionMessageParamUnion{}
     for _, m := range history {
        if m.Role == domain.UserRole {
            messages = append(messages, openai.UserMessage(m.Content))
        } else {
            messages = append(messages, openai.AssistantMessage(m.Content))
        }
     }

    resp, _ := p.client.Chat.Completions.New(
        context.Background(),
        openai.ChatCompletionNewParams{
            Model: p.model,
            Messages: messages,
        },
    )

    return resp.Choices[0].Message.Content
}