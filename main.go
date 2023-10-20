package main

import (
	"log"
)

type DBInterface interface {
	Get()
	Delete()
}

type DBInfo3rdParty struct {
	username string
	password string
}

type DBHandler struct {
	DBInfo3rdParty
}

func DBUseInterface(d DBInterface) {

	d.Get()
	d.Delete()

}

func (d *DBHandler) Get() {
	log.Println("Get Process")
}

func (d *DBHandler) Delete() {
	log.Println("Delete Process")
}

func main() {

	d := &DBHandler{
		DBInfo3rdParty{
			username: "user",
			password: "pass",
		},
	}
	DBUseInterface(d)
}
