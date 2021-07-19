package parser

import (
	"encoding/base64"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestGetBase64URLParams(t *testing.T) {
	tests := []struct {
		description string

		b64URLDecoder b64URLDecoder
		urlParser urlParser

		b64URL string
		paramsIn []string

		paramsOut map[string]string
		err error
	}{
		{
			description: "should fail to decode url to bytes",
			b64URLDecoder: func(url string) ([]byte, error) {
				return nil, errors.New("fail to decode")
			},
			err: errors.New("fail to decode"),
		},
		{
			description: "should fail to parse url",
			b64URLDecoder: func(url string) ([]byte, error) {
				if url == "testurl" {
					return []byte(url), nil
				}
				return nil, errors.New("fail to decode")
			},
			urlParser: func(url string) (valGetter, error) {
				return nil, errors.New("fail to parse url")
			},

			b64URL: "testurl",
			err: errors.New("fail to parse url"),
		},
		{
			description: "should succeed in getting params",
			b64URLDecoder: func(url string) ([]byte, error) {
				if url == "arg1=test1&arg2=test2" {
					return []byte(url), nil
				}
				return nil, errors.New("fail to decode")
			},
			urlParser: func(url string) (valGetter, error) {
				if url == "arg1=test1&arg2=test2" {
					return &mockValGetter{params: map[string]string{
						"arg1":"test1", "arg2":"test2","arg3":"test3",
					}}, nil
				}
				return nil, errors.New("fail to parse url")
			},

			b64URL: "arg1=test1&arg2=test2",
			paramsIn: []string{"arg1","arg2"},
			paramsOut: map[string]string{"arg1":"test1", "arg2":"test2"},
			err: nil,
		},
	}

	for _, test := range tests {
		getParams := newb64URLParamsGetter(test.b64URLDecoder, test.urlParser)
		params, err := getParams(test.b64URL, test.paramsIn...)
		assert.Equal(t, test.err, err, test.description)
		assert.Equal(t, test.paramsOut, params, test.description)
	}
}

func TestEncodeParams(t *testing.T) {
	tests := []struct {
		description string

		inParams map[string]string

		encodedParams string
	}{
		{
			description: "should fail to encode params",
			inParams: map[string]string{"host":"1234"},
			encodedParams: "aG9zdD0xMjM0",
		},
	}

	for _, test := range tests {
		encode := newB64ParamsEncoder(func() paramsEncoder{return url.Values{}}, base64.RawURLEncoding.EncodeToString)
		encodedParams := encode(test.inParams)
		assert.Equal(t, test.encodedParams, encodedParams, test.description)
	}

}

type mockValGetter struct {
	params map[string]string
}

func (m *mockValGetter) Get(key string) string {
	return m.params[key]
}
