package jwt

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Claims :
type Claims struct {
	ID int64
	jwt.StandardClaims
}

//ParseToken :
func ParseToken(token string) (*jwt.Token, error) {

	return jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		verifyBytes, err := ioutil.ReadFile("./certificates/public.key")
		if err != nil {
			panic(err)
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			panic(err)
		}
		return verifyKey, nil
	})
}

//CreateTokenWithClaims :
func CreateTokenWithClaims(id int64) string {

	//valid for 60 days
	expirationTime := time.Now().Add(24 * 60 * time.Hour)

	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "anshu",
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	keyBytes, err := ioutil.ReadFile("./certificates/private.key")
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)

	if err != nil {
		panic(err)
	}

	// Create the JWT string
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		panic(err)
	}

	return tokenString
}
