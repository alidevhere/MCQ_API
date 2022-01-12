package main

//	TODO: Apply pagination
// TODO : Return internal error on failing Marshal

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ApiMcq struct {
	Statement     string `json:"statement"`
	CorrectOption string `json:"CorrectOption"`
	A             string `json:"A"`
	B             string `json:"B"`
	C             string `json:"C"`
	D             string `json:"D"`
}

// Gets all records
//QUERY:  SELECT * FROM mcqs;
func getAllMcqs(w http.ResponseWriter, r *http.Request) {
	// METHOD 1
	var mcqs []ApiMcq
	DB.Model(&MCQ{}).Find(&mcqs)
	print(len(mcqs))

	// METHOD 2
	// By  RAW SQL
	//https://gorm.io/docs/sql_builder.html

	w.Header().Set("Content-Type", "application/json")
	output, err := json.Marshal(mcqs)
	CheckError(err, " In func getAllMcqs")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

}

// TODO : internal Error Response
func getRandomMcq(w http.ResponseWriter, r *http.Request) {
	var randomMcqs []Mcq
	limit, _ := strconv.Atoi(mux.Vars(r)["limit"])
	var iteration int = 0
	rand.Seed(time.Now().UnixNano())
	var doneList = make(map[string]bool)

	for len(randomMcqs) < limit {
		iteration++
		k := strconv.Itoa(rand.Intn(iteration))

		value, valid_key := data[k]
		_, repeated_key := doneList[k]

		if valid_key && !repeated_key {
			randomMcqs = append(randomMcqs, value)
			doneList[k] = true
			fmt.Println(k, " accepted")

		}

	}
	w.Header().Set("content-type", "application/json")
	out, err := json.Marshal(randomMcqs)
	CheckError(err, "in getRandomMcq")
	w.Write(out)
	fmt.Println("completed ", len(randomMcqs), " in ", iteration, " iterations")

}

func getMcqById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var mcq ApiMcq
	result := DB.Model(&MCQ{}).First(&mcq, id)

	fmt.Println(result.RowsAffected)
	if result.RowsAffected == 1 {
		out, err := json.Marshal(mcq)
		CheckError(err, " in func getMcqByID")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err.Error())
		}

		//w.Header().Set("content-type", "application/json")
		w.Header().Set(ContentType.ContentType())
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	} else {
		w.Header().Set(ContentType.ContentType())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"error\":\"Id Not Found\"}"))
	}
}

func UpdateMcqByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var newMcq MCQ
	var OldMcq MCQ

	result := DB.Model(&MCQ{}).First(&OldMcq, id)

	if result.RowsAffected == 1 {
		err := json.NewDecoder(r.Body).Decode(&newMcq)
		//fmt.Print(r.Body)
		CheckError(err, " In UpdateMcqById function")
		//fmt.Print(updatedMcq)
		// ID updating should not be permissible
		DB.Model(&OldMcq).Updates(&newMcq)
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusNoContent)

	}

	//	w.WriteHeader(http.StatusNotFound)

}

func DeleteMcqById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	result := DB.Delete(&MCQ{}, id)
	if result.RowsAffected == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func CreateMcq(w http.ResponseWriter, r *http.Request) {
	var mcq MCQ
	err := json.NewDecoder(r.Body).Decode(&mcq)
	CheckError(err, " inside CreateMcq")
	result := DB.Create(&mcq)
	fmt.Println(mcq)

	if result.RowsAffected == 1 {
		j, err := json.Marshal(&mcq)
		CheckError(err, " inside CreateMcq: Marshelling failed")
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
	}

}

func SearchByStmt(w http.ResponseWriter, r *http.Request) {
	stmt := mux.Vars(r)["stmt"]
	var mcqs []MCQ
	//println("Statement===", stmt)
	DB.Where("statement LIKE ?", "%"+stmt+"%").Find(&mcqs)
	if len(mcqs) > 0 {
		output, err := json.Marshal(mcqs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
