package ai

import (
	"eval"
	"fmt"
	"github.com/malbrecht/chess"
	"github.com/malbrecht/chess/engine"
	"time"
)

type KacceAI struct {
	board *chess.Board
}

func (k *KacceAI) SetPosition(board *chess.Board) {
	k.board = board
}

//only search implimented
//starts recursive function minimax
func (k *KacceAI) SearchDepth(depth int) <-chan engine.Info {
	//result channel
	infoChan := make(chan engine.Info, 1)
	//local info struct
	info := Info{}
	_, tempMove := minimax(depth, k.board)
	info.chosenMove = tempMove
	fmt.Println(info.BestMove())
	infoChan <- info
	close(infoChan)
	return infoChan
}

//minimax search algorithm
func minimax(depth int, board *chess.Board) (score int, move chess.Move) {
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
				tempScore, _ := minimax(depth-1, boardCopy)
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
				tempScore, _ := minimax(depth-1, boardCopy)
				if tempScore < bestMinScore {
					bestMinScore = tempScore
					bestMinMove = m
				}
			}
			return bestMinScore, bestMinMove
		}
	}
}

func (k *KacceAI) SearchTime(t time.Duration) <-chan engine.Info {
	fmt.Println("not implimented")
	return nil
}

func (k *KacceAI) SearchClock(wtime, btime, winc, binc time.Duration, movesToGo int) <-chan engine.Info {
	fmt.Println("not implimented")
	return nil
}

func (k *KacceAI) Stop() {
	fmt.Println("not implimented")
}

func (k *KacceAI) Quit() {
	fmt.Println("not implimented")
}

func (k *KacceAI) Ping() error {
	fmt.Println("not implimented")
	return nil
}

// Options returns the settable options of the engine.
func (k *KacceAI) Options() map[string]engine.Option {
	fmt.Println("not implimented")
	return nil
}

func (k *KacceAI) Search() <-chan engine.Info {
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
