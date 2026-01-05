package main

import "fmt"

func hitungPersegi(sisi int) (luas int, keliling int) {
    luas = sisi * sisi
    keliling = 4 * sisi
    return 
}

func main() {
    l, k := hitungPersegi(10)
    fmt.Printf("Luas: %d, Keliling: %d\n", l, k)
}