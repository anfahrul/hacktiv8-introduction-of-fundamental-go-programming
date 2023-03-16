package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func getFriendById(id int) (Friend, error) {
	friends := map[int]Friend{
		1: {
			ID:      1,
			Name:    "Paulin Mayasari",
			Address: "Jl Kebahagiaan 54, Dki Jakarta",
			Job:     "Backend Developer",
			Reason:  "Belajar",
		},
		2: {
			ID:      2,
			Name:    "Cayadi Natsir",
			Address: "Jl Tubagus Ismail Raya 3 B, Jawa Barat",
			Job:     "Software Engineer",
			Reason:  "Belajar",
		},
		3: {
			ID:      3,
			Name:    "Mala Puspita",
			Address: "Jl Tebet Tmr Dlm 113, Dki Jakarta",
			Job:     "DevOps Engineer",
			Reason:  "Belajar",
		},
		4: {
			ID:      4,
			Name:    "Laksana Simanjuntak",
			Address: "JL Tj Pura 12, Benua Melayu Darat, Pontianak Selatan",
			Job:     "Backend Developer",
			Reason:  "Belajar",
		},
		5: {
			ID:      5,
			Name:    "Lala Nuraini",
			Address: "Jl Sindang Palay 9, Jawa Barat",
			Job:     "Software Engineer",
			Reason:  "Belajar",
		},
		6: {
			ID:      6,
			Name:    "Ina Nuraini",
			Address: "Jl PKP Kelapa Dua Wetan, Dki Jakarta",
			Job:     "Software Engineer",
			Reason:  "Belajar",
		},
		7: {
			ID:      7,
			Name:    "Qori Pudjiastuti",
			Address: "Jl KH Samanhudi 21 D, Dki Jakarta",
			Job:     "Backend Developer",
			Reason:  "Belajar",
		},
		8: {
			ID:      8,
			Name:    "Kezia Mardhiyah",
			Address: "Jl Pakuningratan 25, Jawa Tengah",
			Job:     "Software Engineer",
			Reason:  "Belajar",
		},
		9: {
			ID:      9,
			Name:    "Capa Damanik",
			Address: "Jl P Tubagus Angke Bl AL/6, Dki Jakarta",
			Job:     "Software Engineer",
			Reason:  "Belajar",
		},
		10: {
			ID:      10,
			Name:    "Karsa Maulana",
			Address: "Jl Kutilang 1, Jawa Timur",
			Job:     "Backend Developer",
			Reason:  "Belajar",
		},
	}

	if value, exist := friends[id]; exist {
		return value, nil
	}

	return Friend{}, errors.New(fmt.Sprintf("Friend with id %v not found", id))
}

func main() {
	if len(os.Args) == 1 {
		panic("ID belum dimasukkan!")
	}
	var friendID = os.Args[1]
	var friendIDInt, _ = strconv.Atoi(friendID)

	friend, err := getFriendById(friendIDInt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan ikut kelas Golang: %s",
			friend.Name,
			friend.Address,
			friend.Job,
			friend.Reason)
	}
}
