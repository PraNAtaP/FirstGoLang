package main

import "fmt"

func getBiodata() (string, int) {
    panggilan := "King"
    semester := 4
    return panggilan, semester
}

func main() {
    panggilan, sem := getBiodata()
    fmt.Printf("Halo %s, sekarang sudah semester %d ya?\n", panggilan, sem)

    _ , semesterSaja := getBiodata()
    fmt.Println("Saya Semester:", semesterSaja)
}