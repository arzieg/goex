package db

import (
	"abstraction/storage"
	"fmt"
	"log"
)

type DB interface {
	Create() error
	Insert() error
	Dump() error
}

type CreateDB struct {
	DBname string
}

type InsertDB struct {
	DBContent string
}

type DumpDB struct {
	CreateDB
	InsertDB
}

func (d *CreateDB) Create() error {
	fmt.Printf("Create DB %s\n", d.DBname)
	return nil
}

func (d *InsertDB) Insert() error {
	fmt.Printf("Insert into DB %s\n", d.DBContent)
	return nil
}

func (d *DumpDB) Dump() error {
	var s storage.Storage

	s = &storage.FileStorage{Filename: d.DBname, Content: d.DBContent}

	if err := s.Save(); err != nil {
		log.Fatal(err)
	}
	return nil
}
