//ff:func feature=parse type=rule control=sequence
//ff:what markerGraph — the judgment graph of one marker candidate, stamped from its table row, its exception list and its rows of the case exception table: the claim "this token sequence is this marker" and the structural exceptions against it, among them, but for a suffix, that its unit is Hangul glossed with as many Han characters (and for a closing written wrong the rebuttal of a misspelt form)
package parse

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Every marker needs a stem before it, and no particle, slot or closing
// stands in a unit of Hangul syllables glossed right after with as many
// Han characters in a round bracket (Glossed: 대곤(大棍), 관치(官治) read
// no 곤, no 치); a suffix (a graph named 접미:<form>) is no native
// particle, closing or slot and may be a Sino-Korean one the gloss
// holds (보장률(保障率): 보장+률), so it is not turned down so. A
// whole-syllable form checks
// how the token before it ends; a particle whether its kind takes the particle after it; a suffix that needs a particle after it that one follows; a
// slot its order; a lone final consonant the syllable it sits under
// (with a word table the word before it is asked instead, where the
// stem is known: lexed); a piece that begins with 어/아/여 is rebutted
// when it does not fit the vowel before it, and a form with 으 read
// after a stem's ㄹ (loose) as one that does not fit the consonant
// before it; a particle read only with a word table (wide: a shortened
// particle, 선, 론, 으론) that no word is written right onto it, since it
// closes its chain and ends its word; an auxiliary particle that sets a word before what is said
// of it (leading) whether it stands right on its stem with a sign that
// ends a sentence closing the unit after it; a closing written only
// onto a Hangul stem (hangulBound) whether a Hangul syllable stands
// before it; a particle whose letters are a closing written wrong's
// (echoes) the rebuttal that it stands right on a stem with no final
// consonant before 요; and a form that ends a
// closed-set word yields to that word. A graph the case exception table
// names (cases) yields to the words of its rows, by an exception of its
// own: its reason tells it from a closed-set word in a trace, and the
// rows it carries (caseRows) tell which row applied. The graph of a closing written
// wrong (wrote) also rebuts its claim as a misspelt form, as the graphs
// of the kinds of stem do, and no exception lifts that. The number of edges is returned.
func markerGraph(e *entry, whole []string, cases []caseException, lexed bool) int {
	g := toulmin.NewGraph(e.name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: e.name, Concludes: "표지"})
	except(g, rule.PieceIs, whyNoStem, fact(rule.PieceStem), not).Attacks(claim)
	edges := 1
	if !strings.HasPrefix(e.name, "접미:") {
		except(g, rule.PieceIs, whyGlossed, fact(rule.PieceGlossed)).Attacks(claim)
		edges++
	}
	if e.lead == 0 && e.after != rule.AfterAny {
		except(g, rule.PieceAfter, whyAfter, &rule.Bits{Of: "after", Mask: e.after}).Attacks(claim)
		edges++
	}
	if e.class == classParticle {
		except(g, rule.ChainLinks, whyChain).Attacks(claim)
		edges++
	}
	if slices.Contains(leading, e.name) {
		except(g, rule.PieceIs, whyClosed, fact(rule.PieceClosed)).Attacks(claim)
		edges++
	}
	if slices.Contains(hangulBound, e.name) {
		except(g, rule.PieceAfter, whyNoHangul, &rule.Bits{Of: "after", Mask: rule.AfterHangul}).Attacks(claim)
		edges++
	}
	if e.inner {
		except(g, rule.PieceIs, whyLast, fact(rule.PieceLast)).Attacks(claim)
		edges++
	}
	if e.class == classSlot {
		except(g, rule.SlotOrder, whySlot).Attacks(claim)
		edges++
	}
	if e.class == classClosing && e.form == "" && e.lead > 0 && !lexed {
		except(g, rule.Listed, whyBase, listed("bases", rule.ListBase, rule.BareBases(e.lead)), not).Attacks(claim)
		edges++
	}
	if e.wide && e.class == classClosing {
		except(g, rule.PieceRight, whyRight, &rule.Bits{Of: "right", Mask: rule.RightParticle}).Attacks(claim)
		edges++
	}
	if e.wide && e.class == classParticle {
		except(g, rule.PieceRight, whyRight, &rule.Bits{Of: "right", Mask: rule.RightOther}).Attacks(claim)
		edges++
	}
	if e.head&headWhole != 0 {
		g.Counter(rule.PieceClash).Attacks(claim, forceClash)
		edges++
	}
	if slices.Contains(echoes, e.name) {
		g.Counter(rule.PieceIs).With(fact(rule.PieceYoOpen)).Attacks(claim, forceEcho)
		edges++
	}
	if e.loose {
		g.Counter(rule.PieceAfter).With(&rule.Bits{Of: "after", Mask: rule.AfterHard}).Attacks(claim, forceClash)
		edges++
	}
	if len(whole) > 0 {
		except(g, rule.Listed, whyWhole, listed("wholeWords", rule.ListWord, whole)).Attacks(claim)
		edges++
	}
	if len(cases) > 0 {
		except(g, rule.Listed, whyCase, listed("caseExceptions", rule.ListWord, caseWords(cases)), &caseRows{Rows: cases}).Attacks(claim)
		edges++
	}
	if e.wrote != "" {
		edges += misspeltRebuttal(g, claim)
	}
	e.g = g
	return edges
}
