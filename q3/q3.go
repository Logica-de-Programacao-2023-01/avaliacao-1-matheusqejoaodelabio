package q3

import "fmt"

func DominoPieces(m, n int) (int, error) {
	if m <= 0 || n <= 0 {
		return 0, fmt.Errorf("erro")
	}
	
	return (m*n)/2, nil
}
