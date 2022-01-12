package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MCQ struct {
	gorm.Model
	Statement     string
	CorrectOption string
	A             string
	B             string
	C             string
	D             string
}

func (m *MCQ) setup_with_MySql(data_string string, path string) {
	db := initializeMySQL(data_string)
	db.AutoMigrate(&MCQ{})
	loadDataMySQL(db, path)
}

func initializeMySQL(data_string string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(data_string), &gorm.Config{})
	//fmt.Println("DB: ", db)
	if err != nil {
		fmt.Print("ERROR: ", err)
		panic(err)
	}
	return db
}

func loadDataMySQL(db *gorm.DB, csv_path string) {
	var data []MCQ
	file, err1 := os.Open(csv_path)
	CheckError(err1, "in loadData func")
	reader, err2 := csv.NewReader(file).ReadAll()
	CheckError(err2, "in Reading CSV file")

	for _, mcq := range reader {
		m := MCQ{
			Statement:     mcq[0],
			A:             mcq[1],
			B:             mcq[2],
			C:             mcq[3],
			D:             mcq[4],
			CorrectOption: mcq[5],
		}
		data = append(data, m)
		//fmt.Println(m)
	}
	//fmt.Println("Data length=", len(data))

	tx := db.Create(&data)
	fmt.Println("Data loaded into MySql database")
	fmt.Println("Row Affected = ", tx.RowsAffected)
	//fmt.Println("Row Affected = ", tx.RowsAffected)

}

func (m *MCQ) setup_with_MngoDB(path string) {

	panic("Not Implemented")
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
