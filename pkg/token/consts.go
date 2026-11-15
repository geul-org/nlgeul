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

// judgedMax is how many verdicts an index remembers (judged) before it
// drops them all and starts again. It bounds memory only and changes no
// verdict. With no bound the number kept growing (270 to 390 thousand
// after ten thousand sentences). Measured on ten thousand sentences of a
// novel and of statutes, 100,000 read as fast as 250,000, 500,000 and
// 2,000,000 (within the 3% two runs differ by) with the least memory (a
// peak of 145MB against 210 to 260MB); 20,000 read 5 to 10% slower.
const judgedMax = 100000

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

// bonusSentence is what a rule of a sentence adds to the preference of a
// candidate its claim stands for. The candidates such a rule tells apart
// are mostly ones of the same lambda (the same letters under two names),
// which any value above 0 puts in order; the value also carries a
// candidate over one of another cut that was just above it. It is chosen
// among 0.01, 0.1, 0.5, 1 and 2 tried with the first rule of a sentence on
// the answer sheet: 0.01 and 0.1 gave 29 and 23 fewer right words than
// 0.5, and 1 and 2 gave 5 and 8 more.
const bonusSentence = 0.5

// bonusOrder is what a rule of a sentence adds when it only sets an
// order among names of the same letters and reads nothing of the
// sentence (에서 as the adverbial particle before the nominative): enough
// to put two candidates of the same lambda in order, and too little to
// carry a candidate over one of another cut. With bonusSentence in its
// place the answer sheet had two words fewer right (a noun that ends in
// 서 was cut before the 서).
const bonusOrder = 0.01

// adverbialFirst are the names of 에서 and 서 as adverbial particles,
// which a rule of a sentence puts before the same letters read as the
// nominative (adverbialFirstRule).
var adverbialFirst = []string{"부사격:에서", "부사격:서"}

// complementNames are the names of 이 and 가 as complement particles,
// subjectNames those names and the nominative ones (any 이/가), and
// complementStems the stems of the predicates a complement stands
// before (complementRule): 되다 and 아니다.
var (
	complementNames = []string{"보격:이", "보격:가"}
	subjectNames    = []string{"주격:이", "주격:가", "보격:이", "보격:가"}
	complementStems = []string{"되", "아니"}
)

// quoteNames are the names of 라고 and 이라고 as the particle of direct
// quotation, which a rule of a sentence prefers right after a closing
// quotation mark (quotedRule).
var quoteNames = []string{"부사격:라고", "부사격:이라고"}

// predKinds are the kinds of a closing's name: a word that carries one
// is a predicate (a verb, an adjective or the copula with its ending),
// which is where the rules of a sentence stop waiting.
var predKinds = []string{"맺음", "꾸밈", "이름"}

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

// newStemMax is the most tokens (tokens, for preference) the stem of a
// new word may have: a stem no table lists is a candidate only up to it
// (the stem alone, not the particles and closings after it), unless it
// is the whole of its unit. A word of a table has no such limit.
const newStemMax = 10

// windowSize is the most tokens a stretch of a long unit is read ahead
// before its first words are sent up (the window): the places the net has
// worked out, counted from the last cut, must hold this many tokens before
// a cut is tried. It is a second constant with the same value as
// newStemMax, because it means another thing.
const windowSize = 10

// ontoBorn is the number of birth the words of the letters after a sign
// of a pair begin with: the words carried over from the whole unit have
// theirs, which are below it, so that of the same lambda such a word is
// before the ones the letters make themselves, as when the letters
// were read after the whole unit.
const ontoBorn = 1 << 32

// bonusOne is what the rules of 한 standing alone add (determinerOneRule):
// the word table's counts put 하다 with the modifying ㄴ 1.03 of
// preference above the determiner 한 of the same letters, more than
// bonusSentence can carry. Measured on the answer sheet, 1.1, 1.5, 2 and
// 3 gave the same words right (the rule only turns the order among the
// candidates of a 한 written apart); 1.5 is chosen to stay over that gap
// if the counts move a little.
const bonusOne = 1.5

// bonusLinkBeforeAux is what the rule "the word right before an
// auxiliary verb is the 어/아 form of its main verb" (linkBeforeAuxRule)
// adds. The 어/아 form it puts first is mostly a stem of one syllable
// (남+아, 서+(어)) against a noun of the same letters (남아, 서), whose
// lambda is about 2 to 2.5 above it, so bonusSentence does not carry it.
const bonusLinkBeforeAux = 3.0

// bonusMalLinked is what the rule "말다 after -지 or -고 is the auxiliary
// verb" (malLinkedRule) adds. The forms of 말다 that drop the ㄹ are the
// letters of other words (마라, 마시고, 마소서) whose lambda is up to
// about 3.2 above them. Measured on the answer sheet, the words put
// right (none put wrong at any of them) were 2 at 0.5, 4 at 1.5, 5 at 3,
// 6 at 3.5 and 6 at 6; the least value of the most is chosen.
const bonusMalLinked = 3.5

// bonusAuxLinked is what the rule "an auxiliary verb stands after a
// linking form" (auxLinkedRule) adds; it had bonusSentence. Where the
// auxiliary's letters are also another word (줄 the bound noun or 줄다,
// 주기 the noun, 보아 the noun), that word's lambda is up to about 1.5
// above the auxiliary's, which 0.5 does not carry. Measured on the answer
// sheet against 0.5, the words put right (none put wrong at any of them)
// were 6 at 1, 8 at 1.5, 2, 3, 4 and 6; the least value of the most is
// chosen.
const bonusAuxLinked = 1.5

// tableNoun is the name of a noun of the word table (열쇠:체언), and
// nounNames the kinds of the names of the closed sets' nouns (의존명사:수,
// 대명사:나): the rule of a sentence gluedShortRule speaks of such a noun
// of one syllable, under any of the names its letters are read with.
var (
	tableNoun = keyKind + ":" + kindLabel[kindNominal]
	nounNames = []string{"의존명사", "대명사"}
)

// bonusGluedShort is what the rule of a sentence gluedShortRule adds to a
// noun of one syllable with its particle written onto a nominal: what the
// rebuttal "one syllable long" alone takes from lambda (its verdict is
// 2/(1+forceShort)-1, so the value is about 3.33), so that the
// preferences of the noun with its particle and of the longer word of the
// same letters decide. Measured on the answer sheet against no rule, in
// words right more: 2 gave 28, 2.7 gave 37, this value 56 (65 fixed, 9
// broken), 4 gave 58 (75 fixed, 17 broken), 5 gave 51 and 8 gave 19
// fewer. It is kept for its reason and for half the words broken at 4.
var bonusGluedShort = -Lambda * math.Log10(2/(1+forceShort)-1)

// copulaAdnForms are the letters of the copula 이 with a modifying
// closing fused into it (이+ㄴ, 이+ㄹ), and bonusUnmodified what the rule of
// a sentence unmodifiedRule adds to a noun of the word table that ends in
// them where nothing an adnominal modifies follows: as much as the
// rebuttal "rarely seen" takes at most from lambda (a row seen once: its
// verdict is 2/(1+forceFew)-1, so the value is about 3.19). Measured on
// the answer sheet after gluedShortRule, in words right more: 1 gave 10,
// 2 gave 13, this value and 5 gave 14, none broken.
var (
	copulaAdnForms  = []string{"인", "일"}
	bonusUnmodified = -Lambda * math.Log10(2/(1+forceFew)-1)
)

// topicNames are the names of 이란 and 란 as the topic particle, which a
// rule of a sentence prefers right after a closing sign (shutTopicRule).
var topicNames = []string{"보조사:이란", "보조사:란"}

// bonusShutTopic is what that rule adds. bonusSentence is too little:
// after a closing sign the noun 이란 and 이+란 read as the copula with
// the modifying 란 (이라는) stand above the particle 이란 by about 1.2
// in preference, because the count table gives 이란 47 against 란's
// 7,069. Tried on the answer sheet: 0.5 and 1 put no word right, 1.2,
// 1.5, 2, 3 and 5 each put eight right and none wrong; 1.5 is taken to
// stand clear of the gap without reaching far past it.
const bonusShutTopic = 1.5

// copulaEndNames are the names of the ending 야, which follows the
// copula 이 only, and of the ending 라고, and sentenceEnds the names of
// the last piece of the word right after a word that ends its sentence:
// none (nothing follows), a full stop or a question mark (copulaEndRule).
var (
	copulaEndNames = []string{"맺음:야", "맺음:라고"}
	sentenceEnds   = []string{"", "마침표", "물음표"}
)

// bonusCopula is what a rule of a sentence adds to a candidate that reads
// the copula 이 with its ending over a particle of the same letters (the
// ending 야 over the vocative 야). Such a candidate pays for the copula's
// piece and comes about two below the particle (2.05 and 2.16 in the
// words looked at), more than bonusSentence carries. It is chosen among
// 0.5, 2.2, 2.5, 3 and 5 tried with copulaEndRule on the answer sheet:
// 0.5 changed nothing, and the others gave the same seven more right
// words and none fewer; 2.5 leaves room over the gap.
const bonusCopula = 2.5

// numberKinds are the names of the stem of a number (and of an ordinal
// number with 제, so that the two readings of the same digits keep their
// order) that a rule of a sentence reads 라고 after as the particle of
// quotation (numberQuoteRule).
var numberKinds = []string{"수", "차례"}

// copulaNames are the names of the copula 이 written out, contrastNames
// those of the linking ending 나 after it, and commaNames those of the
// comma, which a rule of a sentence reads right after them
// (copulaCommaRule).
var (
	copulaNames   = []string{"서술격:이"}
	contrastNames = []string{"맺음:나"}
	commaNames    = []string{"쉼표"}
)
