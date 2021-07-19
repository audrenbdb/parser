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

func TestDecode(t *testing.T) {
	//decodedURL := "host=127.0.0.1&port=12345"
	b64EncodedURL := "aG9zdD0xMjcuMC4wLjEmcG9ydD0xMjM0NQ"

	p := New()
	d, _ := p.GetB64URLParams(b64EncodedURL, "host", "port")
	//output map[string]string{"host":"127.0.0.1","port":"12345"}

	e := p.EncodeB64Params(map[string]string{"host":"127.0.0.1","port":"12345"})
	//output aG9zdD0xMjcuMC4wLjEmcG9ydD0xMjM0NQ

	assert.Equal(t, b64EncodedURL, e)
	assert.Equal(t, map[string]string{"host":"127.0.0.1","port":"12345"}, d)
}
