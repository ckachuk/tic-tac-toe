package tictactoe

type Player struct {
	Name string
	Wins int8
}

type Game struct {
	Player1    Player
	Player2    Player
	Board      [9]string
	PlayerTurn int8
}

func InitGame(player1Name string, player2Name string) (game *Game) {
	player1 := &Player{Name: player1Name}
	player2 := &Player{Name: player2Name}

	board := [9]string{"", "", "", "", "", "", "", "", ""}

	game = &Game{Player1: *player1, Player2: *player2, Board: board, PlayerTurn: 1}

	return game
}

func (game *Game) ResetGame() {
	board := [9]string{"", "", "", "", "", "", "", "", ""}
	game.PlayerTurn = 1
	game.Board = board
}

func CheckArrays(winningFormula *[8][3]int8, board [9]string, playerCharacter string) bool {
	count := 0
	for _, i := range winningFormula {
		count = 0
		for _, j := range i {
			for index_k, k := range board {
				if j == int8(index_k+1) && k == playerCharacter {
					count++
				}
			}
		}
		if count == 3 {
			return true
		}
	}
	return false
}

func (game *Game) CheckGameState() bool {

	winningFormula := [8][3]int8{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 5, 9}, {3, 5, 7}, {1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	if game.PlayerTurn == 1 {
		return CheckArrays(&winningFormula, game.Board, "X")
	} else if game.PlayerTurn == 2 {
		return CheckArrays(&winningFormula, game.Board, "O")
	}
	return false
}

func (game *Game) UpdateBoard(choice int8) {

	if game.PlayerTurn == 1 && game.Board[choice] == "" {
		game.Board[choice] = "X"
		game.PlayerTurn = 2
	} else if game.PlayerTurn == 2 && game.Board[choice] == "" {
		game.Board[choice] = "O"
		game.PlayerTurn = 1
	}
}

//multidimentional arrays aproach funtions
// func SearchElement(arr []int8, element int8) int8 {
// 	leftPointer := 0
// 	rightPointer := len(arr) - 1

// 	for leftPointer <= rightPointer {
// 		midPointer := int8((leftPointer + rightPointer) / 2)
// 		midValue := arr[midPointer]

// 		if midValue == element {
// 			return midPointer
// 		} else if midValue < element {
// 			leftPointer = int(midPointer + 1)
// 		} else if midValue > element {
// 			rightPointer = int(midPointer - 1)
// 		}
// 	}
// 	return -1
// }

// func RemoveElement(arr []int8, element int8) []int8 {
// 	posElement := SearchElement(arr, element)

// 	if posElement == -1 {
// 		panic("The value does not exist")
// 	}
// 	return append(arr[:posElement], arr[posElement+1:]...)
// }
