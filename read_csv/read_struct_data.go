package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord 一条CSV记录
type CSVRecord struct {
	Id          int
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error // 当前一行可能存在的解析错误
}

func ReadStructData() {
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	var csvData []CSVRecord

	line := 1

	for {
		record, err := reader.Read()
		// 文件读取完毕
		if err == io.EOF {
			break
		}

		var csvRecord CSVRecord

		for idx, value := range record {
			// idx == 5 是最后一列，品种描述
			if idx == 5 {
				// 读出了空的品种信息
				if value == "" {
					log.Printf("解析第 %d 行失败，无法预知类型，列 %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Empty species value")
					break
				}

				// 读到最后一列，退出循环
				csvRecord.Species = value
				break
				// 读取第0列，ID值
			} else if idx == 0 {
				ID, err := strconv.Atoi(value)
				if err != nil {
					log.Printf("解析第 %d 行失败，ID解析错误，列 %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Invalid ID")
					// ID解析出错，跳出本次循环，不再解析
					break
				}
				csvRecord.Id = ID
				// 处理其他列
			} else {
				var floatValue float64
				if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
					log.Printf("解析第 %d 行失败，无法预知类型，列 %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Could not parse float")
					break
				}

				switch idx {
				case 1:
					csvRecord.SepalLength = floatValue
				case 2:
					csvRecord.SepalWidth = floatValue
				case 3:
					csvRecord.PetalLength = floatValue
				case 4:
					csvRecord.PetalWidth = floatValue
				}
			}

		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
		line++
	}

	fmt.Printf("Successfully parsed %d lines\n", len(csvData))
	fmt.Println(csvData[:2])
}
