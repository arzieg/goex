package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.Handle("/list", http.HandlerFunc(db.list))
	http.Handle("/read", http.HandlerFunc(db.read))
	http.Handle("/update", http.HandlerFunc(db.update))
	http.Handle("/delete", http.HandlerFunc(db.delete))
	http.Handle("/create", http.HandlerFunc(db.create))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// create
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusFound)
		fmt.Fprintf(w, "I found already %s with price %s\n", item, db[item])
		return
	}
	floatprice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "could not get the price: %q\n", price)
		return
	}
	db[item] = dollars(floatprice)
	fmt.Fprintf(w, "Add %s with price %s\n", item, db[item])

}

// update
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	floatprice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "price could not be converted to float: %q\n", price)
		return
	}
	db[item] = dollars(floatprice)
	fmt.Fprintf(w, "%s\n", db[item])

}

// delete
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "Deleted %s\n", item)

}
