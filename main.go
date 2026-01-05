package main

import "fmt"

func getBiodata() (string, int) {
    panggilan := "King"
    semester := 4
    return panggilan, semester
}

func main() {
    panggilan, sem := getBiodata()
	// %s	String (Teks)	
	// %d	Decimal (Angka bulat)	
	// %f	Float (Angka desimal)	
	// %t	True/False (Boolean)	
	// %v	Value (Apa saja)	
    fmt.Printf("Halo %s, sekarang sudah semester %d ya?\n", panggilan, sem)
    fmt.Println("Halo", panggilan, "sekarang sudah semester", sem, "ya?")

    _ , semesterSaja := getBiodata()
    fmt.Println("Saya Semester:", semesterSaja)
}