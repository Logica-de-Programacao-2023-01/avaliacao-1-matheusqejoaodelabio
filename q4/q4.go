package q4

import (
	"fmt"
)

func ClassifyPrices(prices []int) (int, error) {
	if prices == nil {
		return 0, fmt.Errorf("erro")
	}
	crescente := make([]int, len(prices))
	copy(crescente, prices)
	for i := 0; i < len(prices)-1; i++ {
		if crescente[i] > crescente[i+1] {
			break

		}
		if i == len(prices)-2 {
			return 1, nil
		}
	}

	decrescente := make([]int, len(prices))
	copy(decrescente, prices)
	for i := len(prices) - 1; i > 0; i-- {
		if decrescente[i] < decrescente[i-1] {
			break
		}
		return 2, nil
	}
	if crescente == nil && decrescente == nil {
		return 3, nil
	}
	if len(prices) == 1 {
		return 3, nil
	}
	return 0, nil
}
