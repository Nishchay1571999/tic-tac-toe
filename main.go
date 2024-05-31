package main

import (
	"fmt"
	"game/game"
	"time"
)

func main() {

	fullGame := *game.NewGame()
	player := "X"
	for {
		fmt.Println(fullGame)
		var x, y int
		fmt.Println("PLAYER \t : ", player)
		fmt.Print("Enter Row : ")
		fmt.Scan(&x)
		fmt.Print("\nEnter Col : ")
		fmt.Scan(&y)
		pos := [2]int{x - 1, y - 1}
		if game.CheckOver(&fullGame, player) {
			fmt.Printf("Player %s gets a point\n", player)
			fmt.Println("Player X : ", game.GetPlayerWinRates(fullGame)[0])
			fmt.Println("Player O : ", game.GetPlayerWinRates(fullGame)[1])
			time.Sleep(5 * time.Second)
			game.Clear()
		}
		game.UpdateBoard(fullGame, pos, &player)
		game.Clear()
	}

}
