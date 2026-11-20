//ff:func feature=parse type=helper control=sequence
//ff:what firstNode — the first word of the best way of the first stretch of a sentence read with words-sentence.tsv and no rule of a sentence (학생이 gives the node of 학생이[0-3] with its tied 주격 and 보격 이)
package parse

import "testing"

func firstNode(t *testing.T, s string) Node {
	t.Helper()
	return firstNodeOf(t, sentenceRegions(t, s)[0])
}
