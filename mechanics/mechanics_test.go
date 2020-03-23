package mechanics_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/domino14/macondo/board"
	"github.com/domino14/macondo/gaddagmaker"
	"github.com/domino14/macondo/gcgio"
	"github.com/domino14/macondo/mechanics"
	"github.com/matryer/is"
)

var LexiconDir = os.Getenv("LEXICON_PATH")

func TestMain(m *testing.M) {
	for _, lex := range []string{"America", "NWL18", "pseudo_twl1979", "CSW19"} {
		gdgPath := filepath.Join(LexiconDir, "gaddag", lex+".gaddag")
		if _, err := os.Stat(gdgPath); os.IsNotExist(err) {
			gaddagmaker.GenerateGaddag(filepath.Join(LexiconDir, lex+".txt"), true, true)
			err = os.Rename("out.gaddag", gdgPath)
			if err != nil {
				panic(err)
			}
		}
	}
	os.Exit(m.Run())
}

func TestPlayToTurn(t *testing.T) {
	is := is.New(t)
	curGameRepr, err := gcgio.ParseGCG("../gcgio/testdata/vs_frentz.gcg")
	is.NoErr(err)
	game := mechanics.StateFromRepr(curGameRepr, "CSW19", 0)
	err = game.PlayGameToTurn(curGameRepr, 21)
	is.NoErr(err)
	expectedBoardConfig := `
   A B C D E F G H I J K L M N O
   ------------------------------
 1|=     '       D       V     =|
 2|  -     E N   A   " Q U O T H|
 3|    -   N O N V I R I L E    |
 4|'     A D O   Y       G     '|
 5|      W O K         Z O      |
 6|  "   A W "       " I     "  |
 7|    '   E   '   '   T   '    |
 8|=     C R A A l E D   '     =|
 9|    '   S   X I '       '    |
10|  "       Y E P   G O R   "  |
11|    J I B E   E U O I        |
12|' S A F E     '       P O N D|
13|    I F     A C E r B E R    |
14|G U L   T U M     "       -  |
15|=   S T E R I L E     '     =|
   ------------------------------
`

	b := board.MakeBoard(board.CrosswordGameBoard)
	b.SetToGame(game.Alphabet(), board.VsWho(expectedBoardConfig))
	b.UpdateAllAnchors()
	b.GenAllCrossSets(game.Gaddag(), game.Bag())
	is.True(b.Equals(game.Board()))
}

func TestSetRandomRack(t *testing.T) {
	is := is.New(t)
	curGameRepr, err := gcgio.ParseGCG("../gcgio/testdata/vs_frentz.gcg")
	is.NoErr(err)
	game := mechanics.StateFromRepr(curGameRepr, "CSW19", 0)
	err = game.PlayGameToTurn(curGameRepr, 21)
	is.NoErr(err)
	// The rack should be a subset of MARLINS + AHNTT
	possibleLetters := map[rune]bool{
		'M': true, 'A': true, 'R': true, 'L': true, 'I': true, 'N': true, 'S': true,
		'H': true, 'T': true,
	}

	uniqueRacks := map[string]bool{}

	for i := 0; i < 1000; i++ {
		game.SetRandomRack(1)

		rack := game.RackFor(1)
		for _, l := range rack.String() {
			is.True(possibleLetters[l])
		}
		uniqueRacks[rack.String()] = true
	}
	is.True(len(uniqueRacks) > 200)
}
