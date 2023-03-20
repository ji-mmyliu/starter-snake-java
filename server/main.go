package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {

	var board [11][11]byte
	var verbose = len(os.Args) > 1 && os.Args[1] == "verbose"

	for _, snake := range state.Board.Snakes {
		var is_you bool = snake.ID == state.You.ID
		var head Coord = snake.Head
		for j, coordinate := range snake.Body {
			if j > 0 && coordinate.X == head.X && coordinate.Y == head.Y {
				continue // Don't re-mark the head
			}

			if j == 0 {
				if is_you == true {
					board[10-coordinate.Y][coordinate.X] = 'H'
				} else {
					board[10-coordinate.Y][coordinate.X] = 'h'
				}
			} else if j == len(snake.Body)-1 {
				if is_you == true {
					board[10-coordinate.Y][coordinate.X] = 'T'
				} else {
					board[10-coordinate.Y][coordinate.X] = 't'
				}
			} else {
				if is_you == true {
					board[10-coordinate.Y][coordinate.X] = 'B'
				} else {
					board[10-coordinate.Y][coordinate.X] = 'b'
				}
			}
		}
	}

	for _, food := range state.Board.Food {
		board[10-food.Y][food.X] = 'f'
	}

	f, _ := os.Create("input.txt")

	// var output = ""
	for row := 0; row < 11; row++ {
		for col := 0; col < 11; col++ {
			if board[row][col] == 0 {
				if verbose {
					fmt.Print(".")
				}
				// output += "."
				f.Write([]byte{'.'})
			} else {
				if verbose {
					fmt.Printf("%c", board[row][col])
				}
				f.Write([]byte{board[row][col]})
			}
		}
		if verbose {
			fmt.Println()
		}
		f.Write([]byte{'\n'})
	}

	f.WriteString(fmt.Sprint(state.You.Health))

	f.Sync()

	if verbose {
		fmt.Println("----------------")
	}

	cmd := exec.Command("java", "Main")
	stdout, err := cmd.Output()

	outputs := strings.Split(string(stdout), "\n")

	if err != nil {
		fmt.Println(err.Error())
		log.Printf("MOVE %d: %s\n", state.Turn, "(error) defaulting to up")
		return BattlesnakeMoveResponse{Move: "up"}
	} else {
		log.Printf("MOVE %d: %s\n", state.Turn, outputs[0])
		if len(outputs) > 1 {
			return BattlesnakeMoveResponse{Move: outputs[0], Shout: outputs[1]}
		} else {
			return BattlesnakeMoveResponse{Move: outputs[0]}
		}
	}
}

func main() {
	RunServer()
}
