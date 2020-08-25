package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type masukc struct {
	ID   string
	Time time.Time
}
type keluarc struct {
	ID    string
	Tipe  string
	hasil int
}

var baseUrl = "http://localhost:8080"

func main() {
	var terminalinput int
menu:
	fmt.Println("===>PARKIR<===")
	fmt.Println("1.Parkir")
	fmt.Println("2.Keluar")
	fmt.Scanln(&terminalinput)
	switch terminalinput {
	case 1:
		tamu, err := user()
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}
		fmt.Println("ID :", tamu.ID)
		fmt.Println("Jam masuk :", tamu.Time)
		goto menu
	case 2:
		var idp, tipekendp string
		fmt.Println("masukkan ID :")
		fmt.Scanln(&idp)
		fmt.Println("masukkan jenis kendaraan:")
		fmt.Scanln(&tipekendp)
		kel, err := userout(idp, tipekendp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("ID :", idp)
		fmt.Println("tagihan :", kel.hasil)
		goto menu
	}
}

func user() (masukc, error) {
	var err error
	var client = &http.Client{}
	var data masukc

	request, err := http.NewRequest("POST", baseUrl+"/masuk", nil)
	if err != nil {
		return data, err
	}
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func userout(Id string, tipe string) (keluarc, error) {
	var err error
	var client = &http.Client{}
	var datac keluarc

	var param = url.Values{}
	param.Set("id", Id)
	param.Set("tipe", tipe)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseUrl+"/keluar", payload)

	if err != nil {
		return datac, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return datac, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&datac)

	if err != nil {
		return datac, err
	}
	return datac, nil
}
