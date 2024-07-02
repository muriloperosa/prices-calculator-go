package conversion

import (
	"strconv"
)

func StringsTofloat(strings *[]string) ([]float64, error) {

	floats := make([]float64, len(*strings))
	for idx, str := range *strings {
		floatVal, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return []float64{}, err
		}

		floats[idx] = floatVal
	}

	return floats, nil
}
