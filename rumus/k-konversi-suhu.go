package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
)

func CelciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}

func CelciusToKelvin(celcius float64) float64 {
	return celcius + 273
}

func FahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273
}

func KelvinToCelcius(kelvin float64) float64 {
	return kelvin - 273
}

func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273)*9/5 + 32
}

func HitungKonversiSuhu() {
	fmt.Println("\n===Aplikasi menghitung konversi suhu===")
	fmt.Println("\nPilih rumus yang akan digunakan:")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celcius")
	fmt.Println("4. Fahrenheit ke Kelvin")
	fmt.Println("5. Kelvin ke Celcius")
	fmt.Println("6. Kelvin ke Fahrenheit")
	fmt.Print("Pilihan: ")
	pilihan, _ := lib.GetInput()
	for pilihan < 1 || pilihan > 6 {
		fmt.Println("Pilihan tidak tersedia")
		fmt.Print("Pilihan: ")
		pilihan, _ = lib.GetInput()
	}
	switch pilihan {
	case 1:
		fmt.Println("\n===Aplikasi menghitung Celcius ke Fahrenheit===")
		fmt.Println("Celcius ke Fahrenheit = (Celcius * 9/5) + 32")
		fmt.Print("Masukkan suhu dalam Celcius: ")
		celcius, _ := lib.GetInput()
		result := CelciusToFahrenheit(celcius)
		fmt.Printf("Hasil konversi adalah %.2f Fahrenheit\n", result)
	case 2:
		fmt.Println("\n===Aplikasi menghitung Celcius ke Kelvin===")
		fmt.Println("Celcius ke Kelvin = Celcius + 273")
		fmt.Print("Masukkan suhu dalam Celcius: ")
		celcius, _ := lib.GetInput()
		result := CelciusToKelvin(celcius)
		fmt.Printf("Hasil konversi adalah %.2f Kelvin\n", result)
	case 3:
		fmt.Println("\n===Aplikasi menghitung Fahrenheit ke Celcius===")
		fmt.Println("Fahrenheit ke Celcius = (Fahrenheit - 32) * 5/9")
		fmt.Print("Masukkan suhu dalam Fahrenheit: ")
		fahrenheit, _ := lib.GetInput()
		result := FahrenheitToCelcius(fahrenheit)
		fmt.Printf("Hasil konversi adalah %.2f Celcius\n", result)
	case 4:
		fmt.Println("\n===Aplikasi menghitung Fahrenheit ke Kelvin===")
		fmt.Println("Fahrenheit ke Kelvin = (Fahrenheit - 32) * 5/9 + 273")
		fmt.Print("Masukkan suhu dalam Fahrenheit: ")
		fahrenheit, _ := lib.GetInput()
		result := FahrenheitToKelvin(fahrenheit)
		fmt.Printf("Hasil konversi adalah %.2f Kelvin\n", result)
	case 5:
		fmt.Println("\n===Aplikasi menghitung Kelvin ke Celcius===")
		fmt.Println("Kelvin ke Celcius = Kelvin - 273")
		fmt.Print("Masukkan suhu dalam Kelvin: ")
		kelvin, _ := lib.GetInput()
		result := KelvinToCelcius(kelvin)
		fmt.Printf("Hasil konversi adalah %.2f Celcius\n", result)
	case 6:
		fmt.Println("\n===Aplikasi menghitung Kelvin ke Fahrenheit===")
		fmt.Println("Kelvin ke Fahrenheit = (Kelvin - 273) * 9/5 + 32")
		fmt.Print("Masukkan suhu dalam Kelvin: ")
		kelvin, _ := lib.GetInput()
		result := KelvinToFahrenheit(kelvin)
		fmt.Printf("Hasil konversi adalah %.2f Fahrenheit\n", result)
	}
}
