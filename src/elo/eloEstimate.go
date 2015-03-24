package elo

import (
	"github.com/malbrecht/chess"
	//"github.com/malbrecht/chess/engine/uci"
	"ai"
	//"fmt"
)

func EstimateElo(eng ai.Engine, depth int) int {
	eloRatings := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	board, _ := chess.ParseFen("r1b3k1/6p1/P1n1pr1p/q1p5/1b1P4/2N2N2/PP1QBPPP/R3K2R b - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			//estimate rating of m
			if m.From == chess.F6 && m.To == chess.F3 {
				eloRatings[0] = 2600
			} else if m.From == chess.C5 && m.To == chess.D4 {
				eloRatings[0] = 1900
			} else if m.From == chess.C6 && m.To == chess.D4 {
				eloRatings[0] = 1900
			} else if m.From == chess.B4 && m.To == chess.C3 {
				eloRatings[0] = 1400
			} else if m.From == chess.C8 && m.To == chess.A6 {
				eloRatings[0] = 1500
			} else if m.From == chess.F6 && m.To == chess.G6 {
				eloRatings[0] = 1400
			} else if m.From == chess.E6 && m.To == chess.E5 {
				eloRatings[0] = 1200
			} else if m.From == chess.C8 && m.To == chess.D7 {
				eloRatings[0] = 1600
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("2nq1nk1/5p1p/4p1pQ/pb1pP1NP/1p1P2P1/1P4N1/P4PB1/6K1 w - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.G2 && m.To == chess.E4 {
				eloRatings[1] = 2600
			} else if m.From == chess.G5 && m.To == chess.H7 {
				eloRatings[1] = 1950
			} else if m.From == chess.H5 && m.To == chess.G6 {
				eloRatings[1] = 1900
			} else if m.From == chess.G2 && m.To == chess.F1 {
				eloRatings[1] = 1400
			} else if m.From == chess.G2 && m.To == chess.D5 {
				eloRatings[1] = 1200
			} else if m.From == chess.F2 && m.To == chess.F4 {
				eloRatings[1] = 1400
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("8/3r2p1/pp1Bp1p1/1kP5/1n2K3/6R1/1P3P2/8 w - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.C5 && m.To == chess.C6 {
				eloRatings[2] = 2500
			} else if m.From == chess.G3 && m.To == chess.G6 {
				eloRatings[2] = 2000
			} else if m.From == chess.E4 && m.To == chess.E5 {
				eloRatings[2] = 1900
			} else if m.From == chess.G3 && m.To == chess.G5 {
				eloRatings[2] = 1700
			} else if m.From == chess.E4 && m.To == chess.D4 {
				eloRatings[2] = 1200
			} else if m.From == chess.D6 && m.To == chess.E5 {
				eloRatings[2] = 1200
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("8/4kb1p/2p3pP/1pP1P1P1/1P3K2/1B6/8/8 w - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.E5 && m.To == chess.E6 {
				eloRatings[3] = 2500
			} else if m.From == chess.B3 && m.To == chess.F7 {
				eloRatings[3] = 1600
			} else if m.From == chess.B3 && m.To == chess.C2 {
				eloRatings[3] = 1700
			} else if m.From == chess.B3 && m.To == chess.D1 {
				eloRatings[3] = 1800
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("b1R2nk1/5ppp/1p3n2/5N2/1b2p3/1P2BP2/4B1PP/6K1 w - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.E3 && m.To == chess.C5 {
				eloRatings[4] = 2500
			} else if m.From == chess.F5 && m.To == chess.H6 {
				eloRatings[4] = 2100
			} else if m.From == chess.E3 && m.To == chess.H6 {
				eloRatings[4] = 1900
			} else if m.From == chess.F5 && m.To == chess.G7 {
				eloRatings[4] = 1500
			} else if m.From == chess.F2 && m.To == chess.G3 {
				eloRatings[4] = 1750
			} else if m.From == chess.C8 && m.To == chess.F8 {
				eloRatings[4] = 1200
			} else if m.From == chess.F2 && m.To == chess.H4 {
				eloRatings[4] = 1200
			} else if m.From == chess.E3 && m.To == chess.B6 {
				eloRatings[4] = 1750
			} else if m.From == chess.E2 && m.To == chess.C4 {
				eloRatings[4] = 1400
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("3rr1k1/pp3pbp/2bp1np1/q3p1B1/2B1P3/2N4P/PPPQ1PP1/3RR1K1 w - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.G5 && m.To == chess.F6 {
				eloRatings[5] = 2500
			} else if m.From == chess.C3 && m.To == chess.D5 {
				eloRatings[5] = 1700
			} else if m.From == chess.C4 && m.To == chess.B5 {
				eloRatings[5] = 1900
			} else if m.From == chess.F2 && m.To == chess.F4 {
				eloRatings[5] = 1700
			} else if m.From == chess.A2 && m.To == chess.A3 {
				eloRatings[5] = 1200
			} else if m.From == chess.E1 && m.To == chess.E3 {
				eloRatings[5] = 1200
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("r1b1qrk1/1ppn1pb1/p2p1npp/3Pp3/2P1P2B/2N5/PP1NBPPP/R2Q1RK1 b - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.F6 && m.To == chess.H7 {
				eloRatings[6] = 2500
			} else if m.From == chess.F6 && m.To == chess.E4 {
				eloRatings[6] = 1800
			} else if m.From == chess.G6 && m.To == chess.G5 {
				eloRatings[6] = 1700
			} else if m.From == chess.A6 && m.To == chess.A5 {
				eloRatings[6] = 1700
			} else if m.From == chess.G8 && m.To == chess.H7 {
				eloRatings[6] = 1500
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("2R1r3/5k2/pBP1n2p/6p1/8/5P1P/2P3P1/7K w - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.B6 && m.To == chess.D8 {
				eloRatings[7] = 2500
			} else if m.From == chess.C8 && m.To == chess.E8 {
				eloRatings[7] = 1600
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("2r2rk1/1p1R1pp1/p3p2p/8/4B3/3QB1P1/q1P3KP/8 w - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.E3 && m.To == chess.D4 {
				eloRatings[8] = 2500
			} else if m.From == chess.E4 && m.To == chess.G6 {
				eloRatings[8] = 1800
			} else if m.From == chess.E4 && m.To == chess.H7 {
				eloRatings[8] = 1800
			} else if m.From == chess.E3 && m.To == chess.H6 {
				eloRatings[8] = 1700
			} else if m.From == chess.D7 && m.To == chess.B7 {
				eloRatings[8] = 1400
			} else {
				eloRatings[0] = 0
			}
		}
	}

	board, _ = chess.ParseFen("r1bq1rk1/p4ppp/1pnp1n2/2p5/2PPpP2/1NP1P3/P2B2PP/R1BQ1RK1 b - - 0 1")
	eng.SetPosition(board)
	board.PrintBoard()
	for i := range eng.SearchDepth(depth) {
		if m, ok := i.BestMove(); ok {
			if m.From == chess.D8 && m.To == chess.D7 {
				eloRatings[9] = 2600
			} else if m.From == chess.F6 && m.To == chess.E8 {
				eloRatings[9] = 2000
			} else if m.From == chess.H7 && m.To == chess.H5 {
				eloRatings[9] = 1800
			} else if m.From == chess.C5 && m.To == chess.D4 {
				eloRatings[9] = 1600
			} else if m.From == chess.C8 && m.To == chess.A6 {
				eloRatings[9] = 1800
			} else if m.From == chess.A7 && m.To == chess.A5 {
				eloRatings[9] = 1800
			} else if m.From == chess.F8 && m.To == chess.E8 {
				eloRatings[9] = 1400
			} else if m.From == chess.D6 && m.To == chess.D5 {
				eloRatings[9] = 1500
			} else {
				eloRatings[0] = 0
			}
		}
	}

	return (eloRatings[0] + eloRatings[1] + eloRatings[2] + eloRatings[3] +
		eloRatings[4] + eloRatings[5] + eloRatings[6] + eloRatings[7] +
		eloRatings[8] + eloRatings[9]) / 10
}
