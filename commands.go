package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExplore(c *config, parameters []string) error {
	if len(parameters) <= 1 {
		return errors.New("missing explore parameters")
	}
	area := "/location-area/" + parameters[1]
	deepLocationsResp, err := c.pokeapiClient.DeepListLocation(&area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", parameters[1])
	for _, poke := range deepLocationsResp.PokemonEncounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}
	return nil
}

func commandExit(c *config, parameters []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, parameters []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(c *config, parameters []string) error {

	locationsResp, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config, parameters []string) error {
	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationResp.Next
	c.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandCatch(c *config, parameters []string) error {
	if len(parameters) <= 1 {
		return errors.New("missing pokemon name parameters")
	}

	pokemonName := parameters[1]
	pokemonData, err := c.pokeapiClient.PokemonData(&pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)
	randNum := rand.Intn(pokemonData.BaseExperience)
	if randNum > (pokemonData.BaseExperience / 3) {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		c.pokedex[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return nil
}

func commandInspect(c *config, paramaters []string) error {
	pokemonName := paramaters[1]

	if pokemon, ok := c.pokedex[pokemonName]; ok {

		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, value := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", value.Stat.Name, value.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	} else {
		return errors.New("invalid pokemon name or pokemon has not been caught")
	}

	return nil
}

func commandPokedex(c *config, parameters []string) error {
	if len(c.pokedex) == 0 {
		return errors.New("pokedex is empty, go catch some pokemon")
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range c.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
