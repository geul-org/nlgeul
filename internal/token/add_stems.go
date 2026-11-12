//ff:func feature=token type=rule control=sequence
//ff:what addStems — the graph of a new word (a stem no table lists) and the graph of a stem's restored last syllable; the new word's claim is rebutted by "it is in no key table" and excepted when its letters read as markers after a boundary, and with a word table when it ends in ㅆ before markers and when the tables know it only as kinds that take no predicate tail before one; the restored syllable's claim is excepted when it stands right after a letter that is no Hangul, when the stem keeps its ㅡ, and when the 어/아 put back does not fit the stem
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A token sequence that no table lists is a new word: the claim that
// it is a word stands, rebutted at the force forceNew by "it is in no
// key table". Its verdict (1/3) is a low positive one, as
// token-stream-algorithm.md section 12 asks. A new word after a
// boundary inside a unit does not stand when its letters read as
// markers there (에 before 서 in 마당 에+서, 께서 after 선생님); without
// a word table that holds only for a word with no marker after it
// (nothing tells 들 the verb's stem from 들 the suffix), so the
// exception is lifted for a stem with markers. With a word table a new
// word with markers after it does not stand when it ends in the final
// consonant ㅆ (such a stem is 있 or a word made with it, which the
// table lists; any other ㅆ is the slot 았/었/겠), nor when the tables
// know its letters, but only as kinds that take no predicate tail,
// before such a tail (뭐+냐, 거+네): a word that is a predicate's stem
// too is left to its kind's graph.
//
// A piece is fused only into a Hangul stem's last syllable: right after
// a letter that is no Hangul inside the stem the syllable is not taken
// apart (api과 is no api고+아), but for 하 and 되, which make a
// predicate of the word before them, and 이, the copula's stem
// (latinBases). A ㅡ put back under ㅏ or ㅓ (따라 = 따르+아) does not
// stand when the word table gives the stem as 르 or 러 irregular (모르
// goes to 몰라, never 모라), when the vowel does not follow the
// syllable before the ㅡ, or when a stem of one syllable is read with
// 아 (써, never 싸).
func (ix *index) addStems() {
	g := toulmin.NewGraph(newWord)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: newWord, Concludes: "열쇠"})
	g.Counter(rule.PieceRead).With(&rule.Claim{Name: noKey, Concludes: "반박"}).Attacks(claim, forceNew)
	marked := except(g, rule.PieceIs, whyMarked, fact(rule.PieceMarked))
	marked.Attacks(claim)
	if ix.lexed {
		except(g, rule.StemCoda, whyCoda, &rule.Bits{Of: "jong", Mask: rule.JongSS}).Attacks(claim)
		known := except(g, rule.StemKind, whyKnown, &rule.Bits{Of: "kinds", Mask: kindAny})
		known.Attacks(claim)
		except(g, rule.StemKind, whyPredicate, &rule.Bits{Of: "kinds", Mask: kindVerb | kindAdj}).Attacks(known)
		except(g, rule.StemTail, whyNoPredTail, &rule.Bits{Of: "tail", Mask: predTails}).Attacks(known)
		ix.edges += 4
	} else {
		except(g, rule.StemTail, whyUntold, &rule.Bits{Of: "tail", Mask: rule.TailBare}).Attacks(marked)
		ix.edges++
	}
	ix.stem, ix.graphs, ix.edges = g, ix.graphs+1, ix.edges+2
	ix.unknown = verdict(g, rule.Piece{})
	r := toulmin.NewGraph(restoreName)
	claim = r.Rule(rule.PieceRead).With(&rule.Claim{Name: restoreName, Concludes: "줄기"})
	latin := except(r, rule.PieceAfter, whyNoHangul, &rule.Bits{Of: "after", Mask: rule.AfterHangul})
	latin.Attacks(claim)
	except(r, rule.PieceIs, whyStemHead, fact(rule.PieceStem), not).Attacks(latin)
	except(r, rule.Listed, whyMakes, listed("latinBases", rule.ListBase, latinBases)).Attacks(latin)
	except(r, rule.PieceIs, whyKeeps, fact(rule.PieceKeeps)).Attacks(claim)
	except(r, rule.PieceClash, whyClash).Attacks(claim)
	one := except(r, rule.Listed, whyOne, listed("oneSyllable", rule.ListForm, []string{"아"}))
	one.Attacks(claim)
	except(r, rule.PieceIs, whyLonger, fact(rule.PieceStem)).Attacks(one)
	ix.restore, ix.graphs, ix.edges = r, ix.graphs+1, ix.edges+7
}
