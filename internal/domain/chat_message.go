package domain

const UserRole = "user"
const BotRole = "assistant"

type ChatMessage struct {
    Role string
    Content string
}