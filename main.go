package main

import (
	"fmt"

	"github.com/muriloperosa/prices-calculator-go/filemanager"
	"github.com/muriloperosa/prices-calculator-go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.0, 0.15}

	for _, taxRate := range taxRates {
		// file io
		fm := filemanager.New("storage/input/prices.txt", fmt.Sprintf("storage/output/result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)

		// cmd io
		// cmd := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(*cmd, taxRate)

		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job!")
			fmt.Println(err)
		}
	}
}
