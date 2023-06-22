package openai

import "github.com/engineer-fumi/openai/chat"

type ChatAPI interface {
	Completions() chat.Completions
}

type _ChatAPI struct {
	completions chat.Completions
}

func newChatAPI(apiKey string) ChatAPI {
	return &_ChatAPI{
		completions: chat.NewCompletions(apiKey),
	}
}

func (api *_ChatAPI) Completions() chat.Completions {
	return api.completions
}
