// Package movegen contains all the move-generating functions. It makes
// heavy use of the GADDAG.
// Implementation notes:
// - Is the specification in the paper a bit buggy? Basically, if I assume
// an anchor is the leftmost tile of a word, the way the algorithm works,
// it will create words blindly. For example, if I have a word FIRE on the
// board, and I have the letter E on my rack, and I specify F as the anchor,
// it will create the word EF! (Ignoring the fact that IRE is on the board)
// You can see this by just stepping through the algorithm.
// It seems that anchors can only be on the rightmost tile of a word
package movegen

import (
	"sort"

	"github.com/domino14/macondo/alphabet"
	"github.com/domino14/macondo/board"
	"github.com/domino14/macondo/gaddag"
	"github.com/domino14/macondo/lexicon"
	"github.com/domino14/macondo/move"
)

// GordonGenerator is the main move generation struct. It implements
// Steven A. Gordon's algorithm from his paper "A faster Scrabble Move Generation
// Algorithm"
type GordonGenerator struct {
	gaddag gaddag.SimpleGaddag
	board  board.GameBoard
	// curRow is the current row for which we are generating moves. Note
	// that we are always thinking in terms of rows, and columns are the
	// current anchor column. In order to generate vertical moves, we just
	// transpose the `board`.
	curRowIdx     int
	curAnchorCol  int
	lastAnchorCol int

	vertical bool // Are we generating moves vertically or not?

	tilesPlayed        int
	plays              []*move.Move
	bag                *lexicon.Bag
	numPossibleLetters int
	leaveMap           LeaveMap
}

func newGordonGenHardcode(gd gaddag.SimpleGaddag) *GordonGenerator {
	// Can change later
	dist := lexicon.EnglishLetterDistribution()
	bag := dist.MakeBag(gd.GetAlphabet())
	gen := &GordonGenerator{
		gaddag:             gd,
		board:              board.MakeBoard(board.CrosswordGameBoard),
		bag:                bag,
		numPossibleLetters: int(gd.GetAlphabet().NumLetters()),
		leaveMap:           loadLeaves(gd.LexiconName()),
	}
	gen.board.SetAllCrosses()
	return gen
}

// NewGordonGenerator returns a Gordon move generator.
func NewGordonGenerator(gd gaddag.SimpleGaddag, bag *lexicon.Bag,
	board board.GameBoard) *GordonGenerator {

	gen := &GordonGenerator{
		gaddag:             gd,
		board:              board,
		bag:                bag,
		numPossibleLetters: int(gd.GetAlphabet().NumLetters()),
		leaveMap:           loadLeaves(gd.LexiconName()),
	}
	gen.board.SetAllCrosses()
	return gen
}

// GenAll generates all moves on the board. It assumes anchors have already
// been updated, as well as cross-sets / cross-scores.
func (gen *GordonGenerator) GenAll(rack *Rack) {
	dim := gen.board.Dim()
	gen.plays = []*move.Move{}
	orientations := []board.BoardDirection{
		board.HorizontalDirection, board.VerticalDirection}
	// Once for each orientation
	for idx, dir := range orientations {
		gen.vertical = idx%2 != 0
		for row := 0; row < dim; row++ {
			gen.curRowIdx = row
			// A bit of a hack. Set this to a large number at the beginning of
			// every loop
			gen.lastAnchorCol = 100
			for col := 0; col < dim; col++ {
				if gen.board.IsAnchor(row, col, dir) {
					gen.curAnchorCol = col
					gen.Gen(col, alphabet.MachineWord([]alphabet.MachineLetter{}),
						rack, gen.gaddag.GetRootNodeIndex())
					gen.lastAnchorCol = col
				}
			}
		}
		gen.board.Transpose()
	}
	gen.dedupeAndSortPlays()
}

// Gen is an implementation of the Gordon Gen function.
func (gen *GordonGenerator) Gen(col int, word alphabet.MachineWord, rack *Rack,
	nodeIdx uint32) {

	var csDirection board.BoardDirection
	// If a letter L is already on this square, then GoOn...
	curSquare := gen.board.GetSquare(gen.curRowIdx, col)
	curLetter := curSquare.Letter()

	if gen.vertical {
		csDirection = board.HorizontalDirection
	} else {
		csDirection = board.VerticalDirection
	}
	crossSet := gen.board.GetCrossSet(gen.curRowIdx, col, csDirection)

	if !curSquare.IsEmpty() {
		nnIdx := gen.gaddag.NextNodeIdx(nodeIdx, curLetter.Unblank())
		gen.GoOn(col, curLetter, word, rack, nnIdx, nodeIdx)

	} else if !rack.empty {
		for ml := alphabet.MachineLetter(0); ml < alphabet.MachineLetter(gen.numPossibleLetters); ml++ {
			if rack.LetArr[ml] == 0 {
				continue
			}
			if crossSet.Allowed(ml) {
				nnIdx := gen.gaddag.NextNodeIdx(nodeIdx, ml)
				rack.take(ml)
				gen.tilesPlayed++
				gen.GoOn(col, ml, word, rack, nnIdx, nodeIdx)
				rack.add(ml)
				gen.tilesPlayed--
			}

		}

		if rack.LetArr[alphabet.BlankMachineLetter] > 0 {
			// It's a blank. Loop only through letters in the cross-set.
			for i := 0; i < gen.numPossibleLetters; i++ {
				if crossSet.Allowed(alphabet.MachineLetter(i)) {
					nnIdx := gen.gaddag.NextNodeIdx(nodeIdx, alphabet.MachineLetter(i))
					rack.take(alphabet.BlankMachineLetter)
					gen.tilesPlayed++
					gen.GoOn(col, alphabet.MachineLetter(i).Blank(), word, rack, nnIdx, nodeIdx)
					rack.add(alphabet.BlankMachineLetter)
					gen.tilesPlayed--
				}
			}
		}

	}
}

// GoOn is an implementation of the Gordon GoOn function.
func (gen *GordonGenerator) GoOn(curCol int, L alphabet.MachineLetter, word alphabet.MachineWord,
	rack *Rack, newNodeIdx uint32, oldNodeIdx uint32) {

	if curCol <= gen.curAnchorCol {
		if !gen.board.GetSquare(gen.curRowIdx, curCol).IsEmpty() {
			word = append([]alphabet.MachineLetter{alphabet.PlayedThroughMarker}, word...)
		} else {
			word = append([]alphabet.MachineLetter{L}, word...)
		}
		// if L on OldArc and no letter directly left, then record play.
		// roomToLeft is true unless we are right at the edge of the board.
		//roomToLeft := true
		noLetterDirectlyLeft := curCol == 0 ||
			gen.board.GetSquare(gen.curRowIdx, curCol-1).IsEmpty()

		// Check to see if there is a letter directly to the left.
		if gen.gaddag.InLetterSet(L, oldNodeIdx) && noLetterDirectlyLeft && gen.tilesPlayed > 0 {
			gen.RecordPlay(word, curCol, rack)
		}
		if newNodeIdx == 0 {
			return
		}
		// Keep generating prefixes if there is room to the left, and don't
		// revisit an anchor we just saw.
		// This seems to work because we always shift direction afterwards, so we're
		// only looking at the first of a consecutive set of anchors going backwards,
		// and then always looking forward from then on.
		if curCol > 0 && curCol-1 != gen.lastAnchorCol {
			gen.Gen(curCol-1, word, rack, newNodeIdx)
		}
		// Then shift direction.
		// Get the index of the SeparationToken
		separationNodeIdx := gen.gaddag.NextNodeIdx(newNodeIdx, alphabet.SeparationMachineLetter)
		// Check for no letter directly left AND room to the right (of the anchor
		// square)
		if separationNodeIdx != 0 && noLetterDirectlyLeft && gen.curAnchorCol < gen.board.Dim()-1 {
			gen.Gen(gen.curAnchorCol+1, word, rack, separationNodeIdx)
		}

	} else {
		if !gen.board.GetSquare(gen.curRowIdx, curCol).IsEmpty() {
			word = append(word, alphabet.PlayedThroughMarker)
		} else {
			word = append(word, L)
		}

		noLetterDirectlyRight := curCol == gen.board.Dim()-1 ||
			gen.board.GetSquare(gen.curRowIdx, curCol+1).IsEmpty()
		if gen.gaddag.InLetterSet(L, oldNodeIdx) && noLetterDirectlyRight && gen.tilesPlayed > 0 {
			gen.RecordPlay(word, curCol-len(word)+1, rack)
		}
		if newNodeIdx != 0 && curCol < gen.board.Dim()-1 {
			// There is room to the right
			gen.Gen(curCol+1, word, rack, newNodeIdx)
		}
	}
}

// RecordPlay records a play.
func (gen *GordonGenerator) RecordPlay(word alphabet.MachineWord, startCol int,
	rack *Rack) {
	row := gen.curRowIdx
	col := startCol
	if gen.vertical {
		// We flip it here because we only generate vertical moves when we transpose
		// the board, so the row and col are actually transposed.
		row, col = col, row
	}
	coords := move.ToBoardGameCoords(row, col, gen.vertical)
	wordCopy := make([]alphabet.MachineLetter, len(word))
	copy(wordCopy, word)

	leave := rack.TilesOn(gen.numPossibleLetters)
	alph := gen.gaddag.GetAlphabet()
	play := move.NewScoringMove(gen.scoreMove(word, startCol),
		wordCopy, leave, gen.vertical,
		gen.tilesPlayed, alph, row, col, coords)

	play.SetEquity(calculateEquity(play, gen.leaveMap, gen.board.IsEmpty()))
	gen.plays = append(gen.plays, play)
}

func (gen *GordonGenerator) dedupeAndSortPlays() {
	dupeMap := map[int]*move.Move{}

	i := 0 // output index

	for _, m := range gen.plays {
		if m.TilesPlayed() == 1 {
			uniqKey := m.UniqueSingleTileKey()
			if _, ok := dupeMap[uniqKey]; !ok {
				dupeMap[uniqKey] = m
				gen.plays[i] = m
				i++
			}
		} else {
			gen.plays[i] = m
			i++
		}
	}
	// Everything after element i is duplicate plays.
	gen.plays = gen.plays[:i]

	sort.Slice(gen.plays, func(i, j int) bool {
		return gen.plays[i].Equity() > gen.plays[j].Equity()
		// if gen.plays[i].Score() != gen.plays[j].Score() {
		// 	return gen.plays[i].Score() > gen.plays[j].Score()
		// }
		// return gen.plays[i].BoardCoords() < gen.plays[j].BoardCoords()
	})

	if len(gen.plays) == 0 {
		gen.plays = append(gen.plays, move.NewPassMove())
	}
}

func (gen *GordonGenerator) crossDirection() board.BoardDirection {
	if gen.vertical {
		return board.HorizontalDirection
	}
	return board.VerticalDirection
}

func (gen *GordonGenerator) scoreMove(word alphabet.MachineWord, col int) int {
	dir := gen.crossDirection()
	var ls int

	mainWordScore := 0
	crossScores := 0
	bingoBonus := 0
	if gen.tilesPlayed == 7 {
		bingoBonus = 50
	}
	wordMultiplier := 1

	for idx, rn := range word {
		ml := alphabet.MachineLetter(rn)
		bonusSq := gen.board.GetBonus(gen.curRowIdx, col+idx)
		letterMultiplier := 1
		thisWordMultiplier := 1
		freshTile := false
		if ml == alphabet.PlayedThroughMarker {
			ml = gen.board.GetLetter(gen.curRowIdx, col+idx)
		} else {
			freshTile = true
			// Only count bonus if we are putting a fresh tile on it.
			switch bonusSq {
			case board.Bonus3WS:
				wordMultiplier *= 3
				thisWordMultiplier = 3
			case board.Bonus2WS:
				wordMultiplier *= 2
				thisWordMultiplier = 2
			case board.Bonus2LS:
				letterMultiplier = 2
			case board.Bonus3LS:
				letterMultiplier = 3
			}
			// else all the multipliers are 1.
		}
		cs := gen.board.GetCrossScore(gen.curRowIdx, col+idx, dir)

		if ml >= alphabet.BlankOffset {
			// letter score is 0
			ls = 0
		} else {
			ls = gen.bag.Score(ml)
		}

		mainWordScore += ls * letterMultiplier
		// We only add cross scores if the cross set of this square is non-trivial
		// (i.e. we have to be making an across word). Note that it's not enough
		// to check that the cross-score is 0 because we could have a blank.
		if freshTile && gen.board.GetCrossSet(gen.curRowIdx,
			col+idx, dir) != board.TrivialCrossSet {

			crossScores += ls*letterMultiplier*thisWordMultiplier + cs*thisWordMultiplier
		}
	}
	return mainWordScore*wordMultiplier + crossScores + bingoBonus
}

// Plays returns the generator's generated plays.
func (gen *GordonGenerator) Plays() []*move.Move {
	return gen.plays
}

func calculateEquity(play *move.Move, lm LeaveMap, boardEmpty bool) float64 {
	// If the leave is not present in the map, this is just 0 + score.
	leave := play.Leave()
	score := play.Score()

	adjustment := 0.0
	if boardEmpty {
		adjustment += placementAdjustment(play)
	}

	// also need a pre-endgame adjustment that biases towards leaving
	// one in the bag, etc.
	return lm[MachineWordString(leave)] + float64(score) + adjustment
}

func placementAdjustment(play *move.Move) float64 {
	// Very simply just checks how many vowels are overlapping bonus squares.
	// This only gets considered when the board is empty.
	row, col, vertical := play.CoordsAndVertical()
	var start, end int
	start = col
	if vertical {
		start = row
	}
	end = start + play.TilesPlayed()

	j := start
	penalty := 0.0
	vPenalty := -0.7 // VERY ROUGH approximation from Maven paper.
	for j < end {
		if play.Tiles()[j-start].IsVowel(play.Alphabet()) {
			switch j {
			case 2, 6, 8, 12:
				// row/col below/above have a 2LS. note this only works
				// for specific board configurations
				penalty += vPenalty
			default:

			}
		}
		j++
	}
	return penalty
}
