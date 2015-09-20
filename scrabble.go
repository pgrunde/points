package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pgrunde/scrabble/game"
)

func main() {
	count := PlayerCount()
	fmt.Printf("Creating game with %d players\n", count)
	g := game.CreateGame(count)

	pointsInput := make(chan game.Game)
	defer close(pointsInput)
	outputCalc := game.PointAdder(pointsInput)
	defer close(outputCalc)
	g.Play(pointsInput, outputCalc)
}

func PlayerCount() game.Count {
	r := bufio.NewReader(os.Stdin)
	var c game.Count
	fmt.Println("Enter number of players:")
	for {
		playerCount, err := r.ReadString('\n')
		n, err := strconv.Atoi(playerCount[:len(playerCount)-1])
		if err == nil && n > 1 {
			c = game.Count(n)
			return c
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Enter valid integer greater than one")
	}
}
