# Parser

Parser is a wrapper around url encode / decode native functions.

```
//decodedURL := "host=127.0.0.1&port=12345"
b64EncodedURL := "aG9zdD0xMjcuMC4wLjEmcG9ydD0xMjM0NQ"

p := parser.New()
d, _ := p.GetB64URLParams(b64EncodedURL, "host", "port")
//output map[string]string{"host":"127.0.0.1","port":"12345"}

e := p.EncodeB64Params(map[string]string{"host":"127.0.0.1","port":"12345"})
//output aG9zdD0xMjcuMC4wLjEmcG9ydD0xMjM0NQ
```
