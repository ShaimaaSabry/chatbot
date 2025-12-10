package main

import (
    "bufio"
    "fmt"
    "os"
    "log"

    "shaimaa/chatbot/internal/domain"
    "shaimaa/chatbot/internal/llm"
)


type AIProvider interface {
    Chat(history []domain.ChatMessage) string
}

const claudeProvider = "claude"
const openaiProvider = "openai"
const ollamaProvider = "ollama"
const aiProvider = openaiProvider

func main() {
    fmt.Println("Simple Chatbot using API provider: ", aiProvider)
    fmt.Println("-----------------------------")

    reader := bufio.NewReader(os.Stdin)

    provider := getProvider()
    history := []domain.ChatMessage{}

    for {
        fmt.Print("You: ")
        userInput, _ := reader.ReadString('\n')

        history = append(
            history,
            domain.ChatMessage{
                Role: domain.UserRole,
                Content: userInput,
            },
        )

        botReply := provider.Chat(history)

        history = append(
            history,
            domain.ChatMessage{
                Role: domain.BotRole,
                Content: botReply,
            },
        )

        fmt.Println("Bot: ", botReply)
    }
}

func getProvider() AIProvider {
    var provider AIProvider

    switch(aiProvider) {
        case openaiProvider:
            provider = llm.NewOpenAIProvider()
        case claudeProvider:
            provider = llm.NewClaudeProvider()
        case ollamaProvider:
            provider = llm.NewOllamaProvider()
        default:
            log.Fatal("Invalid provider flag")
    }

    return provider
}