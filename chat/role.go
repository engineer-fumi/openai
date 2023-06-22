package chat

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser           = "user"
	RoleAssistant      = "assistant"
	RoleFunction       = "function"
)
