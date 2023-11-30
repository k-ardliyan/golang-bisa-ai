package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func LuasPersegi(sisi float64) float64 {
	return sisi * sisi
}

func HitungLuasPersegi()  {
	fmt.Println("\n===Aplikasi menghitung luas persegi===")
	fmt.Println("Luas Persegi = Sisi * Sisi")
	fmt.Print("Masukkan sisi persegi (cm): ")
	sisi, _ := lib.GetInput()
	result := LuasPersegi(sisi)
	fmt.Printf("Luas Persegi adalah %.2f cm2\n", result)
}