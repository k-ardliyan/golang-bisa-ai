package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
	"math"
)

func EnergiPotensial(massa, gravitasi, ketinggian float64) float64 {
	return massa * gravitasi * ketinggian
}

func EnergiKinetik(massa, kecepatan float64) float64 {
	return 0.5 * massa * math.Pow(kecepatan, 2)
}

func HitungEnergiPotensialKinetik() {
	fmt.Println("\n===Aplikasi menghitung Energi Potensial dan Kinetik===")
	fmt.Println("\nPilih rumus yang akan digunakan:")
	fmt.Println("1. Energi Potensial")
	fmt.Println("2. Energi Kinetik")
	fmt.Print("Pilihan: ")
	pilihan, _ := lib.GetInput()
	for pilihan < 1 || pilihan > 2 {
		fmt.Println("Pilihan tidak tersedia")
		fmt.Print("Pilihan: ")
		pilihan, _ = lib.GetInput()
	}
	switch pilihan {
	case 1:
		fmt.Println("\n===Aplikasi menghitung Energi Potensial===")
		fmt.Println("Energi Potensial = Massa * Gravitasi * Ketinggian")
		fmt.Print("Masukkan massa (kg): ")
		massa, _ := lib.GetInput()
		fmt.Print("Masukkan gravitasi (m/s2): ")
		gravitasi, _ := lib.GetInput()
		fmt.Print("Masukkan ketinggian (m): ")
		ketinggian, _ := lib.GetInput()
		result := EnergiPotensial(massa, gravitasi, ketinggian)
		fmt.Printf("Energi Potensial adalah %.2f Joule\n", result)
	case 2:
		fmt.Println("\n===Aplikasi menghitung Energi Kinetik===")
		fmt.Println("Energi Kinetik = 0.5 * Massa * Kecepatan^2")
		fmt.Print("Masukkan massa (kg): ")
		massa, _ := lib.GetInput()
		fmt.Print("Masukkan kecepatan (m/s): ")
		kecepatan, _ := lib.GetInput()
		result := EnergiKinetik(massa, kecepatan)
		fmt.Printf("Energi Kinetik adalah %.2f Joule\n", result)
	}
}