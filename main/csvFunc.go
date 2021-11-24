package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func loadData(path string, m map[string]Mcq) {

	file, err1 := os.Open(path)
	CheckError(err1, "in loadData func")
	reader, err2 := csv.NewReader(file).ReadAll()
	CheckError(err2, "in Reading CSV file")

	for index, mcq := range reader {
		//fmt.Println(index, mcq[0])
		m[strconv.Itoa(index)] = Mcq{
			Id:        uint(index),
			Statement: mcq[0],
			A:         mcq[1],
			B:         mcq[2],
			C:         mcq[3],
			D:         mcq[4],
			Answer:    mcq[5],
		}
	}
	//fmt.Print(m)
}
