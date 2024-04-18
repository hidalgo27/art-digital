package main

import (
	"fmt"
	"github.com/hidalgo27/art-digital.git/data"
	"github.com/hidalgo27/art-digital.git/db"
)

func main() {
	fmt.Println("Poblando la base de datos")

	db.Init()
	db.DB.AutoMigrate(&data.User{})

	db.Init()
	db.DB.AutoMigrate(&data.User{})

	db.DB.Create(&data.User{
		Name:  "John Smith",
		Email: "john.smith@gmail.com",
	})

	db.DB.Create(&data.User{
		Name:  "Daniel Smith",
		Email: "daniel@gmail.com",
	})

}
