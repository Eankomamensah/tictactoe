package main

import (
	"fmt"
)

func main() {
	var isX bool
	played := make([]string, 9)
	// intialising elements as spaces for formating purposes
	for i := 0; i < len(played); i++ {
		played[i] = " "
	}
	printBoard()
	for {
		var keyPressed int
		if isX {
			_, err := fmt.Scanf("%d\n", &keyPressed)
			//matching keyPressed and played array index
			keyUsed := keyPressed - 1
			if err != nil {
				fmt.Println(err)
				break
			}
			Player := "X"
			if played[keyUsed] == " " {
				played[keyUsed] = Player
			} else {
				fmt.Println("There's already a character there")
				continue
			}

			//check winnig conditions for X
			if played[0] == Player && played[3] == Player && played[6] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[1] == Player && played[4] == Player && played[7] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[2] == Player && played[5] == Player && played[8] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[0] == Player && played[4] == Player && played[8] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[2] == Player && played[4] == Player && played[6] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[0] == Player && played[1] == Player && played[2] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[3] == Player && played[4] == Player && played[5] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[6] == Player && played[7] == Player && played[8] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			}
			//O/X turn indicator
			isX = false
		} else if !isX {
			_, err := fmt.Scanf("%d\n", &keyPressed)
			//matching keyPressed and played array index
			keyUsed := keyPressed - 1
			if err != nil {
				fmt.Println(err)
				break
			}
			Player := "O"
			if played[keyUsed] == " " {
				played[keyUsed] = Player
			} else {
				fmt.Println("There's already a character there")
				continue
			}

			//check winning conditions for O
			if played[0] == Player && played[3] == Player && played[6] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[1] == Player && played[4] == Player && played[7] == Player {
				fmt.Println(Player, " wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[2] == Player && played[5] == Player && played[8] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[0] == Player && played[4] == Player && played[8] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[2] == Player && played[4] == Player && played[6] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[0] == Player && played[1] == Player && played[2] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[3] == Player && played[4] == Player && played[5] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			} else if played[6] == Player && played[7] == Player && played[8] == Player {
				fmt.Println(Player, "wins! Congratulations!")
				currentBoard(played)
				break
			}
			isX = true
		}
		currentBoard(played)
		//fmt.Println(played)
		if isFull(played) {
			fmt.Println("Game Over.\nTry again?")
			main()
		}
	}
}

func currentBoard(arr []string) {
	fmt.Println("", arr[0], "|", arr[1], "|", arr[2])
	fmt.Println("---+---+---")
	fmt.Println("", arr[3], "|", arr[4], "|", arr[5])
	fmt.Println("---+---+---")
	fmt.Println("", arr[6], "|", arr[7], "|", arr[8])
}

func isFull(arr []string) bool {
	for i := 0; i < len(arr); i++ {
		place := arr[i]
		switch place {
		case "X":
			continue
		case "O":
			continue
		default:
			return false
		}
	}
	return true
}

func errorHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func printBoard() {
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("---+---+---")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("---+---+---")
	fmt.Println(" 7 | 8 | 9 ")
	fmt.Println("Enter a number where you'd like to play!")
}
