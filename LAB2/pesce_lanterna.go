package main

import "fmt"

func main() {
	pesci := map[int]int{8: 0, 7: 0, 6: 0, 5: 0, 4: 1, 3: 2, 2: 1, 1: 1, 0: 0}
	//gestione dell input
	/*lista := []int{3, 4, 3, 1, 2}
	for _, x := range lista {
		pesci[x]++
	}*/
	temp := 0
	for i := 0; i < 256; i++ { //esegue le operazioni successive tante quanti sono i giorni
		temp = pesci[0]
		for j := 8; j >= 0; j-- { //riduce il timer dei pesci
			pesci[j], temp = temp, pesci[j]
		}

		pesci[6] += temp
		fmt.Println(pesci)
	}
	count := 0
	for i := 0; i < 9; i++ {
		count += pesci[i]
	}
	fmt.Println(count)
}
