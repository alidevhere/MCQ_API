package main

import (
	"fmt"
	"log"
)

//MAP for storing data after reading from file
var data = make(map[string]Mcq)

func CheckError(err error, str string) {
	if err != nil {
		log.Fatal(str, " : ", err.Error())
	}
}

func main() {
	fmt.Println("Started")
	loadData("./data.csv", data)
	//truncateData("./data.csv", "./updatedData.csv")
	fmt.Println("Listening....")
	runAPI()
}
