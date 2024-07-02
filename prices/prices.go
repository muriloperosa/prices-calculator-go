package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/muriloperosa/prices-calculator-go/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could open the file!")
		fmt.Println(err)
		return
	}

	lines := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Could read the file content!")
		fmt.Println(err)
		defer file.Close()
		return
	}

	prices, err := conversion.StringsTofloat(&lines)

	if err != nil {
		fmt.Println("Converting price to float failed!")
		fmt.Println(err)
		defer file.Close()
		return
	}

	job.InputPrices = prices
	defer file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
