package main

func runAPI() {
	/*
		//Router
		r := mux.NewRouter().StrictSlash(true)
		//fmt.Println("Started...")

		//=========	ROUTES  ==========

		r.HandleFunc("/mcqs", getAllMcqs).Methods("GET")
		r.HandleFunc("/rand-mcq/{limit}", getRandomMcq).Methods("GET")
		r.HandleFunc("/mcq/{id}", getMcqById).Methods("GET")
		r.HandleFunc("/mcq/{id}", UpdateMcqByID).Methods("PUT")
		r.HandleFunc("/mcq", CreateMcq).Methods("POST")
		r.HandleFunc("/mcq/{id}", DeleteMcqById).Methods("DELETE")

		//===========

		//Server
		server := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}

		//fmt.Println("Listening....")
		server.ListenAndServe()
		fmt.Println("Accepted...")
	*/

	truncateData("./data.csv", "ds")
}
