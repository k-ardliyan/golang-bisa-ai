package main

import (
	"fmt"
	"golang-bisa-ai/rumus"
	"os"
)

func main() {
	fmt.Println("===Aplikasi Golang Bisa AI===")

	options := []string{
		"Luas Persegi",
		"Luas Segitiga",
		"Luas Lingkaran",
		"Sudut Sinus, Cosinus, Tangen",
		"Gerak Lurus Beraturan",
		"Gerak Lurus Berubah Beraturan",
		"Energi Potensial, Kinetik",
		"Frekuensi atau Getaran",
		"Massa Jenis",
		"Daya, Tekanan, Usaha dan Gaya",
		"Konversi Suhu",
	}
	
	for {
		fmt.Println("\nPilihan Rumus:")
		for i, option := range options {
			fmt.Printf("%d. %s\n", i+1, option)
			if i == len(options)-1 {
				fmt.Println("0. KELUAR APLIKASI")
			}
		}

		var choice int
		fmt.Print("Masukkan angka untuk memilih rumus: ")
		_, err := fmt.Scan(&choice)

		if err != nil || choice < 0 || choice > len(options) {
			fmt.Println("\nInput tidak valid. Silakan coba lagi.")
			continue
		}

		switch choice {
		case 1:
			rumus.HitungLuasPersegi()
		case 2:
			rumus.HitungLuasSegitiga()
		case 3:
			rumus.HitungLuasLingkaran()
		case 4:
			rumus.HitungSinusCosinusTangen()
		case 5:
			rumus.HitungGerakLurusBeraturan()
		case 6:
			rumus.HitungGerakLurusBerubahBeraturan()
		case 7:
			rumus.HitungEnergiPotensialKinetik()
		case 8:
			rumus.HitungFrekuensiGetaran()
		case 9:
			rumus.HitungMassaJenis()
		case 10:
			rumus.HitungDayaTekananUsahaGaya()
		case 11:
			rumus.HitungKonversiSuhu()
		case 0:
			// Keluar dari program
			fmt.Println("Terima kasih! Program selesai.")
			os.Exit(0)
		}

		// Setelah mendapatkan hasil, tanyakan apakah ingin menjalankan ulang aplikasi
		for {
			fmt.Print("\nApakah Anda ingin menjalankan rumus lainnya? (y/n): ")
			var restartChoice string
			fmt.Scan(&restartChoice)

			if restartChoice == "y" || restartChoice == "Y" {
				// Jika jawabannya 'y', keluar dari loop dan kembali ke awal
				break
			} else if restartChoice == "n" || restartChoice == "N" {
				// Jika jawabannya 'n', keluar dari loop dan selesai program
				fmt.Println("Terima kasih! Program selesai.")
				os.Exit(0)
			} else {
				// Jika jawabannya bukan 'y' atau 'n', tanyakan lagi
				fmt.Println("Input tidak valid. Silakan coba lagi.")
			}
		}
	}
}