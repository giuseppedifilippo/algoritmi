package main

import "fmt"

func main() {
	var stringa string
	var a, count int
	stringa = "aaaccccbb"
	for _, x := range stringa {
		if string(x) == "a" {
			a++
		}
		if string(x) == "b" {
			count += a
		}
	}
	fmt.Println(count)
}
