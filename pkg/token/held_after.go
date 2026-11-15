//ff:func feature=token type=helper control=iteration dimension=1
//ff:what heldAfter — whether a wait for the predicate is open after a word: a word whose first candidate is a predicate closes the wait that was open, and a word with a candidate that a rule of a sentence cannot judge before the words after it are read (waitsFor) opens one
package token

func heldAfter(held bool, n Node, rules []sentenceRule) bool {
	if nearOf(n.Contest.Branches[0]).Pred {
		held = false
	}
	for _, br := range n.Contest.Branches {
		held = held || waitsFor(nearOf(br), rules)
	}
	return held
}
