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

func truncateData(inpath string, outpath string) {

	//input file
	file, err1 := os.Open(inpath)
	CheckError(err1, "in loadData func")
	//input file reader
	reader, err2 := csv.NewReader(file).ReadAll()
	CheckError(err2, "in Reading CSV file")
	//output file
	outFile, errOutFile := os.Create(outpath)
	CheckError(errOutFile, "in Reading CSV file")
	//output file writer
	csvWriter := csv.NewWriter(outFile)
	var data [][]string
	for index, mcq := range reader {

		if !isEmpty(mcq) {
			println(index, mcq[0], mcq[1], mcq[2], mcq[3], mcq[4], mcq[5])
			data = append(data, mcq)
		}
	}
	//fmt.Print(m)

	csvWriter.WriteAll(data)
}

func isEmpty(mcq []string) bool {

	if mcq[0] != "" && mcq[1] != "" && mcq[2] != "" && mcq[3] != "" && mcq[4] != "" && mcq[5] != "" {
		return false
	}
	return true
}
