package main

import (
	con "productwebapi/connection"
	rt "productwebapi/routing"
)

func main() {
	con.DataMigration()
	rt.HandlerRouting()
}

//GORM  : Go  relational Mapping that it will auto create tables in DB
//go get -u github.com/gorilla/mux
//go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql
