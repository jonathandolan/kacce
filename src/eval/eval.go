package eval

import (
	"fmt"
	"github.com/malbrecht/chess"
)

type Eval func(b *chess.Board) int

func ScoreTable(b *chess.Board) int {
	score := 0
	for i, v := range b.Piece {
		if v == chess.WP || v == chess.BP { //pawn
			score = score + pawnTable[i]
		} else if v == chess.WN || v == chess.BN { //knight in shining armor
			score = score + knightTable[i]
		} else if v == chess.WB || v == chess.BB { //bishop
			score = score + bishopTable[i]
		} else if v == chess.WR || v == chess.WR { //rook
			score = score + rookTable[i]
		} else if v == chess.WQ || v == chess.WQ { //queen
			score = score + queenTable[i]
		} else if v == chess.WK || v == chess.WK { //king
			score = score + kingTable[i]
		}
	}

	if b.SideToMove == chess.Black {
		return -score
	} else {
		return score
	}
}

//returns balance of material. white is positive, black is negative
func Material(b *chess.Board) int {
	score := 0
	for i, _ := range b.Piece {
		if b.Piece[i] == chess.WR { //white rook
			score = score + rookValue
		} else if b.Piece[i] == chess.WN { // white knight
			score = score + knightValue
		} else if b.Piece[i] == chess.WB { // white bishop
			score = score + bishopValue
		} else if b.Piece[i] == chess.WQ { // white queen
			score = score + queenValue
		} else if b.Piece[i] == chess.WK { // white king
			score = score + kingValue
		} else if b.Piece[i] == chess.WP { // white pawn
			score = score + pawnValue
		} else if b.Piece[i] == chess.BR { //black rook
			score = score - rookValue
		} else if b.Piece[i] == chess.BN { // black knight
			score = score - knightValue
		} else if b.Piece[i] == chess.BB { // black bishop
			score = score - bishopValue
		} else if b.Piece[i] == chess.BQ { // black queen
			score = score - queenValue
		} else if b.Piece[i] == chess.BK { // black king
			score = score - kingValue
		} else if b.Piece[i] == chess.BP { // black pawn
			score = score - pawnValue
		}
	}
	return score
}

var EvaluateBasic = func(b *chess.Board) int {
	return Material(b)
}

var EvaluateWithTables = func(b *chess.Board) int {
	score := 0
	material := Material(b)
	tableScore := ScoreTable(b)
	score = material + tableScore
	return score
}

var EvaluateWithPassedPawns = func(b *chess.Board) int {
	b.PrintBoard()
	return 7
}

func EvalTest(b *chess.Board) bool {
	inputScore := EvaluateBasic(b)
	fmt.Println("inputScore: ", inputScore)

	mirror := b.MirrorBoard()
	mirror.PrintBoard()
	mirrorScore := EvaluateBasic(&mirror)
	fmt.Println("mirrorScore: ", mirrorScore)

	if mirrorScore == -inputScore {
		return true
	} else {
		return false
	}

}
