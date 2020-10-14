package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func ReadAndFilterDF() {
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)

	filter := dataframe.F{
		Colname:    "Species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"SepalWidthCm", "Species"})
	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"SepalWidthCm", "Species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)
}
