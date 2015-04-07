package main

import (
	"ai"
	//"elo"
	"eval"
	"fmt"
	"github.com/malbrecht/chess"
	"github.com/malbrecht/chess/engine/uci"
	//"log"
	"os"
)

func playGame(eng1 *uci.Engine, eng2 *uci.Engine, board *chess.Board) {
	//setup new positions
	eng1.SetPosition(board)
	eng2.SetPosition(board)

	for {
		//black
		for i := range eng1.SearchDepth(20) {
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
		board.PrintBoard()
		eng1.SetPosition(board)
		eng2.SetPosition(board)

		//white
		for i := range eng2.SearchDepth(20) {
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
		eng1.SetPosition(board)
		eng2.SetPosition(board)
		board.PrintBoard()
	}
	eng1.Quit()
	eng2.Quit()
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
				fmt.Println("white move: ", m)
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					fmt.Println("WHITE WINS")
					board.PrintBoard()
					fmt.Println(board.LegalMoves())
					fmt.Println(e1)
					os.Exit(0)
				}
			}
		}

		board.PrintBoard()
		eng1.SetPosition(board)
		eng2.SetPosition(board)
		if board.MoveNr == 50 {
			fmt.Println("DRAW")
			os.Exit(0)
		}

		//black
		for i := range eng2.SearchDepth(5, e2) {
			if m, ok := i.BestMove(); ok {
				fmt.Println("black move: ", m)
				board = board.MakeMove(m)
				if _, mate := board.IsCheckOrMate(); mate {
					fmt.Println("BLACK WINS")
					board.PrintBoard()
					fmt.Println(board.LegalMoves())
					fmt.Println(e2)
					os.Exit(0)
				}
			}
		}

		eng1.SetPosition(board)
		eng2.SetPosition(board)
		board.PrintBoard()
		if board.MoveNr == 50 {
			fmt.Println("DRAW")
			os.Exit(0)
		}
	}
}

func main() {
	board, _ := chess.ParseFen("")
	board.PrintBoard()

	//fmt.Println(chess.IsCastleMove(chess.Move{chess.E1, chess.A1, chess.NoPiece}))
	//fmt.Println(board.LegalMoves())
	//fmt.Println(board.CastleSq)

	//var log *log.Logger
	//eng1, _ := uci.Run("stockfish", nil, log)
	eng1 := ai.Engine{}
	eng2 := ai.Engine{}
	playEvalTestGame(eng1, eng2, board, eval.EvaluateBasic, eval.EvaluateWithTables)
	//playGame(eng1, eng2, board)
	//fmt.Println(elo.EstimateElo(eng2, 5))
}
