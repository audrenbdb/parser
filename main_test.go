package parser

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetB64URLParams_Integration(t *testing.T) {
	//Integration test pass
	t.Skip()
	url := "arg1=test1&arg2=test123"

	inputURL := base64.RawURLEncoding.EncodeToString([]byte(url))
	inputParams := []string{"arg1", "arg2"}
	expectedOutput := map[string]string{"arg1":"test1","arg2":"test123"}

	parser := New()
	outputParams, err := parser.GetB64URLParams(inputURL, inputParams...)
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, outputParams)

	outputParams, err = parser.GetB64URLParams(url, inputParams...)
	assert.Nil(t, outputParams)
	assert.Error(t, err)


}
