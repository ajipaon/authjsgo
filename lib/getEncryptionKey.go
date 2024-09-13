package authjsgo

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"golang.org/x/crypto/hkdf"
	"io"
)

func getDerivedEncryptionKey(enc string, keyMaterial, salt []byte) ([]byte, error) {
	var length int
	switch enc {
	case "A256CBC-HS512":
		length = 64
	case "A256GCM":
		length = 32
	default:
		return nil, errors.New("unsupported JWT content encryption algorithm")
	}

	hash := sha256.New
	reader := hkdf.New(hash, keyMaterial, salt, []byte(fmt.Sprintf("Auth.js Generated Encryption Key (%s)", salt)))
	key := make([]byte, length)
	if _, err := io.ReadFull(reader, key); err != nil {
		return nil, err
	}
	return key, nil
}
