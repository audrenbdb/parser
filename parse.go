package parser

func newb64URLParamsGetter(decodeURL b64URLDecoder, parseURL urlParser) b64URLParamsGetter {
	return func(b64URL string, params ...string) (map[string]string, error) {
		b, err := decodeURL(b64URL)
		if err != nil {
			return nil, err
		}
		values, err := parseURL(string(b))
		if err != nil {
			return nil, err
		}
		return getParamsValues(values, params...), nil
	}
}

func getParamsValues(v valGetter, params ...string) map[string]string {
	pval := map[string]string{}
	for _, p := range params {
		pval[p] = v.Get(p)
	}
	return pval
}

func newB64ParamsEncoder(newParamsEncoder func() paramsEncoder, encodeRawURL rawURLEncoder) b64ParamsEncoder {
	return func(params map[string]string) string {
		paramsEncoder := newParamsEncoder()
		for k, v := range params {
			paramsEncoder.Set(k, v)
		}
		b := []byte(paramsEncoder.Encode())
		return encodeRawURL(b)
	}
}



//b64URLParamsGetter decode a b64url and extracts the given parameters
type b64URLParamsGetter func(b64URL string, params ...string) (map[string]string, error)

//b64UrlParamsEncoder encode params into a b64URL string
type b64ParamsEncoder func(params map[string]string) string

//b64URLDecoder returns a decoded byte array from an b64 url
type b64URLDecoder func(url string) ([]byte, error)

//rawURLEncoder converts a byte array to a string
type rawURLEncoder func(b []byte) string

//urlParser parses an url and returns a value getter to read parameters parsed
type urlParser func(url string) (valGetter, error)

type valGetter interface {
	Get(key string) string
}

type paramsEncoder interface {
	Set(key, val string)
	Encode() string
}
