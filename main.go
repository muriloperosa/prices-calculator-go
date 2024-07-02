package main

import (
	"fmt"

	"github.com/muriloperosa/prices-calculator-go/filemanager"
	"github.com/muriloperosa/prices-calculator-go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.0, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {

		doneChannels[index] = make(chan bool)
		errChannels[index] = make(chan error)

		// file io
		fm := filemanager.New("storage/input/prices.txt", fmt.Sprintf("storage/output/result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)

		// cmd io
		// cmd := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(*cmd, taxRate)

		go priceJob.Process(doneChannels[index], errChannels[index])

		// if err != nil {
		// 	fmt.Println("Could not process job!")
		// 	fmt.Println(err)
		// }
	}

	for index, _ := range taxRates {
		select {
		case err := <-errChannels[index]:
			if err != nil {
				fmt.Println("Could not process job!")
				fmt.Println(err)
			}
		case <-doneChannels[index]:
			fmt.Println("Done!")
		}
	}
}
