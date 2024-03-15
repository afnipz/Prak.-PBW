package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	// Inisialisasi map untuk menyimpan data mahasiswa
	dataMahasiswa := make(map[string]Mahasiswa)

	// Menambahkan data mahasiswa ke dalam map
	dataMahasiswa["4522210120"] = Mahasiswa{"Afni Puspita Zahra", "4522210120", "Teknik Informatika"}
	dataMahasiswa["462220001"] = Mahasiswa{"Daffa Putra", "462220001", "Teknik Mesin"}
    dataMahasiswa["1389274636"] = Mahasiswa{"Reza Daffa", "1389274636", "Akuntansi"}
    dataMahasiswa["1389273641"] = Mahasiswa{"Sri Hartatik", "1389273641", "Akuntansi"}
    dataMahasiswa["4422214503"] = Mahasiswa{"Satrio Adi", "4422214503", "Teknik Elektro"}
    dataMahasiswa["4422214506"] = Mahasiswa{"Kevin Aliya", "4422214506", "Teknik Elektro"}
    dataMahasiswa["300404261"] = Mahasiswa{"Devani Putri", "300404261", "Manajemen"}

	// Menampilkan daftar nama mahasiswa dengan perulangan
	fmt.Println("Daftar Mahasiswa:")
	for _, mahasiswa := range dataMahasiswa {
		fmt.Println("Nama:", mahasiswa.Nama)
	}

	// Mencari data mahasiswa berdasarkan NPM
	MencariDataNPM := "4522210120"
	mahasiswa, found := dataMahasiswa[MencariDataNPM]
	if found {
		fmt.Println("\nData Mahasiswa dengan NPM", MencariDataNPM, ":")
		fmt.Println("Nama   :", mahasiswa.Nama)
		fmt.Println("NPM    :", mahasiswa.NPM)
		fmt.Println("Jurusan:", mahasiswa.Jurusan)
	} else {
		fmt.Println("\nData Mahasiswa dengan NPM", MencariDataNPM, "tidak ditemukan.")
	}
}
