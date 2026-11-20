//ff:func feature=parse type=rule control=sequence
//ff:what pluralNounRule — the rule of a sentence "a word the word table gives both as a nominal and as an adverb, with the suffix 들 after it, is the nominal": the candidate that reads such letters as the nominal of the word table (열쇠:체언) with the suffix 들 (pluralNames) among its markers is preferred (bonusOverCount); it waits for nothing
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The suffix -들 adds the sense of more than one to a noun (the
// dictionary: 셀 수 있는 명사 뒤에 붙는다), and the tables read 들 after
// an adverb too, as the particle of the dictionary that follows an
// adverb or an ending (잘들, 어서들). Letters the word table gives both as
// a nominal and as an adverb (모두, 서로, 대신) with 들 after them are
// the nominal with its plural suffix; the table counts most of them more
// often as the adverb, so the adverb with 들 came first. Letters that are
// an adverb only keep 들 after the adverb.
func pluralNounRule(lex map[string]word) sentenceRule {
	g, _ := sentenceClaim(rulePluralNoun, rule.SelfHas, listed("접미 들", rule.ListSelfMarks, pluralNames))
	is := func(n rule.Near) bool {
		return n.Kind == tableNoun && lex[n.Stem].kinds&kindAdverb != 0 && hasMark(n, pluralNames)
	}
	return sentenceRule{name: rulePluralNoun, graph: g, bonus: bonusOverCount, about: is}
}
