package authjsgo

import (
	"fmt"
)

type Digest string

func GetDigest(encryptionSecret []byte) Digest {
	return Digest(fmt.Sprintf("sha%d", len(encryptionSecret)<<3))
}
