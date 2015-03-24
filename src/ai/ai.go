package ai

import (
	"eval"
	"fmt"
	"github.com/malbrecht/chess"
	"github.com/malbrecht/chess/engine"
	"time"
)

type Engine struct {
	board *chess.Board
}

func (k *Engine) SetPosition(board *chess.Board) {
	k.board = board
}

//only search implimented
//starts recursive function minimax
func (k *Engine) SearchDepth(depth int) <-chan engine.Info {
	//result channel
	infoChan := make(chan engine.Info, 1)
	//local info struct
	info := Info{}
	_, tempMove := minimaxAB(depth, k.board, k.board.SideToMove, -9223372036854775808, 9223372036854775807)
	info.chosenMove = tempMove
	fmt.Println(info.BestMove())
	infoChan <- info
	close(infoChan)
	return infoChan
}

//minimax search algorithm
func minimax(depth int, board *chess.Board, mySide int) (score int, move chess.Move) {
	_, mate := board.IsCheckOrMate()
	if depth == 0 || mate == true {
		return eval.Evaluate(board), chess.NullMove
	} else {
		if depth%2 == 0 { //max
			bestMaxScore := -9223372036854775808
			bestMaxMove := chess.Move{}
			for _, m := range board.LegalMoves() {
				boardCopy := board.CopyBoard()
				boardCopy = boardCopy.MakeMove(m)
				tempScore, _ := minimax(depth-1, boardCopy, mySide)
				if tempScore > bestMaxScore {
					bestMaxScore = tempScore
					bestMaxMove = m
				}
			}
			return bestMaxScore, bestMaxMove

		} else { //min
			bestMinScore := 9223372036854775807
			bestMinMove := chess.Move{}
			for _, m := range board.LegalMoves() {
				boardCopy := board.CopyBoard()
				boardCopy = boardCopy.MakeMove(m)
				tempScore, _ := minimax(depth-1, boardCopy, mySide)
				if tempScore < bestMinScore {
					bestMinScore = tempScore
					bestMinMove = m
				}
			}
			return bestMinScore, bestMinMove
		}
	}
}

//minimax search algorithm with alpha, beta pruning
func minimaxAB(depth int, board *chess.Board, mySide int, alpha int, beta int) (score int, move chess.Move) {
	_, mate := board.IsCheckOrMate()
	if depth == 0 || mate == true {
		return eval.Evaluate(board), chess.NullMove
	} else {
		if depth%2 == 0 { //max
			bestMaxMove := chess.Move{}
			for _, m := range board.LegalMoves() {
				boardCopy := board.CopyBoard()
				boardCopy = boardCopy.MakeMove(m)
				tempScore, _ := minimaxAB(depth-1, boardCopy, mySide, alpha, beta)
				if tempScore > alpha {
					alpha = tempScore
					bestMaxMove = m
					if alpha >= beta { //aplha beta pruning
						break
					}
				}
			}
			return alpha, bestMaxMove

		} else { //min
			bestMinMove := chess.Move{}
			for _, m := range board.LegalMoves() {
				boardCopy := board.CopyBoard()
				boardCopy = boardCopy.MakeMove(m)
				tempScore, _ := minimaxAB(depth-1, boardCopy, mySide, alpha, beta)
				if tempScore < beta {
					beta = tempScore
					bestMinMove = m
					if alpha >= beta { //alpha beta pruning
						break
					}
				}
			}
			return beta, bestMinMove
		}
	}
}

func (k *Engine) SearchTime(t time.Duration) <-chan engine.Info {
	fmt.Println("not implimented")
	return nil
}

func (k *Engine) SearchClock(wtime, btime, winc, binc time.Duration, movesToGo int) <-chan engine.Info {
	fmt.Println("not implimented")
	return nil
}

func (k *Engine) Stop() {
	fmt.Println("not implimented")
}

func (k *Engine) Quit() {
	fmt.Println("not implimented")
}

func (k *Engine) Ping() error {
	fmt.Println("not implimented")
	return nil
}

// Options returns the settable options of the engine.
func (k *Engine) Options() map[string]engine.Option {
	fmt.Println("not implimented")
	return nil
}

func (k *Engine) Search() <-chan engine.Info {
	fmt.Println("not implimented")
	return nil
}

type Info struct {
	chosenMove chess.Move
}

func (ki Info) Err() error {
	return nil
}

func (ki Info) BestMove() (move chess.Move, ok bool) {
	return ki.chosenMove, true
}

func (ki Info) Pv() *engine.Pv {
	return nil
}

func (ki Info) Stats() *engine.Stats {
	return nil
}
