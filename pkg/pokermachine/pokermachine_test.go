package pokermachine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerMachine(t *testing.T) {
	tests := map[string]struct {
		firstHand      string
		secondHand     string
		expectedWinner string
	}{
		"AAAQQ vs QQAAA": {firstHand: "AAAQQ", secondHand: "QQAAA", expectedWinner: "Tie"},
		"53QQ2 vs Q53Q2": {firstHand: "53QQ2", secondHand: "Q53Q2", expectedWinner: "Tie"},
		"53888 vs 88385": {firstHand: "53888", secondHand: "88385", expectedWinner: "Tie"},
		"QQAAA vs AAAQQ": {firstHand: "QQAAA", secondHand: "AAAQQ", expectedWinner: "Tie"},
		"Q53Q2 vs 53QQ2": {firstHand: "AAAQQ", secondHand: "QQAAA", expectedWinner: "Tie"},
		"88385 vs 53888": {firstHand: "88385", secondHand: "53888", expectedWinner: "Tie"},
		"AAAQQ vs QQQAA": {firstHand: "AAAQQ", secondHand: "QQQAA", expectedWinner: "Hand 1"},
		"Q53Q4 vs 53QQ2": {firstHand: "Q53Q4", secondHand: "53QQ2", expectedWinner: "Hand 1"},
		"53888 vs 88375": {firstHand: "53888", secondHand: "88375", expectedWinner: "Hand 1"},
		"33337 vs QQAAA": {firstHand: "33337", secondHand: "QQAAA", expectedWinner: "Hand 1"},
		"22333 vs AAA58": {firstHand: "22333", secondHand: "AAA58", expectedWinner: "Hand 1"},
		"33389 vs AAKK4": {firstHand: "33389", secondHand: "AAKK4", expectedWinner: "Hand 1"},
		"44223 vs AA892": {firstHand: "44223", secondHand: "AA892", expectedWinner: "Hand 1"},
		"22456 vs AKQJT": {firstHand: "22456", secondHand: "AKQJT", expectedWinner: "Hand 1"},
		"99977 vs 77799": {firstHand: "99977", secondHand: "77799", expectedWinner: "Hand 1"},
		"99922 vs 88866": {firstHand: "99922", secondHand: "88866", expectedWinner: "Hand 1"},
		"9922A vs 9922K": {firstHand: "9922A", secondHand: "9922K", expectedWinner: "Hand 1"},
		"99975 vs 99965": {firstHand: "99975", secondHand: "99965", expectedWinner: "Hand 1"},
		"99975 vs 99974": {firstHand: "99975", secondHand: "99974", expectedWinner: "Hand 1"},
		"99752 vs 99652": {firstHand: "99752", secondHand: "99652", expectedWinner: "Hand 1"},
		"99752 vs 99742": {firstHand: "99752", secondHand: "99742", expectedWinner: "Hand 1"},
		"99753 vs 99752": {firstHand: "99753", secondHand: "99752", expectedWinner: "Hand 1"},
		"QQQAA vs AAAQQ": {firstHand: "QQQAA", secondHand: "AAAQQ", expectedWinner: "Hand 2"},
		"53QQ2 vs Q53Q4": {firstHand: "53QQ2", secondHand: "Q53Q4", expectedWinner: "Hand 2"},
		"88375 vs 53888": {firstHand: "88375", secondHand: "53888", expectedWinner: "Hand 2"},
		"QQAAA vs 33337": {firstHand: "QQAAA", secondHand: "33337", expectedWinner: "Hand 2"},
		"AAA58 vs 22333": {firstHand: "AAA58", secondHand: "22333", expectedWinner: "Hand 2"},
		"AAKK4 vs 33389": {firstHand: "AAKK4", secondHand: "33389", expectedWinner: "Hand 2"},
		"AA892 vs 44223": {firstHand: "AA892", secondHand: "44223", expectedWinner: "Hand 2"},
		"AKQJT vs 22456": {firstHand: "AKQJT", secondHand: "22456", expectedWinner: "Hand 2"},
		"77799 vs 99977": {firstHand: "77799", secondHand: "99977", expectedWinner: "Hand 2"},
		"88866 vs 99922": {firstHand: "88866", secondHand: "99922", expectedWinner: "Hand 2"},
		"9922K vs 9922A": {firstHand: "9922K", secondHand: "9922A", expectedWinner: "Hand 2"},
		"99965 vs 99975": {firstHand: "99965", secondHand: "99975", expectedWinner: "Hand 2"},
		"99974 vs 99975": {firstHand: "99974", secondHand: "99975", expectedWinner: "Hand 2"},
		"99652 vs 99752": {firstHand: "99652", secondHand: "99752", expectedWinner: "Hand 2"},
		"99742 vs 99752": {firstHand: "99742", secondHand: "99752", expectedWinner: "Hand 2"},
		"99752 vs 99753": {firstHand: "99752", secondHand: "99753", expectedWinner: "Hand 2"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			hand1 := NewPokerHand(test.firstHand)
			hand2 := NewPokerHand(test.secondHand)

			winner, _ := FindWinner(hand1, hand2)
			assert.Equal(t, test.expectedWinner, winner)
		})
	}
}
