package main

import (
	"fmt"
	"io"
	"os"

	"log"

	"encoding/csv"
)

func ReadDirtyData() {
	f, err := os.Open("iris_unexpected.csv")
	if err != nil {
		log.Fatal(err)
	}

	// 最终注意关闭文件
	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = 6
	var rawCSVData [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
			continue
		}

		rawCSVData = append(rawCSVData, record)
	}

	fmt.Printf("Parsed %d lines successfully\n", len(rawCSVData))
	fmt.Println(rawCSVData[:3])
}
