package signin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	con "productwebapi/connection"
	hash "productwebapi/hashkey"
	"productwebapi/tokens"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var myDB *gorm.DB
var cred hash.Credentials

type Claims struct {
	Username string `json:"username"`
	Usermail string `json:"usermail"`
	jwt.StandardClaims
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	myDB = con.RegdMigration()
	var getCred hash.Credentials
	r.Header.Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	if errInDecode := json.NewDecoder(r.Body).Decode(&getCred); errInDecode != nil {
		log.Println("errInDecode : ", errInDecode)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("jsondecode getCred : ", getCred)

	fmt.Println("getCred : ", getCred)
	fmt.Println(" mail, psw,  = ", getCred.UserEmail, getCred.Password)
	if keyErr := myDB.Where("user_email = ?", getCred.UserEmail).First(&cred).Error; keyErr != nil {
		fmt.Println("keyErr : ", keyErr)
		w.Write([]byte("User not found !"))
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}
	okPassword := cred.CheckPassword(getCred.Password)

	if getCred.UserName != cred.UserName && !okPassword {
		log.Println("Wrong Password.")
		w.Write([]byte("Wrong Password."))
		return
	}

	if okPassword {
		okCookies := tokens.GenerateToken(w, cred)
		if !okCookies {
			//w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Signin again.Cookies are not saved."))
			return
		}
		w.Write([]byte("successful Login."))
		return
	}

}
