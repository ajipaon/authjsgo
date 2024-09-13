package authjsgo

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"hash"
)

// this code based from https://github.com/panva/jose/blob/main/src/jwk/thumbprint.ts
// @param jwk The JWK to calculate the thumbprint of.
// @param digestAlgorithm The digest algorithm to use to calculate the thumbprint.
// @returns The JWK thumbprint.

func calculateJwkThumbprint(jwk JWK, digestAlgorithm string) (string, error) {
	if digestAlgorithm == "" {
		digestAlgorithm = "sha256"
	}

	hashFunc, err := getHashFunc(digestAlgorithm)
	if err != nil {
		return "", err
	}

	components, err := getJwkComponents(jwk)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(components)
	if err != nil {
		return "", err
	}

	hashes := hashFunc()
	hashes.Write(data)
	bytes := hashes.Sum(nil)
	return Base64urlEncode(bytes), nil
}

func getHashFunc(digestAlgorithm string) (func() hash.Hash, error) {
	switch digestAlgorithm {
	case "sha256":
		return sha256.New, nil
	case "sha384":
		return sha512.New384, nil
	case "sha512":
		return sha512.New, nil
	default:
		return nil, errors.New(`digestAlgorithm must be one of "sha256", "sha384", or "sha512"`)
	}
}

func getJwkComponents(jwk JWK) (map[string]string, error) {
	switch jwk.Kty {
	case "EC":
		if jwk.Crv == "" || jwk.X == "" || jwk.Y == "" {
			return nil, errors.New(`"crv", "x", and "y" are required for EC keys`)
		}
		return map[string]string{"crv": jwk.Crv, "kty": jwk.Kty, "x": jwk.X, "y": jwk.Y}, nil
	case "OKP":
		if jwk.Crv == "" || jwk.X == "" {
			return nil, errors.New(`"crv" and "x" are required for OKP keys`)
		}
		return map[string]string{"crv": jwk.Crv, "kty": jwk.Kty, "x": jwk.X}, nil
	case "RSA":
		if jwk.E == "" || jwk.N == "" {
			return nil, errors.New(`"e" and "n" are required for RSA keys`)
		}
		return map[string]string{"e": jwk.E, "kty": jwk.Kty, "n": jwk.N}, nil
	case "oct":
		if jwk.K == "" {
			return nil, errors.New(`"k" is required for oct keys`)
		}
		return map[string]string{"k": jwk.K, "kty": jwk.Kty}, nil
	default:
		return nil, errors.New(`"kty" parameter missing or unsupported`)
	}
}
