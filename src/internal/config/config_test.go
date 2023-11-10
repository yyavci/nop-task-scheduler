package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {

	tests := map[string]struct {
		configPath    string
		errorExpected bool
		resultNil     bool
	}{
		"happy case": {
			configPath:    "../../config.json",
			errorExpected: false,
			resultNil:     false,
		},
		"invalid config path": {
			configPath:    "asd/config.json",
			errorExpected: true,
			resultNil:     true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := ReadConfiguration(test.configPath)
			assert.Equal(t, test.resultNil, result == nil)
			assert.Equal(t, test.errorExpected, err != nil)
		})
	}
}
