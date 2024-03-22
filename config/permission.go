package config

import (
	"errors"

	"github.com/backent/fra-golang/models"
)

var errMessageNoPermissionAvailable = "no permission available"

type permissionFunc func(user models.User) (bool, error)

var permissionMapping map[string]map[string]permissionFunc

func init() {
	permissionMapping = make(map[string]map[string]permissionFunc)

	permissionMapping["approve"] = make(map[string]permissionFunc)

	permissionMapping["approve"]["author"] = func(user models.User) (bool, error) {
		return false, errors.New(errMessageNoPermissionAvailable)
	}
	permissionMapping["approve"]["reviewer"] = func(user models.User) (bool, error) {
		return true, nil
	}

	permissionMapping["reject"] = make(map[string]permissionFunc)

	permissionMapping["reject"]["author"] = func(user models.User) (bool, error) {
		return false, errors.New(errMessageNoPermissionAvailable)
	}
	permissionMapping["reject"]["reviewer"] = func(user models.User) (bool, error) {
		return true, nil
	}

	permissionMapping["submit"] = make(map[string]permissionFunc)

	permissionMapping["submit"]["author"] = func(user models.User) (bool, error) {
		return true, nil
	}
	permissionMapping["submit"]["reviewer"] = func(user models.User) (bool, error) {
		return true, nil
	}

	permissionMapping["update"] = make(map[string]permissionFunc)

	permissionMapping["update"]["author"] = func(user models.User) (bool, error) {
		return true, nil
	}
	permissionMapping["update"]["reviewer"] = func(user models.User) (bool, error) {
		return true, nil
	}

	permissionMapping["draft"] = make(map[string]permissionFunc)

	permissionMapping["draft"]["author"] = func(user models.User) (bool, error) {
		return true, nil
	}
	permissionMapping["draft"]["reviewer"] = func(user models.User) (bool, error) {
		return true, nil
	}

}

func ValidatePermission(user models.User, action string) (bool, error) {
	funcNotif, found := permissionMapping[action][user.Role]
	if found {
		return funcNotif(user)
	}
	return false, errors.New(errMessageNoPermissionAvailable)
}
