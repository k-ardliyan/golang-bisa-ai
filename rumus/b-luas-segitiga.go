package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func LuasSegitiga(alas, tinggi float64) float64 {
	return alas * tinggi / 2
}

func HitungLuasSegitiga() {
	fmt.Println("\n===Aplikasi menghitung luas segitiga===")
	fmt.Println("Luas Segitiga = Alas * Tinggi / 2")
	fmt.Print("Masukkan alas segitiga (cm): ")
	alas, _ := lib.GetInput()
	fmt.Print("Masukkan tinggi segitiga (cm): ")
	tinggi, _ := lib.GetInput()
	result := LuasSegitiga(alas, tinggi)
	fmt.Printf("Luas Segitiga adalah %.2f cm2\n", result)
	
}