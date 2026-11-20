//ff:func feature=parse type=helper control=sequence
//ff:what TestNearOf — a candidate as the rules of a sentence read it: a stem with a case particle gives the stem, its name and the particle's name, is no predicate and may stand after an adnominal; a stem with a closing (맺음) is a predicate, and so is one with a modifying (꾸밈) or naming (이름) closing; a new word alone gives its letters with no name and no last piece, ends in a nominal and may stand after an adnominal; a run of signs alone with no name (○○) is a run of signs, not with a marker after it nor with a letter in it; a word that begins with no stem (attached to the word before) gives no stem; a branch with no word, or a word with no part, gives nothing
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNearOf(t *testing.T) {
	word := func(parts ...Part) Branch { return Branch{Words: []Eojeol{{Parts: parts}}} }
	stem := Part{Text: "학생", Name: Name{Kind: "열쇠", Form: "체언"}, Stem: true}
	verb := Part{Text: "되", Name: Name{Kind: "열쇠", Form: "동사"}, Stem: true}
	comp := Part{Text: "이", Name: Name{Kind: "보격", Form: "이"}}
	if got := nearOf(word(stem, comp)); got != (rule.Near{Stem: "학생", Kind: "열쇠:체언", Last: "보격:이", Marks: "보격:이", Modifiable: true}) {
		t.Fatal(got)
	}
	past := Part{Text: "었", Name: Name{Kind: "선어말", Form: "었"}}
	end := Part{Text: "다", Name: Name{Kind: "맺음", Form: "다"}}
	if got := nearOf(word(verb, past, end)); got != (rule.Near{Stem: "되", Kind: "열쇠:동사", Last: "맺음:다", Pred: true, Marks: "선어말:었+맺음:다"}) {
		t.Fatal(got)
	}
	adn := Part{Text: "ㄴ", Name: Name{Kind: "꾸밈", Form: "ㄴ"}}
	naming := Part{Text: "ㅁ", Name: Name{Kind: "이름", Form: "ㅁ"}}
	obj := Part{Text: "을", Name: Name{Kind: "목적격", Form: "을"}}
	if !nearOf(word(verb, adn)).Pred || !nearOf(word(verb, naming, obj)).Pred || nearOf(word(verb, past)).Pred {
		t.Fatal("꾸밈 and 이름 make a predicate; a pre-final ending alone does not")
	}
	if got := nearOf(word(Part{Text: "학생이", Stem: true})); got != (rule.Near{Stem: "학생이", Nominal: true, Modifiable: true}) {
		t.Fatal("a new word alone", got)
	}
	if got := nearOf(word(Part{Text: "로", Name: Name{Kind: "부사격", Form: "로"}})); got != (rule.Near{Last: "부사격:로", Marks: "부사격:로"}) {
		t.Fatal("no stem first", got)
	}
	if got := nearOf(word(Part{Text: "○○", Stem: true})); got != (rule.Near{Stem: "○○", Nominal: true, Modifiable: true, Signs: true}) {
		t.Fatal("a run of signs alone with no name", got)
	}
	if nearOf(word(Part{Text: "○○", Stem: true}, comp)).Signs || nearOf(word(Part{Text: "○가", Stem: true})).Signs {
		t.Fatal("a run of signs with a marker after it, or with a letter in it")
	}
	if nearOf(Branch{}) != (rule.Near{}) || nearOf(word()) != (rule.Near{}) {
		t.Fatal("nothing to read")
	}
}
