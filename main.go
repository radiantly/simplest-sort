package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number of elements you'd like to sort: ")
	fmt.Scan(&n)

	Ar := make([]int, n)

	fmt.Println("Enter the elements: ")
	for i := range Ar {
		fmt.Scan(&Ar[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if Ar[i] < Ar[j] {
				Ar[i], Ar[j] = Ar[j], Ar[i]
			}
		}
	}

	fmt.Println(Ar)
}
