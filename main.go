package main

import (
	"github.com/muriloperosa/prices-calculator-go/cmdmanager"
	"github.com/muriloperosa/prices-calculator-go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.0, 0.15}

	for _, taxRate := range taxRates {
		// file io
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)

		// cmd io
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(*cmd, taxRate)

		priceJob.Process()
	}
}
