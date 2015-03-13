package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Mnist struct {
	label string
	data  []int
}

func NewMnist(csvFile string) []Mnist {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rawCsvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal("ReadAll: ", err)
	}

	numMnistData := len(rawCsvData) - 1
	mnistData := make([]Mnist, numMnistData)

	for i, elem := range rawCsvData {
		if i == 0 {
			continue
		}
		mnistData[i-1].data = make([]int, len(elem[1:]))

		mnistData[i-1].label = elem[0]
		for j, pixel := range elem[1:] {
			mnistData[i-1].data[j], _ = strconv.Atoi(pixel)
		}
	}

	return mnistData
}

func main() {
	mnists := NewMnist("./train.csv")
	fmt.Printf("Num of Data: %d\n", len(mnists))
}