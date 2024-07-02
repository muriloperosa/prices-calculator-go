package main

import (
	"fmt"

	"github.com/muriloperosa/prices-calculator-go/filemanager"
	"github.com/muriloperosa/prices-calculator-go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.0, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		priceJob.Process()
	}
}
