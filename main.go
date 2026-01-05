package main

import "fmt"

type Motor struct {
    Name  string
    Power int
}


// (r Motor) sebelum nama func Artinya fungsi 'SayHello' cuma bisa dipakai oleh struct Motor
func (r Motor) SayHello() {
    fmt.Printf("Motor saya %s, tenaganya %d CC.\n", r.Name, r.Power)
}

func (r Motor) CekStatus() {
	if r.Power > 300 {
		fmt.Printf("gila lu cepet juga lu bang bisa kecepatan %d CC.\n", r.Power)
	} else {
		fmt.Printf("ya standart sih kalo kecepatan luwh %d CC.\n", r.Power)
	}
}

func main() {
    bot := Motor{Name: "Beat 2024", Power: 120}
    
    bot.SayHello()
	bot.CekStatus()

}