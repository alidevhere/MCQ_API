package main

/*
1- load data from csv
2- create models in required database


*/

import (
	"log"

	"gorm.io/gorm"
)

func CheckError(err error, str string) {
	if err != nil {
		log.Fatal(str, " : ", err.Error())
	}
}

var DB *gorm.DB
var data = make(map[string]Mcq)

func main() {
	dsn := "rana:ali123@tcp(localhost:3306)/MCQ_API?charset=utf8mb4&parseTime=True&loc=Local"
	//var m MCQ
	//m.setup_with_MySql(dsn, "./data.csv")
	DB = initializeMySQL(dsn)
	runAPI()

}
