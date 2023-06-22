package chat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Completions interface {
	Request(req *RequestBody) (res *ResponseBody, err error)
}

type completionsImpl struct {
	endPoint    string
	contentType string
	apiKey      string
}

func NewCompletions(apiKey string) Completions {
	return &completionsImpl{
		endPoint:    "https://api.openai.com/v1/chat/completions",
		contentType: "application/json",
		apiKey:      apiKey,
	}
}

func (c *completionsImpl) Request(req *RequestBody) (res *ResponseBody, err error) {
	if req == nil {
		return nil, fmt.Errorf("request body is nil")
	}

	var reqJson []byte
	if reqJson, err = req.MarshalJson(); err != nil {
		return
	}

	var request *http.Request
	if request, err = http.NewRequest(http.MethodPost, c.endPoint, bytes.NewBuffer(reqJson)); err != nil {
		return
	}
	request.Header.Set("Content-Type", c.contentType)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}

	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return
	}

	defer func() {
		if r := response.Body.Close(); r != nil {
			err = r
		}
	}()

	var body []byte
	if body, err = io.ReadAll(response.Body); err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		err = errors.Join(fmt.Errorf("invalid status code: %d", response.StatusCode), err)
		return
	}

	err = json.Unmarshal(body, &res)
	return
}
