package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	pointsEarned int
	turn         int
	players      []player
}

type player struct {
	Name   string
	Scores []int
}

func main() {
	count := PlayerCount()
	fmt.Printf("Creating game with %d players\n", count)
	game := CreateGame(count)

	pointsInput := make(chan Game)
	defer close(pointsInput)
	outputCalc := PointAdder(pointsInput)
	defer close(outputCalc)
	game.Play(pointsInput, outputCalc)
}

func PointAdder(input chan Game) chan Game {
	outputGame := make(chan Game)
	go func() {
		for {
			g := <-input
			g.players[g.turn].Scores = append(g.players[g.turn].Scores, g.pointsEarned)
			g.turn += 1
			if g.turn == len(g.players) {
				g.turn = 0
			}
			outputGame <- g
		}
	}()
	return outputGame
}

func (g *Game) Play(input chan Game, output chan Game) {
	g.DisplayPoints()
	for {
		fmt.Println("Points Earned:")
		g.pointsEarned = Earned()
		input <- *g
		game := <-output
		g = &game
		g.DisplayPoints()
	}
}

func Earned() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		earned, err := strconv.Atoi(line[:len(line)-1])
		if err == nil {
			return earned
		}
		fmt.Println(err)
		fmt.Println("Enter a valid integer")
	}
}

func (g Game) DisplayPoints() {
	fmt.Printf("Turn: %s\n", g.players[g.turn].Name)
	for i, p := range g.players {
		fmt.Printf("%d | %s%s\n", Sum(p.Scores), p.Name, g.TurnMarker(i))
	}
}

func (g Game) TurnMarker(i int) string {
	if g.turn == i {
		return " <--"
	}
	return ""
}

func CreateGame(n int) (g Game) {
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter name for player %d: ", i+1)
		for {
			name, err := r.ReadString('\n')
			if err == nil && len(strings.TrimSpace(name)) != 0 {
				g.players = append(g.players, player{Name: name[:len(name)-1]})
				break
			}
			fmt.Println("Enter a name containing characters")
		}
	}
	return
}

func PlayerCount() int {
	r := bufio.NewReader(os.Stdin)
	var c int
	fmt.Println("Enter number of players:")
	for {
		playerCount, err := r.ReadString('\n')
		n, err := strconv.Atoi(playerCount[:len(playerCount)-1])
		if err == nil && n != 1 {
			c = n
			return c
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Enter valid integer greater than one")
	}
}

func Sum(score []int) (total int) {
	for _, s := range score {
		total += s
	}
	return
}
