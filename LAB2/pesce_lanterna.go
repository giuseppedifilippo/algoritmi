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
	temp2 := 0
	for i := 0; i < 80; i++ {
		temp = pesci[0]
		for j := 8; j > 0; j-- {
			pesci[j-1] = pesci[j]
		}
		//pesci[8], pesci[7], pesci[6], pesci[5], pesci[4], pesci[3], pesci[2], pesci[1], pesci[0] = pesci[0], pesci[8], pesci[7], pesci[6], pesci[5], pesci[4], pesci[3], pesci[2], pesci[1]
		pesci[8] += temp
		pesci[6] += temp
		fmt.Println(pesci)
	}
	count := 0
	for i := 0; i < 9; i++ {
		count += pesci[i]
	}
	fmt.Println(count)
}
