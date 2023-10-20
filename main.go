package main

import (
	"log"
)

type DBInfo3rdParty struct {
	username string
	password string
}

type DBHandler struct {
	DBInfo3rdParty
}

func DBNormal(username, password string) {

	d := &DBHandler{
		DBInfo3rdParty{
			username: username,
			password: password,
		},
	}

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
	DBNormal("user", "pass")
}
