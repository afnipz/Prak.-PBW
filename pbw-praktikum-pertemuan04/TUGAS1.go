package main

import "fmt"

func cari(n int) int {
	if n == 1 {
		return 1
	}
	return n * cari(n-1)
}

func main() {
	// Fungsi diatas sama saja dengan
	// 7 x (6 x(5 x(4 x(3 x (2 x (1 x 1))))))
	fmt.Println(cari(7))
}
