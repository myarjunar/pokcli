package pokermachine

type LongestCardSequence struct {
	card   string
	count  int
	value  int
	length int
}

type PokerHand struct {
	lcs          LongestCardSequence
	occurenceMap map[string]int
	rules        map[string]int
}

var rules = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func NewPokerHand(cards string) PokerHand {
	occurenceMap := map[string]int{}

	cardList := []rune(cards)
	for idx := range cardList {
		card := string(cardList[idx])
		occurenceMap[card] += 1
	}

	lcs := GenerateLongestCardSequence(occurenceMap)

	return PokerHand{
		lcs:          lcs,
		occurenceMap: occurenceMap,
		rules:        rules,
	}
}

func GenerateLongestCardSequence(occurenceMap map[string]int) LongestCardSequence {
	var lcs LongestCardSequence
	lcs.length = 0
	for occurence := range occurenceMap {
		if occurenceMap[occurence] > lcs.length {
			lcs.length = occurenceMap[occurence]
			lcs.card = occurence
		} else if occurenceMap[occurence] == lcs.length {
			lcs.length = occurenceMap[occurence]
			if rules[occurence] > rules[lcs.card] {
				lcs.card = occurence
			}
		}
	}

	lcs.count = 0
	for occurence := range occurenceMap {
		if occurenceMap[occurence] == lcs.length {
			lcs.count += 1
		}
	}

	lcs.value = rules[lcs.card]

	return lcs
}
