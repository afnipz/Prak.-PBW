package main

import "fmt"

func BubbleSort(x []int) {
	n := len(x)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	fmt.Println("Bubble Sort")

	arrayNumber := [...]int{30, 17, 26, 1, 27, 101, 73, 34, 68, 93, 91, 12, 45, 20, 83, 55, 72, 33, 61, 38}

	fmt.Println("Sebelum dilakukan pengurutan:")
	fmt.Println(arrayNumber)

	
	BubbleSort(arrayNumber[:])

	fmt.Println("Setelah dilakukan pengurutan:")
	fmt.Println(arrayNumber)
}
