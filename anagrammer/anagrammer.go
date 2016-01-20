// Package anagrammer implements an anagrammer based on a GADDAG data
// structure.
// I am basing this (and the movegen in the future) on my ujamaa program
// at github.com/domino14/ujamaa
// However, in order to simplify data structures and (hopefully) speed things
// up, we replace ARC* with just the index of the node that the ARC would
// point to in all the algorithms.
package anagrammer

import (
	"fmt"
	"github.com/domino14/macondo/gaddag"
	"github.com/domino14/macondo/movegen"
	"strings"
	"time"
)

const (
	ModeBuild   = iota
	ModeAnagram = iota
)

type Rack struct {
	rack  []uint8
	count uint8
}

var answerSet map[string]bool

// The main anagram logic.
// This function could be faster if we got rid of the inner for-loop below.
// However we would need a new data structure (with no LetterSets) for this.
// Hopefully this GADDAG will still be fast when doing move generation
// (compared to Quackle).
func anagram(g gaddag.SimpleGaddag, nodeIdx uint32, prefix []byte,
	rack *Rack, mode uint8, alphabet *Alphabet) {
	// Go through all arc children.
	var numArcs, rackPos, i, j byte
	var newPrefix []byte
	var nextNodeIdx, letterSet uint32
	var letter rune
	numArcs = byte(g[nodeIdx] >> gaddag.NumArcsBitLoc)
	for i = byte(1); i <= numArcs; i++ {
		nextNodeIdx, letter = g.ArcToIdxLetter(nodeIdx+uint32(i), alphabet)
		if letter == gaddag.SeparationToken {
			break
		}
		rackPos = letter - 'A'
		if rack.rack[rackPos] == 0 {
			if rack.rack[movegen.BlankPosition] == 0 {
				continue
			} else {
				rackPos = movegen.BlankPosition
			}
		}
		rack.rack[rackPos]--
		rack.count--
		newPrefix = []byte{letter}
		newPrefix = append(newPrefix, prefix...)
		// Now check the LetterSet of this new node and add plays.
		// This is kind of awkward.
		//letterSet = g[nextNodeIdx] & ((1 << gaddag.NumArcsBitLoc) - 1)
		letterSet = g.GetLetterSet(nextNodeIdx)
		// This stupid loop is slow.
		if letterSet != 0 {
			for j = uint8(0); j < 26; j++ {
				if letterSet&(1<<j) != 0 {
					if mode == ModeAnagram &&
						(rack.rack[j] == 1 || rack.rack[movegen.BlankPosition] == 1) &&
						rack.count == 1 {
						addPlay(append([]byte{j + 'A'}, newPrefix...))
					} else if mode == ModeBuild &&
						(rack.rack[j] != 0 || rack.rack[movegen.BlankPosition] != 0) {
						addPlay(append([]byte{j + 'A'}, newPrefix...))
					}
				}
			}
		}

		anagram(g, nextNodeIdx, newPrefix, rack, mode, alphabet)
		// Restore letter to rack.
		rack.rack[rackPos]++
		rack.count++
	}

}

// turnStringIntoRack Turns a given rack into a uint8 slice of 27 integers,
// one for each letter of the alphabet (blank is the 27th).
func turnStringIntoRack(str string) Rack {
	fmt.Println("Turning", str, "into rack")
	rack := make([]uint8, movegen.NumTotalLetters)
	str = strings.ToUpper(str)
	ct := 0
	for _, c := range str {
		if c == '_' || c == '?' {
			rack[movegen.BlankPosition]++
		} else {
			rack[c-'A']++
		}
		ct++
	}
	r := Rack{rack, uint8(ct)}
	fmt.Println(r)
	return r
}

func rackToString(rack *Rack) string {
	str := ""
	for idx, val := range rack.rack {
		for j := uint8(0); j < val; j++ {
			if idx < 26 {
				str += string(idx + 'A')
			} else {
				str += "?"
			}
		}
	}
	return str
}

func addPlay(word []byte) {
	answerSet[string(word)] = true
}

// Anagram anagrams or builds the passed in string.
func Anagram(gaddagData []uint32, str string, mode uint8) {
	answerSet = make(map[string]bool)
	rack := turnStringIntoRack(str)
	t0 := time.Now()
	alphabet := gaddag.SimpleGaddag(gaddagData).GetAlphabet()
	anagram(gaddagData, 0, []byte(nil), &rack, mode, alphabet)
	t1 := time.Now()
	//fmt.Println(answerSet)
	fmt.Println(len(answerSet), "answers")
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

}
