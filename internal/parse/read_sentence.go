//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what readSentence — read every word of the sentence, then the facts of the whole: the last word's stems under its longest tails, how many earlier words carry 이/가, and for each word whether the next begins with a dependent noun
package parse

import "strings"

func readSentence(text string) []*word {
	sent := &sentence{}
	fields := strings.Fields(text)
	var words []*word
	for i, f := range fields {
		w := readWord(f, i == len(fields)-1, sent)
		if i > 0 {
			words[i-1].nextBound = w.bound
		}
		words = append(words, w)
		if !w.last && w.chain.kase&caseComp != 0 {
			sent.nominative++
		}
	}
	if len(words) > 0 {
		sent.stems = fullStems(words[len(words)-1])
	}
	return words
}
