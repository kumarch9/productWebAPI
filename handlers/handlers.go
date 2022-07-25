package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	cn "productwebapi/connection"
	md "productwebapi/model"
	tk "productwebapi/tokens"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var myDB *gorm.DB

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	readCook, validCook, expireCook, errInCook := tk.ValidateToken(r)
	fmt.Println(" validCook,expireCook, errInCook :: ", validCook, expireCook, errInCook)
	userhithost := r.Host
	userhiturl := r.URL.Path

	if !validCook && !expireCook {
		log.Println("Unauthorizd User. Accessed the path : ", userhithost, userhiturl)
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if expireCook && readCook != "" {
		log.Println("Session is expired or invalid token, userlogin try again. Accessed the path : ", userhithost, userhiturl)
		w.Write([]byte("Session is expired or invalid token, userlogin try again."))
		return
	} else {
		var MyProduct md.ProductInfo
		myDB = cn.DataMigration()
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("cache-control", "no-cache")
		w.Header().Add("Content-Type", "application/json")

		if errInDecode := json.NewDecoder(r.Body).Decode(&MyProduct); errInDecode != nil {
			log.Println("errInDecode", errInDecode)
			w.Write([]byte("Could not decoded the data."))
			return
		} else if MyProduct.ProductName == "" && MyProduct.ProductPrice == "" {
			w.Write([]byte("Information is incomplete"))
			return
		}

		if err := myDB.Create(&MyProduct).Error; err != nil {
			log.Fatalln("Err:", myDB.Error)
			w.Write([]byte("Not Created."))
			return
		}

		if err := json.NewEncoder(w).Encode(MyProduct); err != nil {
			log.Fatalln("err in Encoding:", err)
			w.Write([]byte("err in Encoding"))
			return
		}
		defer r.Body.Close()
		log.Println("Data has been save successful.")
		w.WriteHeader(http.StatusCreated)
		return
	}

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	readCook, validCook, expireCook, errInCook := tk.ValidateToken(r)
	fmt.Println(" validCook,expireCook, errInCook :: ", validCook, expireCook, errInCook)
	userhithost := r.Host
	userhiturl := r.URL.Path
	if !validCook && !expireCook {
		log.Println("Unauthorizd User. Accessed the path : ", userhithost, userhiturl)
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if expireCook && readCook != "" {
		log.Println("Session is expired or invalid token, userlogin try again.Accessed the path : ", userhithost, userhiturl)
		w.Write([]byte("Session is expired or invalid token, userlogin try again."))
		return
	} else {
		var MyProducts []md.ProductInfo
		myDB = cn.DataMigration()
		w.Header().Add("Content-Type", "application/json")
		if err := myDB.Find(&MyProducts).Error; err != nil {
			log.Fatalln("Err in find db :", err)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if errInEnCode := json.NewEncoder(w).Encode(MyProducts); errInEnCode != nil {
			//fmt.Fprintln(w, "Data not fatch ")
			log.Fatalln("err in Encoding:", errInEnCode)
			w.Write([]byte("Data could not data fatch."))
			return
		}
		defer r.Body.Close()
	}

}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	readCook, validCook, expireCook, errInCook := tk.ValidateToken(r)
	fmt.Println(" validCook,expireCook, errInCook :: ", validCook, expireCook, errInCook)
	userhithost := r.Host
	userhiturl := r.URL.Path
	if !validCook && !expireCook {
		log.Println("Unauthorizd User. Accessed the path : ", userhithost, userhiturl)
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if expireCook && readCook != "" {
		log.Println("Session is expired or invalid token, userlogin try again.Accessed the path : ", userhithost, userhiturl)
		w.Write([]byte("Session is expired or invalid token, userlogin try again."))
		return
	} else {
		var MyProduct md.ProductInfo
		myDB = cn.DataMigration()
		w.Header().Add("Content-Type", "application/json")
		if err := myDB.First(&MyProduct, mux.Vars(r)["id"]).Error; err != nil {
			log.Println("Err in db :", err)
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if errInEnCode := json.NewEncoder(w).Encode(MyProduct); errInEnCode != nil {
			log.Fatalln("err in Encoding:", errInEnCode)
			w.Write([]byte("Data could not data fatch."))
			return
		}
		defer r.Body.Close()
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	readCook, validCook, expireCook, errInCook := tk.ValidateToken(r)
	fmt.Println(" validCook,expireCook, errInCook :: ", validCook, expireCook, errInCook)
	userhithost := r.Host
	userhiturl := r.URL.Path
	if !validCook && !expireCook {
		log.Println("Unauthorizd User. Accessed the path : ", userhithost, userhiturl)
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if expireCook && readCook != "" {
		log.Println("Session is expired or invalid token, userlogin try again.Accessed the path : ", userhithost, userhiturl)
		w.Write([]byte("Session is expired or invalid token, userlogin try again."))
		return
	} else {
		var MyProduct md.ProductInfo
		myDB = cn.DataMigration()
		w.Header().Add("Content-Type", "application/json")

		if err := myDB.First(&MyProduct, mux.Vars(r)["id"]).Error; err != nil {
			log.Println("Err in db :", err)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if errInDecode := json.NewDecoder(r.Body).Decode(&MyProduct); errInDecode != nil {
			log.Println("errInDecode", errInDecode)
			w.Write([]byte("Error in encoding data."))
			return
		}
		myDB.Save(&MyProduct)
		if errInEnCode := json.NewEncoder(w).Encode(MyProduct); errInEnCode != nil {
			log.Fatalln("err in Encoding:", errInEnCode)
			w.Write([]byte("Error in encoding data."))
			return
		}
		defer r.Body.Close()
		w.Write([]byte("data has been updated."))
		return
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	readCook, validCook, expireCook, errInCook := tk.ValidateToken(r)
	fmt.Println(" validCook,expireCook, errInCook :: ", validCook, expireCook, errInCook)
	userhithost := r.Host
	userhiturl := r.URL.Path
	if !validCook && !expireCook {
		log.Println("Unauthorizd User. Accessed the path : ", userhithost, userhiturl)
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if expireCook && readCook != "" {
		log.Println("Session is expired or invalid token, userlogin try again. Accessed the path : ", userhithost, userhiturl)
		w.Write([]byte("Session is expired or invalid token, userlogin try again."))
		return
	} else {
		var MyProduct md.ProductInfo
		myDB = cn.DataMigration()
		w.Header().Add("Content-Type", "application/json")
		if errDel := myDB.Delete(&MyProduct, mux.Vars(r)["id"]).Error; errDel != nil {
			log.Println("Err in deletion:", errDel)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if errInEncode := json.NewEncoder(w).Encode(("data deleted !!")); errInEncode != nil {
			log.Println("err in Encoding:", errInEncode)
			w.Write([]byte("Error in encoding data."))
			return
		}
		defer r.Body.Close()
		w.WriteHeader(http.StatusOK)
		return
	}

}
