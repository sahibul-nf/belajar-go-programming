package main

import (
	"fmt"
)

type Gamer struct {
	Name  string
	Games []string
}

// method
func (gamer *Gamer) AddGame(name string) {
	gamer.Games = append(gamer.Games, name)
}

func main() {

	gamer := Gamer{}

	gamer.AddGame("Sahibul")
	gamer.AddGame("Nuzul")
	gamer.AddGame("Firdaus")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
