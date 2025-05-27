package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type response struct {
	Item   string `json:"item"`
	Album  string
	Title  string
	Artist string
}

type respWrapper struct {
	response
}

var j1 = `{
 "item" : "album", 
 "album" : {"title": "Dark Side of the Moon"}
 }`

var j2 = `{
  "item" : "song", 
  "song" : {"title": "Bella Donna", "artist": "Stevie Nicks"}
  }`

func main() {
	var resp1, resp2 respWrapper
	var err error

	if err = json.Unmarshal([]byte(j1), &resp1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp1.response)

	if err = json.Unmarshal([]byte(j2), &resp2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp2.response)
}

func (r *respWrapper) UnmarshalJSON(b []byte) (err error) {
	var raw map[string]interface{}       // any json object
	err = json.Unmarshal(b, &r.response) // ignore errors
	err = json.Unmarshal(b, &raw)        // maping string to interface
	switch r.Item {
	case "album":
		inner, ok := raw["album"].(map[string]interface{}) // wenn album gefunden wird, wird wieder ein map to interface erstellet

		if ok {
			if album, ok := inner["title"].(string); ok { // wenn inner["title] ein string ist, dann ...
				r.Album = album
			}
		}

	case "song":
		inner, ok := raw["song"].(map[string]interface{})

		if ok {
			if title, ok := inner["title"].(string); ok { // wenn inner["title] ein string ist, dann ...
				r.Title = title
			}
			if artist, ok := inner["artist"].(string); ok { // wenn inner["title] ein string ist, dann ...
				r.Artist = artist
			}
		}
	}
	return nil
}
