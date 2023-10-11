package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostJsonRequest(t *testing.T) {

	tests := map[string]struct {
		url         string
		json        string
		expected    HttpResponse
		expectedErr error
	}{
		"happy case": {
			url:         "https://postman-echo.com/post",
			json:        `"{"hey":1}"`,
			expectedErr: nil,
			expected: HttpResponse{
				StatusCode: 200,
				Status:     "200 OK",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, _ := PostJsonRequest(test.url, test.json)

			assert.Equal(t, test.expected.StatusCode, result.StatusCode)
			assert.Equal(t, test.expected.Status, result.Status)
		})
	}
}

func TestPostJsonRequestErrors(t *testing.T) {
	tests := map[string]struct {
		url              string
		json             string
		isExpectingError bool
		expectedMessage  string
	}{
		"empty json": {
			url:              "https://postman-echo.com/post",
			json:             "",
			isExpectingError: true,
			expectedMessage:  "json is empty",
		},
		"empty url": {
			url:              "",
			json:             `"{"hey":1}"`,
			isExpectingError: true,
			expectedMessage:  "url is empty",
		},
		"invalid url": {
			url:              "https1://postman-echo.com/post",
			json:             `"{"hey":1}"`,
			isExpectingError: true,
			expectedMessage:  "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := PostJsonRequest(test.url, test.json)

			assert.Equal(t, test.isExpectingError, err != nil)

			if err != nil && len(test.expectedMessage) > 0 {
				assert.Equal(t, test.expectedMessage, err.Error())
			}
		})
	}
}
