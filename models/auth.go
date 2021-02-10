package models

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// "fmt"

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"username"`
	jwt.StandardClaims
}

//完成签名函数的剥离工作！
// Create the Signin handler
func GenarateJwt(email string) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Println(err)
		return "bad", err
	}
	return tokenString, nil

}

type VerifyTimeError struct {
	info string
}

func (e *VerifyTimeError) Error() string {
	return fmt.Sprintf("%v: 还在有效时间内", e.info)
}

func Refresh(oldtoken string) (string, error) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	tknStr := oldtoken
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "bad", err
		}
	}
	if !tkn.Valid {
		return "bad", jwt.ErrSignatureInvalid
	}
	// (END) The code up-till this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	//不超出时间就不写入新的token，这里就不再设置新的cookie
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return "time", &VerifyTimeError{info: "还在有效时间内"}
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "bad", err
	}
	return tokenString, nil

}

func VerifyJwt(token string) (string, error) {
	// We can obtain the session token from the requests cookies, which come with every request

	// Get the JWT string from the cookie
	tknStr := token
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println(err)
			// w.WriteHeader(http.StatusUnauthorized)
			return "bad", err
		}
	}
	if !tkn.Valid {
		// w.WriteHeader(http.StatusUnauthorized)
		return "bad", jwt.ErrSignatureInvalid
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	fmt.Sprintf("Welcome %s!", claims.Email)
	return claims.Email, nil
}
