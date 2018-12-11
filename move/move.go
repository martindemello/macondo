package move

import (
	"fmt"
	"strconv"

	"github.com/domino14/macondo/alphabet"
)

// MoveType is a type of move; a play, an exchange, pass, etc.
type MoveType uint8

const (
	MoveTypePlay MoveType = iota
	MoveTypeExchange
	MoveTypePass
	MoveTypePhonyTilesReturned

	MoveTypeEndgameTiles
	MoveTypeLostTileScore
)

// Move is a move. It can have a score, position, equity, etc. It doesn't
// have to be a scoring move.
type Move struct {
	action      MoveType
	score       int
	equity      int
	desc        string
	coords      string
	tiles       alphabet.MachineWord
	leave       alphabet.MachineWord
	rowStart    uint8
	colStart    uint8
	vertical    bool
	bingo       bool
	tilesPlayed uint8
	alph        *alphabet.Alphabet
}

// ByScore is a helper type to help sort moves by score.
type ByScore []*Move

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].score > a[j].score }

// ByEquity is a helper type to help sort moves by equity.
type ByEquity []*Move

func (a ByEquity) Len() int           { return len(a) }
func (a ByEquity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEquity) Less(i, j int) bool { return a[i].equity > a[j].equity }

func (m Move) String() string {
	switch m.action {
	case MoveTypePlay:
		return fmt.Sprintf(
			"<action: play word: %v %v (%v) score: %v tp: %v leave: %v>",
			m.coords, m.tiles, m.tiles.UserVisible(m.alph), m.score,
			m.tilesPlayed, m.leave.UserVisible(m.alph))
	case MoveTypePass:
		return fmt.Sprint("<action: pass>")
	}
	return fmt.Sprint("<Unhandled move>")

}

func (m Move) Action() MoveType {
	return m.action
}

func (m Move) TilesPlayed() uint8 {
	return m.tilesPlayed
}

func NewScoringMove(score int, tiles alphabet.MachineWord,
	leave alphabet.MachineWord, vertical bool, tilesPlayed uint8,
	alph *alphabet.Alphabet, rowStart uint8, colStart uint8, coords string) *Move {

	move := &Move{
		action: MoveTypePlay, score: score, tiles: tiles, leave: leave, vertical: vertical,
		bingo: tilesPlayed == 7, tilesPlayed: tilesPlayed, alph: alph,
		rowStart: rowStart, colStart: colStart, coords: coords,
	}
	return move
}

func (m Move) Score() int {
	return m.score
}

func (m Move) Leave() alphabet.MachineWord {
	return m.leave
}

func (m Move) Tiles() alphabet.MachineWord {
	return m.tiles
}

func (m Move) UniqueSingleTileKey() int {
	// Find the tile.
	var idx int
	var ml alphabet.MachineLetter
	for idx, ml = range m.tiles {
		if ml != alphabet.PlayedThroughMarker {
			break
		}
	}

	var row, col uint8
	row = m.rowStart
	col = m.colStart
	// We want to get the coordinate of the tile that is on the board itself.
	if m.vertical {
		row, col = col, row
		row += uint8(idx)
	} else {
		col += uint8(idx)
	}
	// A unique, fast to compute key for this play.
	return int(row) + alphabet.MaxAlphabetSize*int(col) +
		alphabet.MaxAlphabetSize*alphabet.MaxAlphabetSize*int(ml)
}

func (m *Move) CoordsAndVertical() (uint8, uint8, bool) {
	return m.rowStart, m.colStart, m.vertical
}

func ToBoardGameCoords(row uint8, col uint8, vertical bool) string {
	colCoords := string(rune('A' + col))
	rowCoords := strconv.Itoa(int(row + 1))
	var coords string
	if vertical {
		coords = colCoords + rowCoords
	} else {
		coords = rowCoords + colCoords
	}
	return coords
}

func NewPassMove() *Move {
	return &Move{
		action: MoveTypePass,
	}
}