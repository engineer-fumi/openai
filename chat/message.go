package chat

import (
	"fmt"
	"regexp"
)

var (
	compareName = regexp.MustCompile("^[a-zA-Z0-9_]*$")
)

type Message struct {
	Role         Role        `json:"role"`
	Content      string      `json:"content,omitempty"`
	Name         string      `json:"name,omitempty"`
	FunctionCall interface{} `json:"functionCall,omitempty"`
}

func isInvalidMessage(message Message) (err error) {
	switch message.Role {
	case RoleSystem, RoleUser, RoleAssistant, RoleFunction:
		break
	case "":
		return fmt.Errorf("role is empty")
	default:
		return fmt.Errorf("invalid role")
	}

	if message.Content == "" {
		if message.Role == RoleAssistant {
			if message.FunctionCall == nil {
				return fmt.Errorf("content is empty")
			}
		} else {
			return fmt.Errorf("content is empty")
		}
	}

	if message.Name != "" {
		if !compareName.Match([]byte(message.Name)) {
			return fmt.Errorf("name contains invalid character")
		}
		if 64 < len(message.Name) {
			return fmt.Errorf("name is too long")
		}
	} else {
		if message.Role == RoleFunction {
			return fmt.Errorf("name is empty")
		}
	}
	return
}
