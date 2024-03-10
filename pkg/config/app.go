package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open("mysql", "Ankit:Ankit@tcp(localhost:3306)/bookstoreDB?charset=utf8&parseTime=True&loc=Local")


    if err != nil {
        panic(err)
    }
}

func GetDB() *gorm.DB {
    return DB
}
