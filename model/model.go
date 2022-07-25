package model

import "gorm.io/gorm"

//information of product in save struct type formate
type ProductInfo struct {
	gorm.Model
	ProductName     string `bson:"productName" json:"productName"`
	ProductType     string `bson:"productType" json:"productType"`
	ProductMaterial string `bson:"productMaterial" json:"productMaterial"`
	ProductSize     string `bson:"productSize" json:"productSize"`
	ProductColor    string `bson:"productColor" json:"productColor"`
	ProductPrice    string `bson:"productPrice" json:"productPrice"`
}

//credential get information of user those will handle the products as like edit
type Credentials struct {
	gorm.Model
	UserName  string `json:"username"`
	UserEmail string `json:"email"`
	Password  string `json:"password" `
}
