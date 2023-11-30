package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
	"math"
)

func LuasLingkaran(jariJari float64) float64 {
	return math.Pi * math.Pow(jariJari, 2)
}

func HitungLuasLingkaran() {
	fmt.Println("\n===Aplikasi menghitung luas lingkaran===")
	fmt.Println("Luas Lingkaran = π * r^2")
	fmt.Print("Masukkan jari-jari lingkaran (cm): ")
	jariJari, _ := lib.GetInput()
	result := LuasLingkaran(jariJari)
	fmt.Printf("Luas Lingkaran adalah %.2f cm2\n", result)
}