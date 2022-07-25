package connection

import (
	"log"

	md "productwebapi/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const conStr string = "root:12345@tcp(127.0.0.1:3306)/productdbs?"

//auto create database in server according the given data model by help of gorm driver
func DataMigration() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conStr+"parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal("error in connection : ", err)
	}
	log.Println("server is started.")
	db.AutoMigrate(&md.ProductInfo{})
	return db
}

//auto create database in server according the given data model by help of gorm driver
func RegdMigration() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conStr+"parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal("error in connection : ", err)
	}
	db.AutoMigrate(&md.Credentials{})
	log.Println("Server Started.")
	return db
}
