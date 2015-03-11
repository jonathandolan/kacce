package main

import (
	"elo"
	"fmt"
	"github.com/malbrecht/chess"
	"github.com/malbrecht/chess/engine/uci"
	"log"
	"os"
)

func playGame(eng1, eng2 *uci.Engine, board *chess.Board) {
	//setup new positions
	eng1.SetPosition(board)
	eng2.SetPosition(board)

	for {
		//black
		for i := range eng1.SearchDepth(20) {
			if m, ok := i.BestMove(); ok {
				fmt.Println("black move: ", m)
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					fmt.Println("BLACK WINS")
					board.PrintBoard()
					os.Exit(0)
				}
			}

		}
		board.PrintBoard()
		eng1.SetPosition(board)
		eng2.SetPosition(board)

		//white
		for i := range eng2.SearchDepth(20) {
			if m, ok := i.BestMove(); ok {
				fmt.Println("white move: ", m)
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					fmt.Println("WHITE WINS")
					board.PrintBoard()
					os.Exit(0)
				}
			}

		}
		eng1.SetPosition(board)
		eng2.SetPosition(board)
		board.PrintBoard()
	}
}

func main() {
	board, _ := chess.ParseFen("")
	board.PrintBoard()
	var log *log.Logger

	eng1, _ := uci.Run("stockfish", nil, log)
	//eng2, _ := uci.Run("stockfish", nil, log)

	fmt.Println(elo.EstimateElo(eng1, 20))
}
