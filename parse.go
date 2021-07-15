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

//b64URLParamsGetter decode a b64url and extracts the given parameters
type b64URLParamsGetter func(b64URL string, params ...string) (map[string]string, error)

//b64URLDecoder returns a decoded byte array from an b64 url
type b64URLDecoder func(url string) ([]byte, error)

//urlParser parses an url and returns a value getter to read parameters parsed
type urlParser func(url string) (valGetter, error)

type valGetter interface {
	Get(key string) string
}
