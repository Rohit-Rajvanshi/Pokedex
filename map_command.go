package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/Rohit-Rajvanshi/Pokedex/pokecache"
	"time"
)




func commandMap (c *config , _ string) error {

	cache := pokecache.NewCache(5 * time.Second)
	val , ok := cache.Get(c.NextURL)
	var body []byte
	if ok{
		body = val
	}else{
		res , err := http.Get(c.NextURL)
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

		cache.Add(c.NextURL , body)
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