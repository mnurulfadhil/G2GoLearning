package main

import (
	"fmt"
	"time"
)

func main() {
	kendaraan := make(map[int]time.Time)
	id := 1
	var terminalinput int

menu:
	fmt.Println("1.Parkir")
	fmt.Println("2.Keluar")
	fmt.Println("3.Data kendaraan Parkir")

	fmt.Println("masukkan pilihan :")
	fmt.Scanln(&terminalinput)

	switch terminalinput {
	case 1:
		x := id
		waktumasuk := time.Now()
		kendaraan[x] = waktumasuk
		fmt.Println("\n====================")
		fmt.Println("Kendaraan dengan id :", x, "\nmasuk :", waktumasuk)
		fmt.Println("====================")
		id++
		goto menu

	case 2:
		var id int
		var tipe string

		fmt.Println("id parkir :")
		fmt.Scanln(&id)
		fmt.Println("Jenis Kendaraam : [mobil/motor]")
		fmt.Scanln(&tipe)
		waktukeluar := time.Now()
		waktuparkir := kendaraan[id]
		durasi := waktukeluar.Sub(waktuparkir)
		detik := int(durasi / time.Second)

		if tipe == "mobil" {
			harga := (3000 * detik) + 5000
			if harga > 100000 {
				fmt.Println("\n====================")
				fmt.Println(tipe, "dengan id :", id)
				fmt.Println("durasi parkir :", detik, " detik")
				fmt.Println("biaya parkir anda :", 100000)
				fmt.Println("====================")
				delete(kendaraan, id)
			} else {
				fmt.Println("\n====================")
				fmt.Println(tipe, "dengan id :", id)
				fmt.Println("durasi parkir :", detik, " detik")
				fmt.Println("biaya parkir anda :", harga)
				fmt.Println("====================")
				delete(kendaraan, id)
			}
			goto menu
		} else if tipe == "motor" {
			harga := (2000 * detik) + 3000
			if harga > 50000 {
				fmt.Println("\n====================")
				fmt.Println(tipe, "dengan id :", id)
				fmt.Println("durasi parkir :", detik, " detik")
				fmt.Println("biaya parkir anda :", 50000)
				fmt.Println("====================")
				delete(kendaraan, id)
			} else {
				fmt.Println("\n====================")
				fmt.Println(tipe, "dengan id :", id)
				fmt.Println("durasi parkir :", detik, " detik")
				fmt.Println("biaya parkir anda :", harga)
				fmt.Println("====================")
				delete(kendaraan, id)
			}
			goto menu
		} else {
			fmt.Println("\n====================")
			fmt.Println("tidak ada kendaraan dengan tipe :", tipe, " yang parkir")
			fmt.Println("====================")
			goto menu
		}
	case 3:
		for key, value := range kendaraan {
			fmt.Println("id :", key, "\t", value)
			goto menu
		}
	default:
	}
}
