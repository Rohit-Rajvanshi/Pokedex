package main

import (
	"fmt" 
	"net/http"
	"log"
	"io"
	"encoding/json"
	"github.com/Rohit-Rajvanshi/Pokedex/pokecache"
	"time"
)


type exploreCommand struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}


func commandExplore (c *config , area string) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area/" + area

	cache := pokecache.NewCache(5 * time.Second)
	val , ok := cache.Get(c.NextURL)

	var body []byte

	if ok{
		body = val
	}else{
		res , err := http.Get(fullUrl)
		if err != nil {
			log.Fatal(err)
		}

		body , err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode > 299 {
				log.Fatalf("Respone failed with status code: %d", res.StatusCode)
			}
		
		cache.Add(fullUrl , body)

		defer res.Body.Close()
	}


	var explorec exploreCommand

	if err := json.Unmarshal(body , &explorec); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exploring %s...\n" , area)
	
	for _ , elem := range explorec.PokemonEncounters {
		fmt.Println(elem.Pokemon.Name)
	}

	return nil 
}