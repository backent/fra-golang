package config

import (
	"errors"
	"fmt"
)

var errMessageNoNotificationAvailable = "no notification available"

type notificationFunc func(string) (string, string, error)

var notificationMapping map[string]map[string]notificationFunc

func init() {
	notificationMapping = make(map[string]map[string]notificationFunc)
	notificationMapping["approve"] = make(map[string]notificationFunc)
	notificationMapping["submit"] = make(map[string]notificationFunc)
	notificationMapping["reject"] = make(map[string]notificationFunc)
	notificationMapping["update"] = make(map[string]notificationFunc)

	notificationMapping["approve"]["author"] = func(documentTitle string) (string, string, error) {
		title := "Your assessment has been Approved"
		subtitle := fmt.Sprintf(`Your assessment "%s" has been Approved`, documentTitle)
		return title, subtitle, nil
	}
	notificationMapping["approve"]["reviewer"] = func(documentTitle string) (string, string, error) {
		return "", "", errors.New(errMessageNoNotificationAvailable)
	}
	notificationMapping["approve"]["superadmin"] = func(documentTitle string) (string, string, error) {
		return "", "", errors.New(errMessageNoNotificationAvailable)
	}

	notificationMapping["reject"]["author"] = func(documentTitle string) (string, string, error) {
		title := "Your assessment has been Returned"
		subtitle := fmt.Sprintf(`Your assessment "%s" has been Returned`, documentTitle)
		return title, subtitle, nil
	}
	notificationMapping["reject"]["reviewer"] = func(documentTitle string) (string, string, error) {
		return "", "", errors.New(errMessageNoNotificationAvailable)
	}
	notificationMapping["reject"]["superadmin"] = func(documentTitle string) (string, string, error) {
		return "", "", errors.New(errMessageNoNotificationAvailable)
	}

	notificationMapping["submit"]["author"] = func(documentTitle string) (string, string, error) {
		return "", "", errors.New(errMessageNoNotificationAvailable)
	}
	notificationMapping["submit"]["reviewer"] = func(documentTitle string) (string, string, error) {
		title := "There is an assessment that has been Submitted"
		subtitle := fmt.Sprintf(`Assessment "%s" has been Submitted`, documentTitle)
		return title, subtitle, nil
	}
	notificationMapping["submit"]["superadmin"] = func(documentTitle string) (string, string, error) {
		title := "There is an assessment that has been Submitted"
		subtitle := fmt.Sprintf(`Assessment "%s" has been Submitted`, documentTitle)
		return title, subtitle, nil
	}

	notificationMapping["update"]["author"] = func(documentTitle string) (string, string, error) {
		return "", "", errors.New(errMessageNoNotificationAvailable)
	}
	notificationMapping["update"]["reviewer"] = func(documentTitle string) (string, string, error) {
		title := "There is an assessment that has been Updated"
		subtitle := fmt.Sprintf(`Assessment "%s" has been Updated`, documentTitle)
		return title, subtitle, nil
	}
	notificationMapping["update"]["superadmin"] = func(documentTitle string) (string, string, error) {
		title := "There is an assessment that has been Updated"
		subtitle := fmt.Sprintf(`Assessment "%s" has been Updated`, documentTitle)
		return title, subtitle, nil
	}

}

func NotificationGenerator(role string, action string, documentTitle string) (title string, subtitle string, err error) {
	funcNotif, found := notificationMapping[action][role]
	if found {
		return funcNotif(documentTitle)
	}
	return "", "", errors.New(errMessageNoNotificationAvailable)
}
