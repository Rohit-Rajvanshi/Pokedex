package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"github.com/Rohit-Rajvanshi/Pokedex/pokecache"
)




func commandMapb (c *config , _ string) error {
	if c.PreviousURL == "" {
		fmt.Println("No entries previous to this")
		return nil 
	}

	cache := pokecache.NewCache(5 * time.Second)
	var body []byte
	val , ok := cache.Get(c.PreviousURL)
	if ok{
		body = val
	}else{
		res , err := http.Get(c.PreviousURL)
		if err != nil {
			log.Fatal(err)
		}
		
		body , err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Respone failed with status code: %d", res.StatusCode)
		}
	}
	var mapc mapCommand
	if err:= json.Unmarshal(body , &mapc) ; err != nil {
		log.Fatal(err)
	}

	for _ , elem := range mapc.Results {
		fmt.Println(elem.Name)
	}

	c.NextURL = mapc.Next
	c.PreviousURL = mapc.Previous

	return nil 
}