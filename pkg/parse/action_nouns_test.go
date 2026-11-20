//ff:func feature=parse type=helper control=sequence
//ff:what TestActionNouns — the nouns of a word table that a verb with 하 is made of, of two syllables or more, sorted: 발병 and 준공 (발병하, 준공하 are verbs), not 장관 (no 하 verb), not 정 (one syllable), not 발병하 itself, and none with no table
package parse

import (
	"slices"
	"testing"
)

func TestActionNouns(t *testing.T) {
	lex := map[string]word{
		"준공": {kinds: kindNominal}, "준공하": {kinds: kindVerb},
		"발병": {kinds: kindNominal}, "발병하": {kinds: kindVerb},
		"장관": {kinds: kindNominal}, "장관하": {kinds: kindNominal},
		"정": {kinds: kindNominal}, "정하": {kinds: kindVerb},
		"먹": {kinds: kindVerb},
	}
	if got := actionNouns(lex); !slices.Equal(got, []string{"발병", "준공"}) {
		t.Fatal(got)
	}
	if got := actionNouns(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
