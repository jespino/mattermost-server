package builtinactions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const EmitHttpRequestID = "emit-http-request"
const EmitHttpRequestResponseID = "emit-http-request-response"

func NewEmitHttpRequest() *actions.ActionDefinition {
	handler := func(config map[string]string, data map[string]string) (map[string]string, error) {
		url := config["url"]
		bodyString := config["body"]
		contentType := config["content-type"]

		body := bytes.NewBufferString(bodyString)

		var resp *http.Response
		var err error
		switch strings.ToUpper(config["method"]) {
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
		default:
			return nil, errors.New("invalid request method")

		}

		if err != nil {
			return nil, err
		}

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		headers, err := json.Marshal(resp.Header)
		if err != nil {
			return nil, err
		}

		return map[string]string{
			"Body":    string(responseBody),
			"Headers": string(headers),
		}, nil
	}

	return &actions.ActionDefinition{
		ID:               EmitHttpRequestID,
		Name:             "Emit http request",
		Description:      "Emits an http request",
		ConfigDefinition: map[string]string{"url": "string", "method": "POST|GET|DELETE|PATCH|PUT", "content-type": "string", "body": "string"},
		Handler:          handler,
	}
}
