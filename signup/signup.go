package signup

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	cn "productwebapi/connection"
	hash "productwebapi/hashkey"

	"gorm.io/gorm"
)

var myDB *gorm.DB
var userCredential hash.Credentials

func RegUserHandler(w http.ResponseWriter, r *http.Request) {
	myDB = cn.RegdMigration()
	//log.Println("Server started.")
	r.Header.Add("Content-Type", "application/json")
	w.Header().Add("Content-Type", "application/json")

	if errDecode := json.NewDecoder(r.Body).Decode(&userCredential); errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("userCredential : ", userCredential)
	hashPsw, hashErr := userCredential.HashPassword()
	fmt.Println("hashPsw : ", hashPsw)
	if hashErr != nil {
		log.Fatalln("Err:", myDB.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if dbErr := myDB.Create(&userCredential).Error; dbErr != nil {
		log.Println("dbErr: ", dbErr.Error())
		w.Write([]byte("Could not signup, Try again."))
		return
	}

	w.WriteHeader(http.StatusCreated)
	defer r.Body.Close()
}
