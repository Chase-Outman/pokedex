package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Chase-Outman/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokedex          map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := strings.Fields(scanner.Text())

		if command, ok := getCommands()[inputs[0]]; ok {
			err := command.callback(cfg, inputs)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	cleanedText := strings.Trim(strings.ToLower(text), " ")

	return strings.Split(cleanedText, " ")
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previos page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemon encounters in area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the named pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display information about pokemon if it has been caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display a list of pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}
