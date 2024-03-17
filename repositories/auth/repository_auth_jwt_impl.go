package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/backent/fra-golang/helpers"
	"github.com/golang-jwt/jwt/v5"
)

type RepositoryAuthJWTImpl struct {
	secretKeys    []byte
	tokenLifeTime int
}

func NewRepositoryAuthJWTImpl() RepositoryAuthInterface {

	tokenLifeTime, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRE_IN_SEC"))
	helpers.PanicIfError(err)

	return &RepositoryAuthJWTImpl{
		secretKeys:    []byte(os.Getenv("APP_SECRET_KEY")),
		tokenLifeTime: tokenLifeTime,
	}
}

func (implementation *RepositoryAuthJWTImpl) Issue(payload string) (string, error) {
	// Create the Claims

	remainingTimeBeforeMidnight := calculateRemainingTimeBeforeMidnight(implementation.tokenLifeTime)

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(remainingTimeBeforeMidnight)),
		Issuer:    payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	generatedToken, err := token.SignedString(implementation.secretKeys)
	return generatedToken, err
}
func (implementation *RepositoryAuthJWTImpl) Validate(tokenString string) (int, bool) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return implementation.secretKeys, nil
	})

	if payload, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := payload["iss"].(string); ok {
			intId, err := strconv.Atoi(id)
			if err != nil {
				return 0, false
			}
			return intId, true
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
}

func calculateRemainingTimeBeforeMidnight(tokenLifeTime int) time.Duration {
	// Your token life time in seconds

	// Get the current time
	now := time.Now()

	// Calculate the end of the day
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	// Calculate the remaining duration until the end of the day
	remainingDuration := endOfDay.Sub(now)

	// Calculate the duration to add (token life time)
	durationToAdd := time.Second * time.Duration(tokenLifeTime)

	// Take the minimum of remainingDuration and durationToAdd
	var duration time.Duration
	if remainingDuration < durationToAdd {
		duration = remainingDuration
	} else {
		duration = durationToAdd
	}

	return duration
}
