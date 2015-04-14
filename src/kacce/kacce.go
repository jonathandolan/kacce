//Main method for the kacce anti computer chess engine.
//Written by Jonathan Dolan
//Spring 2015

package main

import (
	"ai"
	"bufio"
	"elo"
	"eval"
	"flag"
	"fmt"
	"github.com/malbrecht/chess"
	"github.com/malbrecht/chess/engine/uci"
	"log"
	"os"
	"strconv"
	"strings"
)

//Plays a single game, between a uci chess engine, and kacce's AI.
//Parameters: uci = external chess engine, ai = internal ai, e = evaluation function for ai, board = board to use
//No return value
func playGame(uci *uci.Engine, ai ai.Engine, e eval.Eval, board *chess.Board) {
	//setup new positions
	uci.SetPosition(board)
	ai.SetPosition(board)

	for {
		//white
		for i := range uci.SearchDepth(5) {
			if m, ok := i.BestMove(); ok {
				fmt.Println("white move: ", m)
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					fmt.Println("WHITE WINS")
					board.PrintBoard(true)
					os.Exit(0)
				}
			}

		}
		board.PrintBoard(true)
		uci.SetPosition(board)
		ai.SetPosition(board)

		//black
		for i := range ai.SearchDepth(5, e) {
			if m, ok := i.BestMove(); ok {
				fmt.Println("black move: ", m)
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					fmt.Println("BLACK WINS")
					board.PrintBoard(true)
					os.Exit(0)
				}
			}

		}
		uci.SetPosition(board)
		ai.SetPosition(board)
		board.PrintBoard(true)
	}
	uci.Quit()
	ai.Quit()
}

//Plays a single game between two ai's with different evaluation functions.
//Used to test which evaluation function is better.
//Parameters: eng1, eng2 = the two AI objects to play each other, board = board to use, e1, e2 = evaluation functions
//for eng1, and eng2.
//No return value
func playEvalTestGame(eng1 ai.Engine, eng2 ai.Engine, board *chess.Board, e1 eval.Eval, e2 eval.Eval) {
	board.SideToMove = 0

	//setup new positions
	eng1.SetPosition(board)
	eng2.SetPosition(board)

	for {
		//white
		for i := range eng1.SearchDepth(5, e1) {
			if m, ok := i.BestMove(); ok {
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					board.PrintBoard(true)
					fmt.Println(board.LegalMoves())
					viewHistory(board, true, "WHITE WINS")
				}
			}
		}

		board.PrintBoard(true)
		eng1.SetPosition(board)
		eng2.SetPosition(board)
		if board.MoveNr == 50 {
			viewHistory(board, true, "Draw: 50 move rule")
		}

		//black
		for i := range eng2.SearchDepth(5, e2) {
			if m, ok := i.BestMove(); ok {
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					board.PrintBoard(true)
					fmt.Println(board.LegalMoves())
					viewHistory(board, true, "BLACK WINS")
				}
			}
		}

		eng1.SetPosition(board)
		eng2.SetPosition(board)
		board.PrintBoard(true)
		if board.MoveNr == 50 {
			viewHistory(board, true, "Draw: 50 move rule")
		}
	}
}

//Views the history of a chess board.
//Parameters: board = board whose history will be viewed, includeEval = boolean determines whether to include evaluations of each position.
//resultString = string containing the result of the game, ie: Black won, White won, Draw, etc.
func viewHistory(board *chess.Board, includeEval bool, resultString string) {
	var input rune
	index := len(board.History) - 1
	tempBoard := board
	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Print("\033[1;1H")
		//fmt.Print("\033[0J")
		tempBoard.PrintBoard(false)
		if includeEval {
			fmt.Println("Basic: ", eval.EvaluateBasic(tempBoard), " With tables: ", eval.EvaluateWithTables(tempBoard))
		}
		fmt.Println("Options: a: move back one ply, d: move backward one ply, q: quit")
		if index == -1 {
			fmt.Println("Beginning of game!")
		} else if index == len(board.History)-1 {
			fmt.Println("End of game!")
			fmt.Println(resultString)
		}
		input, _, _ = reader.ReadRune()
		if input == 'q' {
			quit(board)
		} else if input == 'a' {
			if index == -1 { //reset index so doesn't run of end of History at first of game
				index++
			}
			index-- //adjust index
			if index == -1 {
				tempBoard, _ = chess.ParseFen("") //index of -1 means initial position (not recorded in History)
			} else {
				tempBoard, _ = chess.ParseFen(board.History[index])
			}

		} else if input == 'd' {
			if index == len(board.History)-1 { //reset index
				index--
			}
			index++ //adjust index
			if index == -1 {
				tempBoard, _ = chess.ParseFen(board.History[index+1])
			} else {
				tempBoard, _ = chess.ParseFen(board.History[index])
			}

		}
	}
}

//Quits after viewHistory, and saves game if required
//Parameters: board = board to save if necessary
func quit(board *chess.Board) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like to save this game? Y/N")
	save, _, _ := reader.ReadRune()
	if save == 'Y' || save == 'y' {
		fmt.Println("Please enter a name for the file.")
		reader.ReadRune() //clear \ng character still in buffer.
		fileName, _ := reader.ReadString('\n')
		saveGame(board, fileName)
		os.Exit(0)
	} else {
		os.Exit(0)
	}
}

//Helper method for quit. Saves game as a text file.
//Parameters: board = board to save, name = file name
func saveGame(board *chess.Board, name string) {
	name = strings.Trim(name, "\n")
	file, _ := os.Create(name)
	writer := bufio.NewWriter(file)
	fmt.Println("Name: ", name)
	for i, v := range board.History {
		writer.WriteString(strconv.Itoa(i+1) + ": ")
		writer.WriteString(v)
		writer.WriteString("\n")
	}
	writer.Flush()
	file.Close()

}

//Opens a previously said game and views it.
func openSavedGame() {
	reader := bufio.NewReader(os.Stdin)
	board, _ := chess.ParseFen("")
	fmt.Println("History: ", board.History)
	//i := 0
	fmt.Println("Please enter the name of the file to open.")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	file, _ := os.Open(input)
	gameReader := bufio.NewReader(file)
	for {
		line, _, err := gameReader.ReadLine()
		if err != nil { // end of file
			break
		}
		fenString := string(line)[3:]
		board.History = append(board.History, fenString)
	}
	viewHistory(board, true, "")

}

//Starts new game
func startNewGame() {
	board, _ := chess.ParseFen("")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like kacce vs kaccee (kvk), or kacce vs stockfish (kvs)?")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	if input == "kvk" {
		playEvalTestGame(ai.Engine{}, ai.Engine{}, board, eval.EvaluateBasic, eval.EvaluateWithTables)
	} else if input == "kvs" {
		var log *log.Logger
		eng1, _ := uci.Run("stockfish", nil, log)
		eng2 := ai.Engine{}
		playGame(eng1, eng2, eval.EvaluateWithTables, board)
	}
}

//Plays a game of one human vs. another human
func humanVsHuman() {
	board, _ := chess.ParseFen("")
	reader := bufio.NewReader(os.Stdin)
	move := chess.Move{}

	board.PrintBoard(false)

	for {
		move = humanTurn(board, reader)
		if !contains(board.LegalMoves(), move) {
			fmt.Println("Invalid move!")
			humanTurn(board, reader)
		}

		board = board.MakeMove(move)
		board.PrintBoard(false)

		if _, mate := board.IsCheckOrMate(); mate {
			if board.SideToMove == 0 {
				fmt.Println("White wins!")
			} else {
				fmt.Println("Black wins!")
			}
		}

		if board.Rule50 >= 50 {
			fmt.Println("Draw: 50 move rule")
		}
	}

}

//Helper method to determine if a human move is valid.
func contains(arr []chess.Move, element chess.Move) bool {
	for _, v := range arr {
		if v.From == element.From && v.To == element.To {
			return true
		}
	}
	return false
}

//Plays a single human turn.
//Parameters: board = current position, reader = object to read input from keyboard
func humanTurn(board *chess.Board, reader *bufio.Reader) chess.Move {
	answer := chess.Move{}
	if board.SideToMove == 0 {
		fmt.Println("Side to move: White")
	} else {
		fmt.Println("Side to move: Black")
	}

	fmt.Println("Please enter your start square (a-h, 1-8 ie: a1, g6, etc.)") //like a g6
	startSquare, _ := reader.ReadString('\n')
	startSquare = startSquare[:len(startSquare)-1]
	answer.From = chess.SquareFromString(startSquare)
	if answer.From == chess.NoSquare {
		fmt.Println("Invalid square.")
		humanTurn(board, reader)
	}

	fmt.Println("Please enter your destination square")
	destSquare, _ := reader.ReadString('\n')
	destSquare = destSquare[:len(destSquare)-1]
	answer.To = chess.SquareFromString(destSquare)
	if answer.To == chess.NoSquare {
		fmt.Println("Invalid square.")
		humanTurn(board, reader)
	}
	return answer
}

//Method that tests an evaluation function and displays its elo score
func testElo() {
	//board, _ := chess.ParseFen("")
	reader := bufio.NewReader(os.Stdin)
	var input rune
	for {
		fmt.Println("Which Evaluation function would you like to test?")
		fmt.Println("	1: Basic (just material balance)")
		fmt.Println("	2: With tables (uses table values, for only one side")
		fmt.Println("	3: With tables for both sides")
		fmt.Println()
		fmt.Println("	e: exit")
		input, _, _ = reader.ReadRune()
		reader.ReadRune() //clear \n character from buffer
		if input == '1' {
			fmt.Println("basic")
			fmt.Println("Estimated ELO score: ", elo.EstimateElo(ai.Engine{}, 5, eval.EvaluateBasic))
			fmt.Println("Test another evaluation function? Y/N")
			testAgain, _, _ := reader.ReadRune()
			reader.ReadRune() //clear \n character from buffer
			if testAgain == 'y' || testAgain == 'Y' {
				continue
			} else if testAgain == 'n' || testAgain == 'N' {
				break
			}
		} else if input == '2' {
			fmt.Println("tables")
			fmt.Println("Estimated ELO score: ", elo.EstimateElo(ai.Engine{}, 5, eval.EvaluateWithTables))
			fmt.Println("Test another evaluation function? Y/N")
			testAgain, _, _ := reader.ReadRune()
			reader.ReadRune() //clear \n character from buffer
			if testAgain == 'y' || testAgain == 'Y' {
				continue
			} else if testAgain == 'n' || testAgain == 'N' {
				break
			}
		} else if input == '3' {
			fmt.Println("mirror")
			fmt.Println("Estimated ELO score: ", elo.EstimateElo(ai.Engine{}, 5, eval.EvaluateWithMirrorTables))
			fmt.Println("Test another evaluation function? Y/N")
			testAgain, _, _ := reader.ReadRune()
			reader.ReadRune() //clear \n character from buffer
			if testAgain == 'y' || testAgain == 'Y' {
				continue
			} else if testAgain == 'n' || testAgain == 'N' {
				break
			}
		} else if input == 'e' {
			break
		}
	}
}

func main() {
	newGame := flag.Bool("ng", false, "This option starts a new game")
	openGame := flag.Bool("og", false, "This option opens an existing game")
	//humanGame := flag.Bool("hg", false, "This option starts a human vs. human game")
	eloTest := flag.Bool("et", false, "This option starts an ELO estimation of a chess engine")
	flag.Parse()
	if *newGame {
		startNewGame()
	} else if *openGame {
		openSavedGame()
	} else if *eloTest {
		testElo()
	}
}
