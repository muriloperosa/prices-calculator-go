package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER (0 to exit)")

	prices := []string{}

	for {
		var price string
		fmt.Print("Prices: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
