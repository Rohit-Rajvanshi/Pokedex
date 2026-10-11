package main

import "fmt"




func commandInspect (c *config , pokemon string) error {
	val , ok := c.Pokedex[pokemon] 
	if ok == false {
		fmt.Println("you have not caught that pokemon")
		return nil 
	}

	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Weight: %d\n", val.Weight)

	fmt.Println("Stats:")
	for _ , elem := range val.Stats{
		fmt.Printf("  -%s: %d\n",elem.Stat.Name , elem.BaseStat)
	}
	
	fmt.Println("Types:")
	for _ , elem := range val.Types {
		fmt.Printf("  -%s\n" , elem.Type.Name)
	}


	return nil 
}
	