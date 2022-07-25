package tokens

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	hash "productwebapi/hashkey"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("key_ABCXXZZLogin")

//var credJwt hash.Credentials

type Claims struct {
	Username string `json:"username"`
	Usermail string `json:"usermail"`
	jwt.StandardClaims
}

//>> w  is the http.ResponseWriter.
//PtrStructCookies is an *interface{} e.g.
// that mycookies := &http.Cookie{Name: "token", Value: tokenString, Expires: expirationTime}
//that get the &http.Cookies to set in browser.
//CredJwt is a hash.Credentials(pkg) struct type
//to get info and set in payload for cookies.

func GenerateToken(w http.ResponseWriter, credJwt hash.Credentials) (createdCookies bool) {

	expirationTime := time.Now().Add(time.Minute * 2)

	cliams := &Claims{
		Username:       credJwt.UserName,
		Usermail:       credJwt.UserEmail,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	tokenString, err := token.SignedString(JwtKey)
	fmt.Println("token; tokenString :: ", token, tokenString)
	if err != nil {
		log.Println("if err in tokenString , http.StatusInternalServerError : ", http.StatusInternalServerError)
		//w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	mycookies := &http.Cookie{Name: "token", Value: tokenString, Expires: expirationTime}
	http.SetCookie(w, mycookies)
	return true
}

func ValidateToken(r *http.Request) (valueCookie string, IsValidCookie, IsExpiredCookie bool, err error) {
	claim := &Claims{}
	readCookieByUser, errCookie := r.Cookie("token")
	if errCookie != nil {
		if errCookie == http.ErrNoCookie {
			log.Println(http.StatusUnauthorized)
			return "", false, false, errors.New(fmt.Sprintln(http.ErrNoCookie))
		} else {
			log.Println(http.StatusBadRequest)
			return "", false, false, errors.New(fmt.Sprintln(http.StatusBadRequest))
		}
	}
	log.Println("Got Cookie From User.")
	tokenString := readCookieByUser.Value
	if tokenString == "" {
		log.Println(http.StatusUnauthorized)
		return "", false, false, errors.New(fmt.Sprintln(http.ErrNoCookie))
	}
	tokenKey, errParseTokenKey := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if errParseTokenKey != nil {
		if errParseTokenKey == jwt.ErrSignatureInvalid {
			log.Println("Unauthorized User !.")
			return "", false, false, errors.New("unauthorized user")
		} else {
			log.Println("Could not verified to user, Try again !")
			return tokenString, false, true, errors.New("could not verified to user, Try again")
		}
	}

	if !tokenKey.Valid {
		log.Println("Token Key Is Not Valid.")
		return tokenString, false, true, errors.New("token key is not valid")
	}

	return tokenString, true, false, nil
}
