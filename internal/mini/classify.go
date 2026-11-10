//ff:func feature=mini type=rule control=iteration dimension=1
//ff:what classify — read the sentence, judge each word with g and print it as typed with "+" before its particles, "word:constituent"
package mini

import (
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The punctuation around a word stays where it was typed.
func classify(g *toulmin.Graph, sentence string) ([]string, error) {
	fields := strings.Fields(sentence)
	var out []string
	for i, w := range readSentence(sentence) {
		marked, label, err := judge(g, w)
		if err != nil {
			return nil, err
		}
		out = append(out, strings.Replace(fields[i], w.text, marked, 1)+":"+label)
	}
	return out, nil
}
