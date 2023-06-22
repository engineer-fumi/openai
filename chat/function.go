package chat

import "fmt"

type Function struct {
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Parameters  []interface{} `json:"parameters,omitempty"`
}

func isInvalidFunction(function Function) (err error) {
	if function.Name == "" {
		return fmt.Errorf("name is empty")
	}
	if !compareName.Match([]byte(function.Name)) {
		return fmt.Errorf("name contains invalid character")
	}
	if 64 < len(function.Name) {
		return fmt.Errorf("name is too long")
	}
	return
}
