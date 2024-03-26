package config

import "os"

func GetUserLDAPAccount() (string, string) {
	return os.Getenv("LDAP_ACCOUNT_USERNAME"), os.Getenv("LDAP_ACCOUNT_PASSWORD")
}
