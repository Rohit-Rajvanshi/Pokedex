package main 

import (
	"fmt"
)

func commandPokedex (c *config , _ string) error {
	fmt.Println("Your Pokedex:")
	for key := range c.Pokedex{
		fmt.Printf(" - %s\n" , key)
	}
	return nil 
}