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
	dim   int
	data  []float64
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
		mnistData[i-1].data = make([]float64, len(elem[1:]))

		mnistData[i-1].label = elem[0]
		nDim := 0
		for j, pixel := range elem[1:] {
			mnistData[i-1].data[j], _ = strconv.ParseFloat(pixel, 64)
			nDim++
		}
		mnistData[i-1].dim = nDim
	}

	return mnistData
}

func main() {
	// Read MNIST Data
	mnist := NewMnist("./train.csv")
	fmt.Printf("Num of Data: %d\n", len(mnist))
	fmt.Printf("Num of Dims: %d\n", mnist[0].dim)
}
