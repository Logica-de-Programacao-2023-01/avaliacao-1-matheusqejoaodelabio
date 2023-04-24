package bonus

import "sort"

func CalculateTowers(barLengths []int) (int, int) {
	sort.Ints(barLengths)
	unicobarra := make(map[int]bool)
	for _, length := range barLengths {
		unicobarra[length] = true
	}
		caontador:= make(map[int]int)
		for _, length := range barLengths {
			caontador[length]++
		}
		numerotorre:=len(unicobarra)
		var maxaltura int
		for _, contar:= range caontador {
			if contar > maxaltura {
				maxaltura = contar
			}
		}
	return maxaltura, numerotorre
}
