package main 

import (
	"strings"
	"fmt"
	"bufio"
	"os"

)

func cleanInput(text string) []string{
	words  := []string{}
	lowerText := strings.ToLower(text)
	words = strings.Fields(lowerText)
	return words 
}

type config struct{
		commands  map[string]cliCommand 
		NextURL string 
		PreviousURL string 
		Pokedex map[string]Pokemon
	}

type cliCommand struct {
		name string
		description string 
		callback func(*config , string) error
	}

type mapCommand struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string   `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}




func startRepl(c *config) error  {
	scanner := bufio.NewScanner(os.Stdin)


	for { 
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		command := cleanText[0]
		if len(cleanText) == 1 {
			cleanText = append(cleanText, "")
		}
		arg := cleanText[1]
		cmd , ok := c.commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}	
		cmd.callback(c , arg)
			
	}
}




