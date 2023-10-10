package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostJsonRequest(t *testing.T) {

	tests := []struct {
		url      string
		json     string
		response HttpResponse
		err      error
	}{
		{
			url:  "https://postman-echo.com/post",
			json: `"{"hey":1}"`,
			err:  nil,
			response: HttpResponse{
				StatusCode: 200,
			},
		},
	}

	for _, test := range tests {
		result, _ := PostJsonRequest(test.url, test.json)
		assert.Equal(t, test.response.StatusCode, result.StatusCode)
	}
}
