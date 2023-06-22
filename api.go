package openai

type API interface {
	Chat() ChatAPI
}

type _API struct {
	chatAPI ChatAPI
}

func NewAPI(apiKey string) API {
	return &_API{
		chatAPI: newChatAPI(apiKey),
	}
}

func (api *_API) Chat() ChatAPI {
	return api.chatAPI
}
