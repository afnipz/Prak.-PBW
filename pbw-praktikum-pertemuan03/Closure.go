package main

import "fmt"

func main() {
    // Closure untuk menentukan nilai huruf
    batasNilai := map[string]int{
        "A": 88,
        "B": 78,
        "C": 68,
        "D": 58,
    }

    getNilaiHuruf := func(nilaiUjian int) string {
        for huruf, batas := range batasNilai {
            if nilaiUjian >= batas {
                return huruf
            }
        }
        return "E"
    }

    // Menentukan nilai huruf untuk beberapa nilai ujian
    nilaiUjian := []int{85, 75, 95, 55, 65}
    for _, nilai := range nilaiUjian {
        fmt.Println("Nilai ujian", nilai, ":", getNilaiHuruf(nilai))
    }
}
