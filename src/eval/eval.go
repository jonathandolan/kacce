package eval

import (
	//"fmt"
	"github.com/malbrecht/chess"
)

type Eval func(b *chess.Board) int

func GetEval() Eval {
	var e Eval

	e = func(b *chess.Board) int {
		return 12
	}

	return e
}

//tables with relative score for a piece on that square.
//from whites perspective
var pawnTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	20, 20, 20, 30, 30, 20, 20, 20,
	10, 10, 10, 20, 20, 10, 10, 10,
	5, 5, 5, 10, 10, 5, 5, 5,
	0, 0, 10, 20, 20, 10, 0, 0,
	5, 0, 0, 5, 5, 0, 0, 5,
	10, 10, 0, -10, -10, 0, 10, 10,
	0, 0, 0, 0, 0, 0, 0, 0}

var knightTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	5, 10, 10, 20, 20, 10, 10, 5,
	5, 10, 15, 20, 20, 15, 10, 5,
	0, 0, 10, 20, 20, 10, 5, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	0, -10, 0, 0, 0, 0, -10, 0}

var bishopTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, -10, 0, 0, -10, 0, 0}

var rookTable = []int{
	0, 0, 5, 10, 10, 5, 0, 0,
	25, 25, 25, 25, 25, 25, 25, 25,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0}

var queenTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0}

var kingTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0}

//tables are from whites perspective.
// mirror provides indexes for lookups from blacks point of view.
var mirror = []int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7}

func ScoreTable(b *chess.Board) int {
	score := 0
	for i, _ := range b.Piece {
		if b.Piece[i] == chess.WP { //white pawn
			score = score + pawnTable[i]
		} else if b.Piece[i] == chess.WN { //white knight in shining armor
			score = score + knightTable[i]
		} else if b.Piece[i] == chess.WB { //white bishop
			score = score + bishopTable[i]
		} else if b.Piece[i] == chess.WR { //white rook
			score = score + rookTable[i]
		} else if b.Piece[i] == chess.WQ { //white queen
			score = score + queenTable[i]
		} else if b.Piece[i] == chess.WK { //white king
			score = score + kingTable[i]
		} else if b.Piece[i] == chess.BP { //black pawn
			score = score - pawnTable[mirror[i]]
		} else if b.Piece[i] == chess.BN { //black knight
			score = score - knightTable[mirror[i]]
		} else if b.Piece[i] == chess.BB { //black bishop
			score = score - bishopTable[mirror[i]]
		} else if b.Piece[i] == chess.BR { //black rook
			score = score - rookTable[mirror[i]]
		} else if b.Piece[i] == chess.BQ { //black queen
			score = score - queenTable[mirror[i]]
		} else if b.Piece[i] == chess.BK { //black king
			score = score - kingTable[mirror[i]]
		}

	}

	return score
}

//returns balance of material. white is positive, black is negative
func Material(b *chess.Board) int {
	score := 0
	for i, _ := range b.Piece {
		if b.Piece[i] == chess.WR { //white rook
			score = score + 525
		} else if b.Piece[i] == chess.WN { // white knight
			score = score + 350
		} else if b.Piece[i] == chess.WB { // white bishop
			score = score + 350
		} else if b.Piece[i] == chess.WQ { // white queen
			score = score + 1000
		} else if b.Piece[i] == chess.WK { // white king
			score = score + 10000
		} else if b.Piece[i] == chess.WP { // white pawn
			score = score + 100
		} else if b.Piece[i] == chess.BR { //black rook
			score = score - 525
		} else if b.Piece[i] == chess.BN { // black knight
			score = score - 350
		} else if b.Piece[i] == chess.BB { // black bishop
			score = score - 350
		} else if b.Piece[i] == chess.BQ { // black queen
			score = score - 1000
		} else if b.Piece[i] == chess.BK { // black king
			score = score - 10000
		} else if b.Piece[i] == chess.BP { // black pawn
			score = score - 100
		}
	}
	return score
}

func Evaluate(b *chess.Board) int {
	score := 0
	material := Material(b)
	//fmt.Println("material: ", material)
	tableScore := ScoreTable(b)
	//fmt.Println("tableScore: ", tableScore)
	score = material + tableScore
	//fmt.Println("score: ", score)

	//return positive from perspective of player
	//	if b.SideToMove == chess.White {
	//		return score
	//	} else {
	//		return -score
	//	}
	return score
}
