package main

import (
	"fmt" 
	"net/http"
	"log"
	"io"
	"encoding/json"
	"math/rand"
)

type Pokemon struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat struct {
			Name string  `json:"name"`
		}`json:"stat"`
	}`json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	}`json:"types"`
}


func commandCatch (c *config , pokemon string) error {

	if pokemon == "" {
		fmt.Println("Please enter a valid Pokemon name")
		return fmt.Errorf("No valid Pokemon name\n")
	}

	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemon 
	res , err := http.Get(fullUrl)
	if err != nil{
		log.Fatal(err)
	}

	defer res.Body.Close()

	body , err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		fmt.Println("\nPlease enter a valid Pokemon name")
		log.Fatalf("Respone failed with status code: %d", res.StatusCode)
	}

	var poke Pokemon

	if err := json.Unmarshal(body , &poke) ; err != nil {
		log.Fatal(err)
	}

	random := rand.Intn(poke.BaseExperience)
	threshold := 50 

	fmt.Printf("Throwing a Pokeball at %s...\n" , pokemon)

	if random < threshold {
		fmt.Printf("%s was caught!\n" , pokemon)
		c.Pokedex[pokemon] = poke 

	}else{
		fmt.Printf("%s escaped!\n" , pokemon)
	}



	return nil 
}

