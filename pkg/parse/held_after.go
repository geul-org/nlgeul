//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what heldAfter — whether a wait for the predicate is open after a word: a word whose first candidate is a predicate closes the wait that was open, and a word with a candidate that a rule of a sentence cannot judge before the words after it are read (waitsFor) opens one; the candidate is asked as judgeSeat gives it to the rules, with the word's letters and whether it is written onto the word before it
package parse

func heldAfter(held bool, n Node, rules []sentenceRule) bool {
	if nearOf(n.Contest.Branches[0]).Pred {
		held = false
	}
	for _, br := range n.Contest.Branches {
		self := nearIn(n, br)
		self.Text = n.Contest.Text
		held = held || waitsFor(self, rules)
	}
	return held
}
