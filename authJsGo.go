package authjsgo

import (
	"errors"
	"github.com/ajipaon/authjsgo/lib"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/json"
	"time"
)

func DecodeJWT(token string, secret, salt []byte, ignoreExp ...bool) ([]byte, error) {

	if len(ignoreExp) == 0 {
		ignoreExp = []bool{false}
	}
	if token == "" {
		return nil, nil
	}
	jws, err := jose.ParseEncrypted(token)

	if err != nil {
		return nil, errors.New("failed to parse encrypted JWT")
	}
	getKey, err := authjsgo.GeneratedKey(jws, secret, salt)

	if err != nil {
		return nil, err
	}
	var claim map[string]interface{}
	data, err := jws.Decrypt(getKey)

	if err != nil {
		return nil, err
	}

	if ignoreExp[0] {
		return data, nil
	}

	err = json.Unmarshal(data, &claim)
	exp, ok := claim["exp"].(float64)

	if !ok {
		return nil, errors.New("JWT does not contain an exp claim")
	}

	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("JWT has expired")
	}

	return data, nil
}
