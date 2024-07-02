package main

import (
	"github.com/muriloperosa/prices-calculator-go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.0, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
