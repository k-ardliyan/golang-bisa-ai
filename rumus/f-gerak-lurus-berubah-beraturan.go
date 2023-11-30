package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func GerakLurusBerubahBeraturan(kecepatanAwal, percepatan, waktu float64) float64 {
	return kecepatanAwal + percepatan * waktu
}

func HitungGerakLurusBerubahBeraturan() {
	fmt.Println("\n===Aplikasi menghitung Gerak Lurus Berubah Beraturan===")
	fmt.Println("Rumus: v = v0 + at")
	fmt.Print("Masukkan kecepatan awal (m/s): ")
	kecepatanAwal, _ := lib.GetInput()
	fmt.Print("Masukkan percepatan (m/s2): ")
	percepatan, _ := lib.GetInput()
	fmt.Print("Masukkan waktu (detik): ")
	waktu, _ := lib.GetInput()
	result := GerakLurusBerubahBeraturan(kecepatanAwal, percepatan, waktu)
	fmt.Printf("Gerak Lurus Berubah Beraturan adalah %.2f m/s\n", result)
	fmt.Printf("Gerak Lurus Berubah Beraturan adalah %.2f km/jam\n", result * 3.6)
}