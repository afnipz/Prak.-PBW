package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

	//Input Nama User
    fmt.Print("Masukkan Nama Anda: ")
    scanner.Scan()
    Nama := scanner.Text()

	//Input Usia User
    fmt.Print("Masukkan Usia Anda: ")
    scanner.Scan()
    var Usia int
	fmt.Sscanf(scanner.Text(), "%d", &Usia)

    var Kategori string
    if Usia < 18 {
        Kategori = "Anak-anak"
    } else {
        Kategori = "Dewasa"
    }

    fmt.Printf("Selamat datang, %s Anda termasuk kategori %s.\n", Nama, Kategori)
}
