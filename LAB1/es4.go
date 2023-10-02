package main

import "fmt"

func main() {
	lista := []int{4, 5, 1, 3, 6, 2, 1} // input di test
	count := 1
	jump := 0
	var index1, index2 int
	for i := 0; i < len(lista); i++ { //itera lungo la lista e trova l indice che corrisponde al primo numero
		if lista[i] == count { //e trova ne salva l indice
			index1 = i
			count++
		}
		for j := 0; j < len(lista); j++ { //trova l indice del numero successivo
			if lista[j] == count {
				index2 = j
			}
		}
		if index2 < index1 { // confronto dei due indici
			jump++
		}
		//if count == len(lista) {
		//	break
		//}
	}
	fmt.Println(jump)
}
