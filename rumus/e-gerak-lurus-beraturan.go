package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func GerakLurusBeraturan(jarak float64, waktu float64) float64 {
	return jarak / waktu
}

func HitungGerakLurusBeraturan() {
	fmt.Println("\n===Aplikasi menghitung Gerak Lurus Beraturan===")
	fmt.Println("Gerak Lurus Beraturan = Jarak / Waktu")
	fmt.Print("Masukkan jarak tempuh (meter): ")
	jarak, _ := lib.GetInput()
	fmt.Print("Masukkan waktu tempuh (detik): ")
	waktu, _ := lib.GetInput()
	result := GerakLurusBeraturan(jarak, waktu)
	fmt.Printf("Gerak Lurus Beraturan adalah %.2f m/s\n", result)
	fmt.Printf("Gerak Lurus Beraturan adalah %.2f km/jam\n", result * 3.6)
}