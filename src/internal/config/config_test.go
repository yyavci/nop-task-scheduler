package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {

	tests := map[string]struct {
		expectedErr	error
	}{
		"happy case": {
			expectedErr: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result,err := ReadConfiguration("../../config.json")

			assert.NotNil(t,result)
			assert.Equal(t,test.expectedErr,err)
		})
	}
}