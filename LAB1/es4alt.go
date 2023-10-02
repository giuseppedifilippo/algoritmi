package main

import "fmt"

func main() {
	lista := []int{4, 5, 1, 3, 6, 2} // input di test
	jump := 0
	diz := map[int]int{}
	for i, x := range lista {
		diz[x] = i
	}
	for i := 1; i < len(lista); i++ {
		if diz[i+1] < diz[i] {
			jump++
		}
	}
	fmt.Println(jump)
} // non funziona
//richiede gestione di numeri uguali
