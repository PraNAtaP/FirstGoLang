package main

import "fmt"

// Struct adalah kumpulan dari berbagai tipe data yang dibungkus jadi satu kesatuan
type Mahasiswa struct {
    Nama    string
    NIM     string
    Jurusan string
    IPK     float64
}

func main() {
	mhs1 := Mahasiswa{
		Nama:    "Pranata Putrandana",
		NIM:     "244107060114",
		Jurusan: "Teknologi Informasi",
		IPK:     3.90,
	}
	fmt.Printf("Mahasiswa: %s (%s)\n", mhs1.Nama, mhs1.NIM)
	fmt.Printf("Jurusan: %s (IPK: %.2f)\n", mhs1.Jurusan, mhs1.IPK)
}