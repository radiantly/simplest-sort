package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func colorText(text string, color string) string {
	switch color {
	case "red":
		return "\033[31m" + text + "\033[39m\033[49m"
	case "green":
		return "\033[32m" + text + "\033[39m\033[49m"
	case "yellow":
		return "\033[33m" + text + "\033[39m\033[49m"
	default:
		return text
	}
}

func main() {
	var n int
	fmt.Print(colorText("Enter the number of elements you'd like to sort: ", "yellow"))
	fmt.Scan(&n)

	Ar := make([]int, n)
	positions := make([]int, n)

	recomputePositions := func() {
		for i := range Ar {
			positions[i] = int(math.Log10(float64(Ar[i]))) + 1

			if i > 0 {
				positions[i] += positions[i-1] + 1
			}
		}
	}

	printCurrent := func(i, j int, color string) {
		// Line 1
		iSpaces := strings.Repeat(" ", positions[i]-1)
		fmt.Println("\r\033[K"+colorText("i ", "yellow"), iSpaces, "▼")

		// Line 2
		fmt.Print("\r   [")
		for k := range Ar {
			num := strconv.Itoa(Ar[k])
			if i == k || j == k {
				fmt.Print(colorText(num, color))
			} else {
				fmt.Print(num)
			}
			if k != n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")

		// Line 3
		jSpaces := strings.Repeat(" ", positions[j]-1)
		fmt.Println("\r\033[K"+colorText(" j", "yellow"), jSpaces, "▲\033[3A")
	}

	fmt.Print(colorText("Enter elements: ", "yellow"))
	for i := range Ar {
		fmt.Scan(&Ar[i])
	}

	fmt.Println()
	recomputePositions()
	for i := 0; i < n; i++ {
		fmt.Print("\033[<N>A")
		for j := 0; j < n; j++ {
			if Ar[i] < Ar[j] {
				printCurrent(i, j, "green")
				time.Sleep(time.Second)

				Ar[i], Ar[j] = Ar[j], Ar[i]
				recomputePositions()

				printCurrent(i, j, "green")
			} else {
				printCurrent(i, j, "red")
			}
			time.Sleep(700 * time.Millisecond)
		}
	}

	fmt.Print("\r\033[K\r")
	fmt.Println(Ar)
	fmt.Println("\r\033[K\r" + colorText("Sorted! 🎉", "green"))
}
