package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Count int

type Game struct {
	PointsEarned int
	Turn         int
	Players      []Player
}

type Player struct {
	Name   string
	Scores []int
}

func CreateGame(n Count) (g Game) {
	r := bufio.NewReader(os.Stdin)
	for i := Count(0); i < n; i++ {
		fmt.Printf("Enter name for player %d: ", i+1)
		for {
			name, err := r.ReadString('\n')
			if err == nil && len(strings.TrimSpace(name)) != 0 {
				g.Players = append(g.Players, Player{Name: name[:len(name)-1]})
				break
			}
			fmt.Println("Enter a name containing characters")
		}
	}
	return
}

func (g *Game) Play(input chan Game, output chan Game) {
	g.DisplayPoints()
	for {
		fmt.Println("Points Earned:")
		g.PointsEarned = Earned()
		input <- *g
		game := <-output
		g = &game
		g.DisplayPoints()
	}
}

func (g Game) DisplayPoints() {
	fmt.Printf("Turn: %s\n", g.Players[g.Turn].Name)
	for i, p := range g.Players {
		fmt.Printf("%d | %s%s\n", Sum(p.Scores), p.Name, g.TurnMarker(i))
	}
}

func (g Game) TurnMarker(i int) string {
	if g.Turn == i {
		return " <--"
	}
	return ""
}
