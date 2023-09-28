package main

import "fmt"

func main() {
	lista := []int{4, 5, 1, 3, 6, 2, 1}
	count := 1
	jump := 0
	var index1, index2 int
	for i := 0; i < len(lista); i++ {
		if lista[i] == count {
			index1 = i
			count++
		}
		for j := 0; j < len(lista); j++ {
			if lista[j] == count {
				index2 = j
			}
		}
		if index2 < index1 {
			jump++
		}
		//if count == len(lista) {
		//	break
		//}
	}
	fmt.Println(jump)
}
