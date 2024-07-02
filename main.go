package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 1.0, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {

		taxIncludedPrices := make([]float64, len(prices))

		for priceI, price := range prices {
			taxIncludedPrices[priceI] = price * (1 + taxRate)
		}

		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}
