package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

type Pokemon struct {
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

var url string = "https://pokeapi.co/api/v2/pokemon/"

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
func selectPokemon() {
	var pokemonName string
	fmt.Println("Choose a pokemon")
	fmt.Scan(&pokemonName)
	fmt.Println("You choose", pokemonName)
	url = url + pokemonName
}

func main() {
	selectPokemon()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			pokemon := Pokemon{}
			err := json.Unmarshal(data, &pokemon)
			if err != nil {
				panic(err)
			} else {
				openbrowser(pokemon.Sprites.FrontDefault)
			}

		}
	}

}
