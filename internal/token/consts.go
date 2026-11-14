package token

import (
	"math"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Classes of index entry.
const (
	classParticle = iota // 조사
	classSlot            // 선어말 칸
	classClosing         // 맺음
	classKey             // 닫힌 집합의 낱말(열쇠)
)

// keyVowel is the index key of 어/아 fused into a stem's last vowel; a
// final consonant put under the stem's last syllable is keyed by its
// letter (rule.LeadText). The final-consonant indexes, the lead of a
// fused 어/아 and the attachment conditions are the rule package's
// (rule.Jong…, rule.LeadVowel, rule.After…) and have no second name
// here.
const keyVowel = 'ㅓ'

// syllables is the number of Hangul syllables: with no key table, a
// stem syllable may be any of them, so each costs log10(1/syllables)
// of preference.
const syllables = 11172

// newWord and noKey name the graph of a stem that is not in any table.
// keyKind is the kind of a stem's candidate that a word table gives.
const (
	newWord = "새 낱말"
	noKey   = "열쇠 표에 없다"
	keyKind = "열쇠"
)

// typoSign and typoName mark a part written in a wrong form: the sign
// that joins the mark to the part's candidates, and the mark's name
// (corrected). 열쇠:동사 written 맞기 is 열쇠:동사,오타정정:맞기.
const (
	typoSign = ","
	typoName = "오타정정"
)

// restoreName names the graph of a stem's last syllable as it is put
// back when a piece is taken out of it (마신 → 마시, 따라 → 따르).
const restoreName = "줄기:되돌린 끝 음절"

// Kinds of stem, by what may follow them (kindTails). The parts of
// speech of a word table are folded into these.
const (
	kindNominal    = 1 << iota // 체언: 명사, 대명사, 수사, 식별자
	kindVerb                   // 동사 줄기 (보조 용언 포함)
	kindAdj                    // 형용사 줄기
	kindAdverb                 // 부사
	kindDeterminer             // 관형사
	kindAlone                  // 홀로 서는 말: 감탄사와 그 밖
)

// kindAny is every kind: a stem with one of them is a word some table
// knows.
const kindAny = kindNominal | kindVerb | kindAdj | kindAdverb | kindDeterminer | kindAlone

// Layers of a word table's use: whether a stem is in the table; whether
// its kind fits what follows it; and the cost of structure.
const (
	layerNone = iota
	layerKnown
	layerKind
	layerCost
)

// forceNew is the force of the rebuttal "it is in no key table"
// against a new word: with it a new word's verdict is 2/(1+0.5)-1 =
// 1/3, a low positive one. It is a placeholder, not a measured value.
const forceNew = 0.5

// How strong the evidence is that a stem found in the word table is a
// word (the verdict of its kind's graph). A stem observed seenEnough
// times or more is not rebutted as rarely seen; forceFew and forceShort
// are the forces of the rebuttals "rarely seen" and "a key of one
// syllable", and shortKey the length that counts as short. seenEnough,
// forceShort and shortKey are placeholders, not measured values.
// forceFew is chosen among a few values tried on a thousand sentences
// so that a row seen once (or with no count), even of one syllable,
// still stands above a new word: with the two rebuttals in full the
// verdict is 2/(1+0.24+0.25)-1, just over a new word's 1/3.
const (
	seenEnough = 10000
	forceFew   = 0.24
	forceShort = 0.25
	shortKey   = 1
)

// forceWhole is the force of the rebuttal "the number and the unit
// together are a word of the table" against a unit read after a lone
// Hangul numeral (이건, 세대). It is chosen among a few values tried on
// a thousand sentences: below about 0.94 the unit candidate still comes
// first, and at 1 it is not kept at all.
const forceWhole = 0.98

// forceClash is the force of the rebuttal "어/아 does not fit the vowel
// before it" (먹아, 잡어), and of "a form with 으 does not fit the ㄹ
// before it" (줄은): the candidate is kept and comes later. It is chosen
// among a few values tried on a thousand sentences for the first of the
// two; nothing records a measurement for the second, which borrows it.
const forceClash = 0.5

// forceTypo is the force of the rebuttal "the letters are a misspelling
// of the word" against a stem that the word table gives as a wrong way
// to write another word (misspeltRebuttal). With only that rebuttal the
// verdict is 2/(1+0.28)-1, about 0.56, below the verdict of the word
// written right, so that the right candidate of the same letters comes
// first. forceClash (먹아, 줄은) is not used: those forms
// have another candidate of the same letters that fits the norm, and the
// candidate is sent behind it; a misspelt stem has no right candidate of
// those letters to give way to. It is a placeholder, not a measured
// value.
const forceTypo = 0.28

// The costs of structure, in the units of preference (log10). They are
// placeholders, not measured values: nothing in the tables says how
// much a boundary, an unseen piece or a piece is worth. Each is paid
// once: for every word after a unit's first, for every piece that
// leaves no trace in the letters (어/아 fused into ㅏ ㅓ ㅐ ㅔ with
// nothing after it), and for the stem of a word. The markers of the
// index (particles, slots, closings, units, signs) pay no such cost:
// a word read into its pieces is not dearer than the same word whole.
const (
	costBoundary = 1.0
	costUnseen   = 1.0
	costPiece    = 1.0
)

// not turns a yes/no rule round (rule.Not); one value serves every
// node.
var not = &rule.Not{}

// noBorn is the number of birth of a branch that was not made by a scan
// (a sign of an edge, a unit of measure): of the same lambda as another
// it comes after it.
const noBorn = math.MaxInt
