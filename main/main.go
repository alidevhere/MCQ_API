package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MAP for storing data after reading from file
var data = make(map[string]Mcq)

func CheckError(err error, str string) {
	if err != nil {
		log.Fatal(str, " : ", err.Error())
	}
}

func main() {
	//fmt.Println("Started")
	//loadData("./data.csv", data)
	//truncateData("./data.csv", "./updatedData.csv")
	//fmt.Println("Listening....")
	//runAPI()
	loadDataToDB("./data.csv")
}

func loadDataToDB(path string) {

	session, sessionErr := mgo.Dial("localhost")
	if sessionErr != nil {
		panic(sessionErr)
		return
	}
	defer session.Clone()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("McqAPI").C("MCQs")

	//input file
	file, err1 := os.Open(path)
	CheckError(err1, "in loadData func")
	//input file reader
	reader, err2 := csv.NewReader(file).ReadAll()
	CheckError(err2, "in Reading CSV file")

	for _, mcq := range reader {
		m := McqDB{
			Id:        bson.NewObjectId(),
			Statement: mcq[0],
			A:         mcq[1],
			B:         mcq[2],
			C:         mcq[3],
			D:         mcq[4],
			Answer:    mcq[5],
		}
		err := c.Insert(&m)
		if err != nil {
			fmt.Println(err)
		}
	}

	count, CountErr := c.Count()
	if CountErr != nil {
		fmt.Println(CountErr.Error())
	}

	fmt.Print(count)

}
