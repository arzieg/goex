package main

import (
	"abstraction/db"
	"log"
	"os"
)

func main() {

	var d db.DB

	d = &db.DumpDB{CreateDB: db.CreateDB{DBname: "myGloryDB"},
		InsertDB: db.InsertDB{DBContent: "This is my content"}}

	if err := d.Dump(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

}
