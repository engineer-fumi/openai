package chat

import (
	"encoding/json"
	"fmt"
)

type RequestBody struct {
	Model            Model       `json:"model"`
	Messages         []Message   `json:"messages"`
	Functions        []Function  `json:"functions,omitempty"`
	FunctionCall     interface{} `json:"functionCall,omitempty"`
	Temperature      *float64    `json:"temperature,omitempty"`
	TopP             *float64    `json:"top_p,omitempty"`
	N                *int        `json:"n,omitempty"`
	Stream           *bool       `json:"stream,omitempty"`
	Stop             interface{} `json:"stop,omitempty"`
	MaxTokens        *int        `json:"max_tokens,omitempty"`
	PresencePenalty  *float64    `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64    `json:"frequency_penalty,omitempty"`
	LogitBias        interface{} `json:"logit_bias,omitempty"`
	User             *string     `json:"user,omitempty"`
}

func (req *RequestBody) MarshalJson() (data []byte, err error) {
	if req == nil {
		return nil, fmt.Errorf("request body is nil")
	}
	if req.Model == "" {
		return nil, fmt.Errorf("model is empty")
	}
	if req.Messages != nil {
		for _, message := range req.Messages {
			if err = isInvalidMessage(message); err != nil {
				return nil, err
			}
		}
	} else {
		return nil, fmt.Errorf("messages is nil")
	}
	if req.Functions != nil {
		for _, function := range req.Functions {
			if err = isInvalidFunction(function); err != nil {
				return nil, err
			}
		}
	}
	if req.Temperature != nil && req.TopP != nil {
		return nil, fmt.Errorf("temperature and top_p cannot be specified at the same time")
	}
	if req.Temperature != nil {
		if *req.Temperature < 0 || 2 < *req.Temperature {
			return nil, fmt.Errorf("invalid temperature value")
		}
	}
	if req.TopP != nil {
		if *req.TopP < 0 || 1 < *req.TopP {
			return nil, fmt.Errorf("invalid top_p value")
		}
	}
	if req.PresencePenalty != nil {
		if *req.PresencePenalty < -2.0 || 2.0 < *req.PresencePenalty {
			return nil, fmt.Errorf("invalid presence_penalty value")
		}
	}
	data, err = json.Marshal(req)
	return
}
