<img src="http://i.imgur.com/HzBjvdC.png" width="268" height="100" alt="LEGIT" />

[![GoDoc](https://godoc.org/github.com/jamescun/legit?status.png)](https://godoc.org/github.com/jamescun/legit)

Legit is an input validation framework for Go. Legit differs from existing frameworks by constructing validation from types and interfaces, preferring custom validators to complex struct tags.

	go get -u github.com/jamescun/legit

Included validators:

  - Email
  - UUID
  - UUID3
  - UUID4
  - UUID5
  - Credit Card
  - Lowercase
  - Uppercase
  - No whitespace
  - Printable characters
  - Alpha
  - Alphanumeric
  - Numeric
  - ASCII
  - Positive number
  - Negative number


Example
-------

```go
package main

import (
	"fmt"
	"regexp"
	"errors"
	"encoding/json"

	"github.com/jamescun/legit"
)

type Name string

// very simplistic regexp for human name validation
var expName = regexp.MustCompile(`[a-zA-Z\ ]{1,64}`)

// attach Validate() method to our custom name type, satisfying the legit.Object interface,
// defining our custom name validation.
func (n Name) Validate() error {
	if !expName.MatchString(string(n)) {
		return errors.New("invalid name")
	}

	return nil
}

type User struct {
	Email legit.Email `json:"email"`
	Name  Name        `json:"name"`
}

func main() {
	body := []byte(`{"email": "test@example.org", "name": "John Doe"}`)
	
	var user User
	json.Unmarshal(body, &user)

	err := legit.Validate(user)
	if err != nil {
		fmt.Println("invalid user!", err)
	}
}

```

