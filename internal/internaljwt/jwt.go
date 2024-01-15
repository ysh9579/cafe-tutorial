package internaljwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const defaultSecretKey = "default key"

func GetSecretKey() string {
	k := os.Getenv("JWT_SECRET_KEY")
	if k == "" {
		k = defaultSecretKey
	}
	return k
}

func CreateJWT(Id string) (string, error) {
	secretKey := GetSecretKey()
	mySigningKey := []byte(secretKey)

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["Id"] = Id
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}
