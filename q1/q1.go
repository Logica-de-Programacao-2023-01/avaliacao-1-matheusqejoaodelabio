package q1

import "fmt"

func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, fmt.Errorf("erro")
	}
	par := false
	if weight%2 == 0 {
		par = true
		return true, nil
	}
	if par {
		return false, nil
	}

	return false, nil
}
