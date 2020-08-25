package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

menukabataku:
	var kalkulator string
	fmt.Println("pilih operation :")
	fmt.Println("1.Penambahan")
	fmt.Println("2.Pengurangan")
	fmt.Println("3.Perkalian")
	fmt.Println("4.Pembagian")
	fmt.Println("5.Pangkat")
	fmt.Println("6.VolumeTabung")

	fmt.Scanln(&kalkulator)

	var x string
	fmt.Println("x=")
	fmt.Scanln(&x)
	var y string
	fmt.Println("y=")
	fmt.Scanln(&y)
	// var z string
	// fmt.Println("pangkat=")
	// fmt.Scanln(&z)

	switch kalkulator {
	case "1":
		var num1 int = stringToInt(x)
		var num2 int = stringToInt(y)
		fmt.Println("=\t", Tambah(num1, num2))
		goto menukabataku
	case "2":
		var num1 int = stringToInt(x)
		var num2 int = stringToInt(y)
		fmt.Println("=\t", Subtract(num1, num2))
		goto menukabataku
	case "3":
		var num1 float64 = stringToFloat(x)
		var num2 float64 = stringToFloat(y)
		fmt.Println("=\t", Kali(num1, num2))
		goto menukabataku
	case "4":
		var num1 float64 = stringToFloat(x)
		var num2 float64 = stringToFloat(y)
		fmt.Println("=\t", Bagi(num1, num2))
		goto menukabataku
	case "5":
		var num1 float64 = stringToFloat(x)
		var num2 float64 = stringToFloat(y)
		fmt.Println("= ", Pangkat(num1, num2))
		goto menukabataku
	case "6":
		var num1 float64 = stringToFloat(x)
		var num2 float64 = stringToFloat(y)
		fmt.Println("= ", Volumetabung(num1, num2))
		goto menukabataku
	default:
		fmt.Println("invalid keyword")
		goto menukabataku

	}

}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return i
}

func stringToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return f
}

func Tambah(x, y int) int {
	return x + y
}
func Subtract(x, y int) int {
	return x - y
}
func Kali(x, y float64) float64 {
	return x * y
}
func Bagi(x, y float64) float64 {
	return x / y
}
func Pangkat(x, y float64) float64 {
	return math.Pow(x, y)
}

func Volumetabung(x, y float64) float64 {
	return 3.14 * x * x * y
}
