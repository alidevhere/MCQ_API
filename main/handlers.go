package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//	TODO: Apply pagination
// TODO : Return internal error on failing Marshal

func getAllMcqs(w http.ResponseWriter, r *http.Request) {
	var mcqs []Mcq

	for _, mcq := range data {
		mcqs = append(mcqs, mcq)
	}

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
	if mcq, ok := data[id]; ok {
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
	}
	fmt.Println("ID ", id)
	fmt.Println(data[id])

}

func UpdateMcqByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedMcq Mcq
	if _, ok := data[id]; ok {
		err := json.NewDecoder(r.Body).Decode(&updatedMcq)
		//fmt.Print(r.Body)
		CheckError(err, " In UpdateMcqById function")
		//fmt.Print(updatedMcq)
		i, _ := strconv.ParseUint(id, 10, 64)
		updatedMcq.Id = uint(i)
		data[id] = updatedMcq
		w.WriteHeader(http.StatusNoContent)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func DeleteMcqById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if _, ok := data[id]; ok {
		delete(data, id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func CreateMcq(w http.ResponseWriter, r *http.Request) {
	var mcq Mcq
	err := json.NewDecoder(r.Body).Decode(&mcq)
	CheckError(err, " inside CreateMcq")
	mcq.Id = uint(len(data) + 1)
	data[strconv.Itoa(len(data)+1)] = mcq
	j, err := json.Marshal(&mcq)
	CheckError(err, " inside CreateMcq: Marshelling failed")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
