//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what facts — the word as the rule package's facts, which the rules read from the context
package parse

import "github.com/geul-org/koparse/pkg/rule"

func facts(w *word) *rule.Word {
	c := w.chain
	out := &rule.Word{
		Text:      w.text,
		Chain:     rule.Chain{Size: c.size, Parts: c.parts, Last: c.last, Count: c.count, Kase: c.kase, Marked: c.marked},
		Longest:   w.longest,
		Last:      w.last,
		Bound:     w.bound,
		NextBound: w.nextBound,
		Sent:      &rule.Sentence{},
	}
	if w.sent != nil {
		out.Sent = &rule.Sentence{Stems: w.sent.stems, Nominative: w.sent.nominative}
	}
	for _, t := range w.tails {
		out.Tails = append(out.Tails, tailFacts(t))
	}
	return out
}
