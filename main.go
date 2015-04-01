package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"math"

	"github.com/skelterjohn/go.matrix"
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

func (src Mnist) distTo(dst Mnist) (dist float64, err error) {
	dist = 0.0
	if src.dim != dst.dim {
		return dist, errors.New("Mismatched: src.dim != dst.dim")
	}

	m1 := matrix.MakeDenseMatrix(src.data, 1, src.dim)
	m2 := matrix.MakeDenseMatrix(dst.data, 1, dst.dim)
	m3, _ := m1.Times(m2.Transpose())

	dist = math.Sqrt(m3.Get(0, 0))

	return dist, nil
}

func main() {
	// Read MNIST Data
	mnist := NewMnist("./train.csv")
	fmt.Printf("Num of Data: %d\n", len(mnist))
	fmt.Printf("Num of Dims: %d\n", mnist[0].dim)

	// Calculate distance
	for i, _ := range mnist {
		if i == len(mnist)-1 {
			break
		}

		dist, err := mnist[i].distTo(mnist[i+1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("dist(mnist[%d], mnist[%d]) = %f\n", i, i+1, dist)
	}
}
