package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	"net/http"
	"time"
)

type masuk struct {
	ID   string
	Time time.Time
}

type keluar struct {
	ID    string
	Tipe  string
	hasil int
}

var kendaraan = []masuk{}

func main() {
	http.HandleFunc("/masuk", getID)
	http.HandleFunc("/keluar", parkout)
	fmt.Println("running....")
	http.ListenAndServe(":8080", nil)
}

func getID(w http.ResponseWriter, r *http.Request) {
	var id string
	id = xid.New().String()
	time := time.Now()
	kas := masuk{id, time}
	kendaraan = append(kendaraan, kas)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		result, err := json.Marshal(kas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}
func parkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		now := time.Now()
		var id = r.FormValue("id")
		var tipekend = r.FormValue("tipe")
		var result []byte
		var datauser keluar
		var err error

		for _, each := range kendaraan {
			if each.ID == id {
				durasi := int(now.Sub(each.Time).Seconds())
				if tipekend == "mobil" {
					if durasi <= 1 {
						tagihan := 5000
						fmt.Println("id kendaraan :", id)
						fmt.Println("jenis kendaraan :", tipekend)
						fmt.Println("tagiahan parkir :", tagihan)
						datauser = keluar{id, tipekend, tagihan}
						result, err = json.Marshal(datauser)
						w.Write(result)
						return
					} else {
						tagihan := (durasi-1)*3000 + 5000
						fmt.Println("Id kendaraan : ", id)
						fmt.Println("jenis Kendaraan : ", tipekend)
						fmt.Println("Bayar : ", tagihan)
						datauser = keluar{id, tipekend, tagihan}
						result, err = json.Marshal(datauser)
						w.Write(result)
						return
					}
				} else if tipekend == "motor" {
					if durasi <= 1 {
						tagihan := 3000
						fmt.Println("id kendaraan :", id)
						fmt.Println("jenis kendaraan :", tipekend)
						fmt.Println("tagiahan parkir :", tagihan)
						datauser = keluar{id, tipekend, tagihan}
						result, err = json.Marshal(datauser)
						w.Write(result)
						return
					} else {
						tagihan := (durasi-1)*2000 + 3000
						fmt.Println("Id kendaraan : ", id)
						fmt.Println("jenis Kendaraan : ", tipekend)
						fmt.Println("Bayar : ", tagihan)
						datauser = keluar{id, tipekend, tagihan}
						result, err = json.Marshal(datauser)
						w.Write(result)
						return
					}
				}
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return

			}

		}
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}
