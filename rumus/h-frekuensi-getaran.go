package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func Frekuensi(periode float64) float64 {
	return 1 / periode
}

func Getaran(periode, frekuensi float64) float64 {
	return periode * frekuensi
}

func HitungFrekuensiGetaran() {
	fmt.Println("\n===Aplikasi menghitung frekuensi getaran===")
	fmt.Println("\nPilih rumus yang akan digunakan:")
	fmt.Println("1. Frekuensi")
	fmt.Println("2. Getaran")
	fmt.Print("Pilihan: ")
	pilihan, _ := lib.GetInput()
	for pilihan < 1 || pilihan > 2 {
		fmt.Println("Pilihan tidak tersedia")
		fmt.Print("Pilihan: ")
		pilihan, _ = lib.GetInput()
	}
	switch pilihan {
	case 1:
		fmt.Println("\n===Aplikasi menghitung Frekuensi===")
		fmt.Println("Frekuensi = 1 / Periode")
		fmt.Print("Masukkan periode (s): ")
		periode, _ := lib.GetInput()
		result := Frekuensi(periode)
		fmt.Printf("Frekuensi adalah %.2f Hz\n", result)
	case 2:
		fmt.Println("\n===Aplikasi menghitung Getaran===")
		fmt.Println("Getaran = Periode * Frekuensi")
		fmt.Print("Masukkan periode (s): ")
		periode, _ := lib.GetInput()
		fmt.Print("Masukkan frekuensi (Hz): ")
		frekuensi, _ := lib.GetInput()
		result := Getaran(periode, frekuensi)
		fmt.Printf("Getaran adalah %.2f Hz\n", result)
	}
}