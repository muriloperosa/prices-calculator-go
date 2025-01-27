package prices

import (
	"fmt"

	"github.com/muriloperosa/prices-calculator-go/conversion"
	"github.com/muriloperosa/prices-calculator-go/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsTofloat(&lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChannel chan bool, errChannel chan error) {
	err := job.LoadData()

	if err != nil {
		errChannel <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result

	err = job.IOManager.WriteResult(job)
	if err != nil {
		errChannel <- err
		return
	}

	doneChannel <- true
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
