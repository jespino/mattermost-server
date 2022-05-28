package actions

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

const EmitHttpRequestID = "emit-http-request"
const EmitHttpRequestResponseID = "emit-http-request-response"

func NewEmitHttpRequest() *systembus.ActionDefinition {
	handler := func(event *systembus.Event, config map[string]string) (*systembus.Event, error) {
		url, err := applyTemplate(config["url"], event.Data)
		if err != nil {
			return nil, err
		}
		bodyString, err := applyTemplate(config["body"], event.Data)
		if err != nil {
			return nil, err
		}
		contentType, err := applyTemplate(config["content-type"], event.Data)
		if err != nil {
			return nil, err
		}

		body := bytes.NewBufferString(bodyString)

		var resp *http.Response
		switch config["method"] {
		case "":
			resp, err = http.Get(url)
		case "GET":
			resp, err = http.Get(url)
		case "POST":
			resp, err = http.Post(url, contentType, body)
		case "PUT":
			req, err := http.NewRequest("PUT", url, body)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", contentType)
			resp, err = http.DefaultClient.Do(req)
		case "DELETE":
			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				return nil, err
			}
			resp, err = http.DefaultClient.Do(req)
		case "PATCH":
			req, err := http.NewRequest("PATCH", url, body)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", contentType)
			resp, err = http.DefaultClient.Do(req)
		}

		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		headers, err := json.Marshal(resp.Header)
		if err != nil {
			return nil, err
		}

		return &systembus.Event{
			ID: EmitHttpRequestResponseID,
			Data: map[string]string{
				"Body":    string(responseBody),
				"Headers": string(headers),
			},
		}, nil
	}

	return &systembus.ActionDefinition{
		ID:               EmitHttpRequestID,
		Name:             "Emit http request",
		Description:      "Emits an http request",
		ConfigDefinition: map[string]string{"url": "string", "method": "POST|GET|DELETE|PATCH|PUT", "content-type": "string", "body": "string"},
		Handler:          handler,
	}
}
