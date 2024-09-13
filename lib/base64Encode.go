package authjsgo

import "encoding/base64"

func Base64urlEncode(data []byte) string {

	return base64.RawURLEncoding.EncodeToString(data)
}
