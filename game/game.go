package game

import (
	"fmt"
)

type Board [3][3]string

type Game struct {
	x_score int
	y_score int
	board   *Board
}

func (g Game) String() string {
	boardString := ""
	for _, row := range g.board {
		for _, cell := range row {
			if cell == "" {
				boardString += "| # |"
			} else {
				boardString += fmt.Sprintf("| %s |", cell)
			}
		}
		boardString += "\n"
	}

	return fmt.Sprintf("\n%s", boardString)
}

func NewBoard() *Board {
	// Define the dimensions of the 2D array
	rows := 3
	cols := 3

	board := Board{}
	// Assign empty strings to elements of the array
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = ""
		}
	}
	return &board
}

func NewGame() *Game {
	game := Game{}
	game.board = NewBoard()
	game.x_score = 0
	game.y_score = 0
	return &game
}
func UpdateBoard(g Game, pos [2]int, player *string) {
	temp := Board{}
	// Define the dimensions of the 2D array
	rows := 3
	cols := 3
	if pos[0] < 3 && pos[1] < 3 {
		for row := 0; row < rows; row++ {
			for cell := 0; cell < cols; cell++ {
				if row == pos[0] && cell == pos[1] && g.board[row][cell] == "" {
					temp[row][cell] = *player
					if *player == "X" {
						*player = "O"
					} else {
						*player = "X"
					}
				} else {
					temp[row][cell] = g.board[row][cell]
				}
			}
		}
		*g.board = temp
	}
}

func CheckOver(g *Game, player string) bool {
	ans := false
	for row := 0; row < 3; row++ {
		if g.board[row][0] == g.board[row][1] && g.board[row][1] == g.board[row][2] && g.board[row][2] != "" {
			ans = true
		}
	}
	for col := 0; col < 3; col++ {
		if g.board[0][col] == g.board[1][col] && g.board[1][col] == g.board[2][col] && g.board[2][col] != "" {
			ans = true
		}
	}
	if g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] && g.board[2][2] != "" {
		ans = true
	}
	if g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] && g.board[2][0] != "" {
		ans = true
	}
	if ans {
		if player == "X" {
			g.x_score++
		} else {
			g.y_score++
		}
		Clear()
	}
	return ans
}

func GetPlayerWinRates(g Game) [2]int {
	scores := [2]int{g.x_score, g.y_score}
	return scores
}
