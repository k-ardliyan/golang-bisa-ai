package rumus

import (
	"fmt"
	"golang-bisa-ai/lib"
	"math"
)

func Sinus(sudut float64) float64 {
	return math.Sin(sudut * math.Pi / 180)
}

func Cosinus(sudut float64) float64 {
	return math.Cos(sudut * math.Pi / 180)
}

func Tangen(sudut float64) float64 {
	return math.Tan(sudut * math.Pi / 180)
}

func HitungSinusCosinusTangen() {
	fmt.Println("\n===Aplikasi menghitung sudut Sinus, Cosinus, Tangen===")
	fmt.Print("Masukkan sudut (derajat): ")
	sudut, _ := lib.GetInput()
	resultSinus := Sinus(sudut)
	resultCosinus := Cosinus(sudut)
	resultTangen := Tangen(sudut)
	fmt.Printf("Sinus adalah %.2f\n", resultSinus)
	fmt.Printf("Cosinus adalah %.2f\n", resultCosinus)
	fmt.Printf("Tangen adalah %.2f\n", resultTangen)
}