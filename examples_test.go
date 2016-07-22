package legit

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Email Email
}

func ExampleValidate() {
	body := []byte(`{"email": "not an email"}`)

	var user User
	json.Unmarshal(body, &user)

	err := Validate(user)
	if err != nil {
		fmt.Println("invalid user!", err)
	}
}
