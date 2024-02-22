package ldap

import (
	"fmt"
	"testing"

	"github.com/backent/fra-golang/helpers"
)

func TestLoginLdap(t *testing.T) {
	fmt.Println("test")
	token, err := helpers.LoginLdap("402746", "Bwgclp24")
	if err != nil {
		panic(err)
	}

	user, _ := helpers.GetUserLdap("830103", token)

	fmt.Println(user)
}
