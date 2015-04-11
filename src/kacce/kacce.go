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
					board.PrintBoard(false)
					os.Exit(0)
				}
			}

		}
		board.PrintBoard(false)
		uci.SetPosition(board)
		ai.SetPosition(board)

		//black
		for i := range ai.SearchDepth(5, e) {
			if m, ok := i.BestMove(); ok {
				fmt.Println("black move: ", m)
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					fmt.Println("BLACK WINS")
					board.PrintBoard(false)
					os.Exit(0)
				}
			}

		}
		uci.SetPosition(board)
		ai.SetPosition(board)
		board.PrintBoard(false)
	}
	uci.Quit()
	ai.Quit()
}

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

func startNewGame() {
	board, _ := chess.ParseFen("")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like kacce vs kaccee (kvk), or kacce vs stockfish (kvs)?")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	if input == "kvk" {
		playEvalTestGame(ai.Engine{}, ai.Engine{}, board, eval.EvaluateBasic, eval.EvaluateWithTables)
	} else if input == "kvs" {
		fmt.Println("SDKFJH")
		var log *log.Logger
		eng1, _ := uci.Run("stockfish", nil, log)
		eng2 := ai.Engine{}
		playGame(eng1, eng2, eval.EvaluateWithTables, board)
	}
}

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
