package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func runAPI() {
	//Router
	r := mux.NewRouter().StrictSlash(true)
	fmt.Println("Started...")

	//=========	ROUTES  ==========

	r.HandleFunc("/mcqs", getAllMcqs).Methods("GET")
	r.HandleFunc("/rand-mcq/{limit}", getRandomMcq).Methods("GET")
	r.HandleFunc("/mcq/{id}", getMcqById).Methods("GET")
	r.HandleFunc("/mcq/{id}", UpdateMcqByID).Methods("PUT")
	r.HandleFunc("/mcq", CreateMcq).Methods("POST")
	r.HandleFunc("/mcq/{id}", DeleteMcqById).Methods("DELETE")
	r.HandleFunc("/mcqs/search/{stmt}", SearchByStmt).Methods("GET")

	//===========

	//Server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Listening....")
	log.Fatal(server.ListenAndServe())
	//fmt.Println("Accepted...")

	//	truncateData("./data.csv", "ds")
}
