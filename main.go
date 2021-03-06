package parser

import (
	"encoding/base64"
	"net/url"
)

type parser struct {
	GetB64URLParams b64URLParamsGetter
	EncodeB64Params b64ParamsEncoder
}

func New() parser {
	b64URLDecoder := base64.RawURLEncoding.DecodeString
	urlParser := newUrlParser()
	return parser{
		GetB64URLParams: newb64URLParamsGetter(b64URLDecoder, urlParser),
		EncodeB64Params: newB64ParamsEncoder(func() paramsEncoder{
			return url.Values{}
		}, base64.RawURLEncoding.EncodeToString),
	}
}

func newUrlParser() urlParser {
	return func(q string) (valGetter, error) {
		return url.ParseQuery(q)
	}
}