//ff:func feature=token type=reader control=iteration dimension=1
//ff:what readSentence — what the example does with one sentence: each stretch's best way, with at most head candidates a word, printed; where it leans on a new word, the ways after it until one of known words only (againKnown)
package token_test

import "github.com/geul-org/nlgeul/pkg/token"

func readSentence(st *token.Stream, head int) {
	for seg, ok := st.NextHead(head); ok; seg, ok = st.NextHead(head) {
		printWay(seg)
		if !wayKnown(seg) {
			againKnown(st, seg.Region, head)
		}
	}
}
