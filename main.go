package parser

import (
	"encoding/base64"
	"net/url"
)

type parser struct {
	GetB64URLParams b64URLParamsGetter
}

func New() parser {
	b64URLDecoder := base64.RawURLEncoding.DecodeString
	urlParser := newUrlParser()
	return parser{
		GetB64URLParams: newb64URLParamsGetter(b64URLDecoder, urlParser),
	}
}

func newUrlParser() urlParser {
	return func(q string) (valGetter, error) {
		return url.ParseQuery(q)
	}
}