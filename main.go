package main

import "fmt"

func sapaMahasiswa(nama string, jurusan string) {
    fmt.Printf("Halo %s, semangat ngodingnya di jurusan %s!\n", nama, jurusan)
}

func main() {
    sapaMahasiswa("Pranata", "Teknologi Informasi")
}