package jwt

import (
	"errors"
	"fmt"
	"masterplan-backend/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey string

func init() {
	jwtKey = models.Config.JWTKey
}

type jwtCustomClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//GenerateToken for verfied user
func GenerateToken(username string) (string, error) {

	claims := jwtCustomClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(models.Config.TokenExpiryTime * time.Minute).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as resptokenFromRequestonse.
	signedToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// VerifyToken verfies if the JWT token is authentic
func VerifyToken(tokenFromRequest string) (token *jwt.Token, err error) {

	token, err = jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {
		// validate the alg is what you expect
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Error while parsing JWT Token: ", err))
	}
	if !token.Valid {
		return nil, errors.New("Token Expired")
	}
	return token, nil
}

// GetValueInToken get required value from the token after passing the key
func GetValueInToken(tokenFromRequest, key string) (value string, err error) {
	token, err := VerifyToken(tokenFromRequest)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Error while getting claims")
	}

	// check if the key is present and return it as string
	if claims[key] == nil || claims[key] == "" {
		return "", errors.New(key + " not found")
	}

	return claims[key].(string), nil
}
