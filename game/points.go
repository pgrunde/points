package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PointAdder(input chan Game) chan Game {
	outputGame := make(chan Game)
	go func() {
		for {
			g := <-input
			g.Players[g.Turn].Scores = append(g.Players[g.Turn].Scores, g.PointsEarned)
			g.Turn += 1
			if g.Turn == len(g.Players) {
				g.Turn = 0
			}
			outputGame <- g
		}
	}()
	return outputGame
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

func Sum(score []int) (total int) {
	for _, s := range score {
		total += s
	}
	return
}
