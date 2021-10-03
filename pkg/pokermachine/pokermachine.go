package pokermachine

func FindWinner(firstHand PokerHand, secondHand PokerHand) (string, error) {
	if firstHand.lcs.length > secondHand.lcs.length {
		return "Hand 1", nil
	}

	if firstHand.lcs.length < secondHand.lcs.length {
		return "Hand 2", nil
	}

	if len(firstHand.occurenceMap) < len(secondHand.occurenceMap) {
		return "Hand 1", nil
	}

	if len(firstHand.occurenceMap) > len(secondHand.occurenceMap) {
		return "Hand 2", nil
	}

	if firstHand.lcs.count > secondHand.lcs.count {
		return "Hand 1", nil
	}

	if firstHand.lcs.count < secondHand.lcs.count {
		return "Hand 2", nil
	}

	if firstHand.lcs.value > secondHand.lcs.value {
		return "Hand 1", nil
	}

	if firstHand.lcs.value < secondHand.lcs.value {
		return "Hand 2", nil
	}

	if len(firstHand.occurenceMap) > 0 && len(secondHand.occurenceMap) > 0 {
		UpdateNextLcs(&firstHand)
		UpdateNextLcs(&secondHand)
		return FindWinner(firstHand, secondHand)
	}

	return "Tie", nil
}

func UpdateNextLcs(hand *PokerHand) {
	delete(hand.occurenceMap, hand.lcs.card)
	hand.lcs = GenerateLongestCardSequence(hand.occurenceMap)
	for occurence := range hand.occurenceMap {
		if hand.occurenceMap[occurence] <= hand.lcs.count && hand.rules[occurence] > hand.lcs.value {
			hand.lcs.card = occurence
			hand.lcs.value = hand.rules[occurence]
		}
	}
}
