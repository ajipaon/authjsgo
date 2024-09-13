# authjsgo


[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![Go Report Card](https://goreportcard.com/badge/github.com/ajipaon/authjsgo)][goreportcard]

[license]: https://github.com/ajipaon/authjsgo/blob/master/LICENSE
[godocs]: https://godoc.org/github.com/ajipaon/authjsgo
[goreportcard]: https://goreportcard.com/report/github.com/ajipaon/authjsgo

This is a package that aims to decode jwt produced using authjs
only for auth.js or nextauth v5

## example: `config my auth.js`
```js
 const option = {
    // other configuratin
    session: { strategy: 'jwt', maxAge : 120  }, // 2 minute expired
    cookies: {
        sessionToken: {
            name: `mywebsite.session-token`, //session name token
                options: {
                httpOnly: false,
                    sameSite: 'none',
                    path: '/',
                    secure: true,
            },
        }
    },
}
```

### [documentation jwt authjs](https://authjs.dev/reference/core/jwt)

## Usage: `authjsgo.DecodeJWT ignore expired`

```go
package main

import (
	"github.com/ajipaon/authjsgo"
)

func main() {
	token := "your token"
	secret := []byte("your secret")
	salt := []byte("mywebsite.session-token") // use session name token

	payload, err := authjsgo.DecodeJWT(token, secret, salt, true) 
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(payload))
}
```

## Usage: `authjsgo.DecodeJWT wit expired`

```go
package main

import (
	"github.com/ajipaon/authjsgo"
)

func main() {
	
	token := "your token"
	secret := []byte("your secret")
	salt := []byte("mywebsite.session-token") // use session name

	payload, err := authjsgo.DecodeJWT(token, secret, salt, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(payload))
}
```

## Installation

```sh
go get github.com/ajipaon/authjsgo
```
## Author

[Ajipaon](https://github.com/ajipaon)