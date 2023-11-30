package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func Daya(usaha float64, waktu float64) float64 {
	return usaha / waktu
}

func Tekanan(gaya float64, luas float64) float64 {
	return gaya / luas
}

func Usaha(gaya float64, jarak float64) float64 {
	return gaya * jarak
}

func Gaya(massa float64, percepatan float64) float64 {
	return massa * percepatan
}

func HitungDayaTekananUsahaGaya() {
	fmt.Println("\n===Aplikasi menghitung Daya, Tekanan, Usaha dan Gaya===")
	fmt.Println("\nPilih rumus yang akan digunakan:")
	fmt.Println("1. Daya")
	fmt.Println("2. Tekanan")
	fmt.Println("3. Usaha")
	fmt.Println("4. Gaya")
	fmt.Print("Pilihan: ")
	pilihan, _ := lib.GetInput()
	for pilihan < 1 || pilihan > 4 {
		fmt.Println("Pilihan tidak tersedia")
		fmt.Print("Pilihan: ")
		pilihan, _ = lib.GetInput()
	}
	switch pilihan {
	case 1:
		fmt.Println("\n===Aplikasi menghitung Daya===")
		fmt.Println("Daya = Usaha / Waktu")
		fmt.Print("Masukkan usaha (Joule): ")
		usaha, _ := lib.GetInput()
		fmt.Print("Masukkan waktu (detik): ")
		waktu, _ := lib.GetInput()
		result := Daya(usaha, waktu)
		fmt.Printf("Daya adalah %.2f Watt\n", result)
	case 2:
		fmt.Println("\n===Aplikasi menghitung Tekanan===")
		fmt.Println("Tekanan = Gaya / Luas")
		fmt.Print("Masukkan gaya (Newton): ")
		gaya, _ := lib.GetInput()
		fmt.Print("Masukkan luas (m2): ")
		luas, _ := lib.GetInput()
		result := Tekanan(gaya, luas)
		fmt.Printf("Tekanan adalah %.2f Pascal\n", result)
	case 3:
		fmt.Println("\n===Aplikasi menghitung Usaha===")
		fmt.Println("Usaha = Gaya * Jarak")
		fmt.Print("Masukkan gaya (Newton): ")
		gaya, _ := lib.GetInput()
		fmt.Print("Masukkan jarak (m): ")
		jarak, _ := lib.GetInput()
		result := Usaha(gaya, jarak)
		fmt.Printf("Usaha adalah %.2f Joule\n", result)
	case 4:
		fmt.Println("\n===Aplikasi menghitung Gaya===")
		fmt.Println("Gaya = Massa * Percepatan")
		fmt.Print("Masukkan massa (kg): ")
		massa, _ := lib.GetInput()
		fmt.Print("Masukkan percepatan (m/s2): ")
		percepatan, _ := lib.GetInput()
		result := Gaya(massa, percepatan)
		fmt.Printf("Gaya adalah %.2f Newton\n", result)
	}
}