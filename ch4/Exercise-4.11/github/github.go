// Provides Go API for working with Github Issues
package github

import (
	"encoding/base64"
	"fmt"
)

var auth string

// set OAuth username and access token
func OAuth(u, t string) {
	auth = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", u, t)))
}