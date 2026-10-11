package main 


func main()  {
	c := &config{
		commands : map[string]cliCommand {
			"exit" : {
				name: "exit",
				description : "Exit the pokedex",
				callback : commandExit,	
			},
			"help" : {
				name : "help" ,
				description : "Displays a help message",
				callback : commandHelp,
			},
			"map" : {
				name : "map",
				description : "Display the location of pokemon areas",
				callback : commandMap,
			},
			"mapb" : {
				name : "mapb" , 
				description : "Display the previous 20 location of the pokemon areas",
				callback : commandMapb,
			},
			"explore" : {
				name : "explore" ,
				description : "List all the pokemons in the area",
				callback: commandExplore,
			},
			"catch" : {
				name : "catch" ,
				description: "Catch a pokemon",
				callback: commandCatch,
			},
			"inspect" : {
				name : "inspect" , 
				description: "Inspect a pokemon",
				callback: commandInspect,
			},
			"pokedex": {
				name : "pokedex" ,
				description: "Show the list of all your pokemons caught",
				callback: commandPokedex,
			},
		},
		NextURL : "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0",
		PreviousURL : "" ,
		Pokedex : make(map[string]Pokemon),
	}
	startRepl(c)
}







