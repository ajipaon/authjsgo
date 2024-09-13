package authjsgo

import (
	"gopkg.in/square/go-jose.v2"
)

func GeneratedKey(token string, secret, salt []byte) ([]byte, error) {

	jws, err := jose.ParseEncrypted(token)
	if err != nil {
		return nil, err
	}
	kid := jws.Header.KeyID
	enc := jws.Header.ExtraHeaders["enc"].(string)

	encryptionSecret, err := getDerivedEncryptionKey(enc, secret, salt)

	if kid == "" {
		return encryptionSecret, nil
	}

	jwk := JWK{
		Kty: "oct",
		K:   Base64urlEncode(encryptionSecret),
	}

	thumbprint, err := calculateJwkThumbprint(jwk, string(GetDigest(encryptionSecret)))
	if err != nil {
		return nil, err
	}

	if thumbprint == kid {
		return encryptionSecret, nil
	}

	return nil, nil
}
