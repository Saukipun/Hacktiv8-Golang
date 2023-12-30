package main

import (
	helpers "assigment1/helper"
	"os"
)

type User struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	data_users := helpers.Init()

	for i := 1; i < len(os.Args); i++ {
		name := os.Args[i]

		helpers.ShowBiodata(data_users, name)
	}
}
