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

## Usage: `ignore expired`

```go
package main

import (
	"github.com/ajipaon/authjsgo"
	"encoding/json"
	"fmt"
)

func main() {
	token := "your token"
	secret := []byte("your secret")
	salt := []byte("mywebsite.session-token") // use session name

	payload, _ := authjsgo.DecodeJWT(token, secret, salt, true)
	
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		fmt.Println("Error decode json:", err)
		return
	}
	fmt.Println(claims)
}
```

## Usage: `with expired`

```go
package main

import (
	"github.com/ajipaon/authjsgo"
	"encoding/json"
	"fmt"
)

func main() {
	
	token := "your token"
	secret := []byte("your secret")
	salt := []byte("mywebsite.session-token") // use session name

	payload, _ := authjsgo.DecodeJWT(token, secret, salt)
	
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		fmt.Println("Error decode json:", err)
		return
	}
	fmt.Println(claims)
}
```

## Installation

```sh
go get github.com/ajipaon/authjsgo
```
## Author

[Ajipaon](https://github.com/ajipaon)