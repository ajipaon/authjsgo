package authjsgo

type JWK struct {
	Kty string `json:"kty"`
	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Y   string `json:"y,omitempty"`
	E   string `json:"e,omitempty"`
	N   string `json:"n,omitempty"`
	K   string `json:"k,omitempty"`
}
