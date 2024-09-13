package test

import (
	"github.com/ajipaon/authjsgo"
	"testing"
)

func TestDecodeJWT(t *testing.T) {
	token := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIiwia2lkIjoiUVhJcW03TlBELW9IeWhTZGRjc016ei1YNjFSUlpEWThGNURqQTEzRWJlTkdkdnRVZk1rYXJLai1LMDZUQk83eFlSQlV4dERTclE3VWFUOGtuSDFIUVEifQ..2u6DtTw19cygUWmohhJ25A.Fe4V0I1kTFlzvqNvY_nYRCI07LhYR4iTK6EN6nDSuJuQIWa2WzCEtxa28Zyx1ZmWLYl3uW4VQ2QlrCMVjhJmuzvX9Hvs8YufvPe-lPOIGozU83ZdKxd5gG6O5t93jC1XtY_aIph8fNGTKHOQk_lwJEJGtk442Esb4u8NhToJVLwamHft4E2gwBSIrcuH6GCtYC2cJhYYJsBAjm2-EjaSXv4JWmrDHd4zIJBZAj-qwUtJg9CxjjrmuI2G6EnHC2AER6v48ldP1mcm9WmHXNGBfoE9qR2Izz06hRuN3zE3mLAshSOawHDgmxvOqsOnIaD5jUEPk7FvMIEzWwa2MAtReaC0DeUcnxtfC6hl3BDdBtPbbKBsxKlXaILq8Z0THLC7aucp2sUmR1lVRTUjSWnyJg.I31CPkRAKo-JyzTef3mOWwd3c8Hnuus0fPfsUTX4Pjc"

	secret := []byte("2f9340a0d055fa9f6d9c08e52aabb80b4ab25c7d7f6fd5dc0e3f6c14dd104e18")
	salt := []byte("mywebsite.session-token")

	payload, err := authjsgo.DecodeJWT(token, secret, salt, true)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(payload))
}
