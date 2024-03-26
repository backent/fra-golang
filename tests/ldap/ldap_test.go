package ldap

import (
	"fmt"
	"testing"

	"github.com/backent/fra-golang/config"
	"github.com/backent/fra-golang/helpers"
)

func TestLoginLdap(t *testing.T) {
	fmt.Println("test")
	usernameLDAP, passwordLDAP := config.GetUserLDAPAccount()
	token, err := helpers.LoginLdap(usernameLDAP, passwordLDAP)
	if err != nil {
		panic(err)
	}

	user, _ := helpers.GetUserLdap("830103", token)

	fmt.Println(user)
}
