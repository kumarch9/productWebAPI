package hashkey

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Credentials struct {
	gorm.Model
	UserName  string `json:"username"`
	UserEmail string `json:"email"`
	Password  string `json:"password" `
}

func (cred *Credentials) HashPassword() (hashPassword string, errInHashPsw error) {
	local_time := time.Now().Add(time.Hour * 5) //to convert in IST time zone
	cred.Model.CreatedAt = local_time.Add(time.Minute * 30)
	cred.Model.UpdatedAt = local_time.Add(time.Minute * 30)
	bytes, err := bcrypt.GenerateFromPassword([]byte(cred.Password), 14)
	if err != nil {
		return "", err
	}
	cred.Password = string(bytes)
	fmt.Println("HasshPsw :", string(bytes))
	return string(bytes), nil
}

func (cred *Credentials) CheckPassword(password string) (Match bool) {
	err := bcrypt.CompareHashAndPassword([]byte(cred.Password), []byte(password))

	return err == nil
}
