package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/backent/fra-golang/models"
)

type LoginResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    DataLoginResponse `json:"data"`
}

type DataLoginResponse struct {
	Auth string               `json:"auth"`
	Jwt  JWTDataLoginResponse `json:"jwt"`
}

type JWTDataLoginResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type GetUserResponse struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Data    DataGetUserResponse `json:"data"`
}

type DataGetUserResponse struct {
	DataPosisi DataPosisiDataGetUserResponse `json:"dataPosisi"`
}

type DataPosisiDataGetUserResponse struct {
	Nama  string `json:"NAMA"`
	Email string `json:"EMAIL"`
}

func LoginLdap(username string, password string) (string, error) {
	const url string = "https://apifactory.telkom.co.id:8243/hcm/auth/v1/token"

	body := fmt.Sprintf(`{
		"username": "%s",
		"password": "%s"
	}`, username, password)

	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	var response LoginResponse
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&response)

	return response.Data.Jwt.Token, nil
}

func GetUserLdap(nik string, token string) (models.UserLdap, error) {
	var userLdap models.UserLdap
	const url string = "https://apifactory.telkom.co.id:8243/hcm/pwb/v1/profile/%s"

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf(url, nik), nil)
	if err != nil {
		return userLdap, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-authorization", "Bearer "+token)

	client := http.Client{}

	res, err := client.Do(request)
	if err != nil {
		return userLdap, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return userLdap, errors.New("returned from ldap with status code: " + strconv.Itoa(res.StatusCode))
	}

	var response GetUserResponse
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&response)

	userLdap.Email = response.Data.DataPosisi.Email
	userLdap.Nik = nik
	userLdap.Name = response.Data.DataPosisi.Nama

	return userLdap, nil
}
