package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func MassaJenis(massa float64, volume float64) float64 {
	return massa / volume
}

func HitungMassaJenis() {
	fmt.Println("\n===Aplikasi menghitung massa jenis===")
	fmt.Println("Massa Jenis = Massa / Volume")
	fmt.Print("Masukkan massa (kg): ")
	massa, _ := lib.GetInput()
	fmt.Print("Masukkan volume (m3): ")
	volume, _ := lib.GetInput()
	result := MassaJenis(massa, volume)
	fmt.Printf("Massa Jenis adalah %.2f kg/m3\n", result)
}