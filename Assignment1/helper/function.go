package helpers

import (
	"fmt"
	"strings"
)

// struct
type User struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

// data user
func Init() []User {
	data_users := []User{
		{
			Nama:      "Thomas",
			Alamat:    "Kota A",
			Pekerjaan: "Programmer",
			Alasan:    "Alasan Thomas",
		},
		{
			Nama:      "Sauki Adillah",
			Alamat:    "Kp. Cibadak",
			Pekerjaan: "Programmer & Sistem Analis",
			Alasan:    "Upgrade skill",
		},
		{
			Nama:      "Rio",
			Alamat:    "Bekasi",
			Pekerjaan: "Koordinator",
			Alasan:    "Pengen tau aja",
		},
	}
	return data_users
}

// print biodata
func ShowBiodata(data_users []User, nama string) {
	// fmt.Println(nama)
	for _, v := range data_users {
		// fmt.Println(v.Nama)
		if strings.Contains(v.Nama, nama) {
			fmt.Println("Nama : ", v.Nama)
			fmt.Println("Alamat : ", v.Alamat)
			fmt.Println("Pekerjaan : ", v.Pekerjaan)
			fmt.Println("Alasan : ", v.Alasan)
			fmt.Println("============================================")
		}
	}
}
