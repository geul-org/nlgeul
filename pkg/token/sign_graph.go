//ff:func feature=token type=rule control=selection
//ff:what signGraph — the judgment graph of one candidate of a sign at a unit's edge or between letters of a unit: the claim "this sign is that", and the exceptions against it by where the sign stands, what stands on its two sides and how many stand in a row (and for 물결표, 빗금 and 쌍점 an exception against its exception); the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// 쉼표 is one comma. 물음표 and 느낌표 are any number of the sign. 쌍점
// and 쌍반점 are one sign at a word's end or alone, 쌍점 also between
// letters that are not both Latin letters or digits (EPI:입자, 復辟:퇴위,
// 줄:1 — the orthography writes the word after it apart; a:b, one with a
// sign next to it, C:\, and one at a unit's head stay as they were); 물결표 any number
// there, or between two Hangul syllables (the exception that it stands
// neither at a word's end nor alone is itself defeated there). 빗금 is
// a slash between two Hangul syllables, or at a unit's edge (after the
// letters that are read, right before them, or alone: the exception
// that its sides are not Hangul is itself defeated there). 줄임표 (…) and 줄표 (― —) are
// any number of the sign. 샵 and 골뱅이 are one sign with letters right
// after it or alone; 태그 and 명칭 one sign with letters right after it.
// 제목 is the sign at the very start of an input line with a space right
// after it. 가운뎃점 is one sign; 더하기 and 연결 are one sign with no
// other sign right next to it.
func signGraph(name string) (*toulmin.Graph, int) {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "기호"})
	is := func(mask int) *rule.Bits { return &rule.Bits{Of: "mark", Mask: mask} }
	repeated := func() { except(g, rule.MarkIs, whySameSigns, is(rule.MarkRepeated)).Attacks(claim) }
	switch name {
	case "쉼표":
		repeated()
		return g, 1
	case "쌍점":
		repeated()
		e := except(g, rule.MarkIs, whyNoEnd, is(rule.MarkEnds|rule.MarkAlone), not)
		e.Attacks(claim)
		apart := except(g, rule.MarkBoth, whyNotAlnumSides, &rule.Bits{Of: "sides", Mask: rule.SideAlnum}, not)
		apart.Attacks(e)
		except(g, rule.MarkIs, whyBySign, is(rule.MarkBySign)).Attacks(apart)
		except(g, rule.MarkIs, whyAtEdge, is(rule.MarkEnds|rule.MarkLeads|rule.MarkAlone)).Attacks(apart)
		return g, 5
	case "쌍반점":
		repeated()
		except(g, rule.MarkIs, whyNoEnd, is(rule.MarkEnds|rule.MarkAlone), not).Attacks(claim)
		return g, 2
	case "물결표":
		e := except(g, rule.MarkIs, whyNoEnd, is(rule.MarkEnds|rule.MarkAlone), not)
		e.Attacks(claim)
		except(g, rule.MarkBoth, whyHangulSides, &rule.Bits{Of: "sides", Mask: rule.SideHangul}).Attacks(e)
		return g, 2
	case "빗금":
		e := except(g, rule.MarkBoth, whyNotHangulSides, &rule.Bits{Of: "sides", Mask: rule.SideHangul}, not)
		e.Attacks(claim)
		except(g, rule.MarkIs, whyAtEdge, is(rule.MarkEnds|rule.MarkLeads|rule.MarkAlone)).Attacks(e)
		return g, 2
	case "줄임표", "줄표":
		return g, 0
	case "샵", "골뱅이":
		repeated()
		except(g, rule.MarkIs, whyNoLead, is(rule.MarkLeads|rule.MarkAlone), not).Attacks(claim)
		return g, 2
	case "태그", "명칭":
		repeated()
		except(g, rule.MarkIs, whyNoLetters, is(rule.MarkLeads), not).Attacks(claim)
		return g, 2
	case "가운뎃점":
		repeated()
		return g, 1
	case "더하기", "연결":
		repeated()
		except(g, rule.MarkIs, whyBySign, is(rule.MarkBySign)).Attacks(claim)
		return g, 2
	case headingName:
		except(g, rule.MarkIs, whyNoLineHead, is(rule.MarkHeads), not).Attacks(claim)
		except(g, rule.MarkIs, whyNoSpace, is(rule.MarkSpaced), not).Attacks(claim)
		return g, 2
	}
	return g, 0
}
