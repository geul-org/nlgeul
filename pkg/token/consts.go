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

// honorCount and euCount are the observed counts of the honorific 시
// (honorName) and of the link vowel 으 (euName) in the tables; nothing in
// this repository records where they were taken from.
const (
	honorCount = 19566
	euCount    = 225409
)

// afterOthers is how many pieces besides the unit words are counted
// among the entries a piece after a number is taken as one of: 여 and
// 째. The three unitSuffixes are not counted; nothing records why (the
// count was set before they were added and left as it was).
const afterOthers = 2

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
// word (the verdict of its kind's graph). A stem observed seenShare of
// the word table's counts or more (that share of the sum of every line's
// count: index.seenEnough) is not rebutted as rarely seen, so that what
// counts as seen enough follows the size of the table; forceFew and
// forceShort are the forces of the rebuttals "rarely seen" and "a key of
// one syllable", and shortKey the length that counts as short.
// forceShort and shortKey are placeholders, not measured values.
// seenShare took the place of a count, 10000, a placeholder set when the
// table's counts summed to about 175 million; with the table counted
// again (about 1,341 million) the count was measured on the answer sheet
// against 10000: 30,000 gave 21 more right words, 100,000 31, 300,000
// 60, 700,000 67, 1,000,000 74 (166 righted, 92 turned wrong), 2,000,000
// 68, 10,000,000 55 and 100,000,000 15 more; and the share: 6e-4 68,
// 6.5e-4 70, 7.5e-4 75 (167 righted, 92 turned wrong), 8.5e-4 73, 9e-4 72
// and 1e-3 71 more. 7.5e-4 (about a million times in that table) is the
// least past which no more words are righted.
// forceFew is chosen among a few values tried on a thousand sentences
// so that a row seen once (or with no count), even of one syllable,
// still stands above a new word: with the two rebuttals in full the
// verdict is 2/(1+0.24+0.25)-1, just over a new word's 1/3.
// forceFew was so chosen before the word table was counted again; with
// the table counted again, against this value on the answer sheet 0.18
// gives 192 fewer, 0.25 and 0.3 two fewer, 0.36 two more (26 righted, 24
// turned wrong) and 0.5 five fewer.
const (
	seenShare  = 7.5e-4
	forceFew   = 0.24
	forceShort = 0.25
	shortKey   = 1
)

// forceWhole is the force of the rebuttal "the number and the unit
// together are a word of the table" against a unit read after a lone
// Hangul numeral (이건, 세대). It is chosen among a few values tried on
// a thousand sentences: below about 0.94 the unit candidate still comes
// first, and at 1 it is not kept at all.
// That was measured before the word table was counted again; with the
// table counted again any force from 0.49 to 1 gives the same words
// right on the answer sheet, 0.25 four fewer.
const forceWhole = 0.98

// forceClash is the force of the rebuttal "어/아 does not fit the vowel
// before it" (먹아, 잡어), and of "a form with 으 does not fit the ㄹ
// before it" (줄은): the candidate is kept and comes later. It was 0.5,
// chosen among a few values tried on a thousand sentences for the first
// of the two before the word table was counted again. With the table
// counted again (and seenShare and the steps of the rules of a sentence),
// against 0.5 on the answer sheet 0.125 gives two more right (3 righted,
// 1 turned wrong), 0.25 and 0.375 two more (none turned wrong: one word of
// each rebuttal, both in old spelling, 않어도), and before seenShare 0.625
// gave six fewer and 0 ten fewer; the middle of the flat stretch is taken.
const forceClash = 0.25

// forceEcho is the force of the rebuttal that 에 right on a stem with no
// final consonant before 요 is also the closing 예요 written 에요
// (echoes): the reading with 에 and 요 (a verdict of 1/3) stands behind
// the closing written wrong (2/(1+forceTypo)-1, about 0.56). It was
// forceClash's 0.5 and is kept so when forceClash went to 0.25, at which
// the reading with 에 and 요 (0.6) would stand before the closing written
// wrong; on the answer sheet 0.25 and 0.5 give the same words.
const forceEcho = 0.5

// forceRun is the force of the rebuttal "it is in no key table" against
// a run of Latin letters and digits that no table lists (runName):
// lower than a new Hangul word's, because the change of script at its
// edges is itself a sign of a boundary. It is a placeholder, not a
// measured value.
const forceRun = 0.25

// forceDropAdn is the force of the rebuttal against a dropped copula
// stem before a modifying or a naming closing (dropAdn), chosen among a
// few values tried on a thousand sentences.
// That was measured before the word table was counted again; with the
// table counted again, on the answer sheet 0 rights one word and turns
// one wrong, and 0.25 to 1 give the same words.
const forceDropAdn = 0.5

// What a rule of a sentence adds to the preference of a candidate its
// claim stands for (sentenceRule.bonus) comes in four steps, by how far
// the candidate must be carried. bonusSameLetters puts in order candidates
// of the same lambda, mostly the same letters under two names (에서 as the
// adverbial before the nominative, 이/가 as the complement, 수 as the noun),
// and carries a candidate over one of another cut just above it.
// bonusOverCount carries a candidate over another reading of its letters
// that the word table's counts put above it by up to about 1.6 (the noun 한
// after the modifying 는 under the determiner 한, the particle 이란 after a
// closing sign under the noun 이란, the noun 보고 after a modifier under 보+고,
// the auxiliary 주 after a linking form under the noun 주기). bonusShortTaken
// (below) carries a candidate over another word up to what a rebuttal of
// one syllable takes, and bonusOverWhole (below) a dependent noun read
// apart over the whole word of its letters. The four took the place of a
// value of each rule (thirteen names, ten values from 0.01 to 5). Measured
// on the answer sheet against the values of each rule (with the word table
// counted again and seenShare), each of the first three steps changed with
// the others as they are, in words right more: the first 0.375 gave 15 (26
// righted, 11 turned wrong), 0.5 19 (22, 3) and 0.75 20 (24, 4); the
// second 1.75 17 (20, 3), 2 19, 2.25 18 (25, 7) and 2.5 19 (26, 7), and 3
// for the noun after a modifier alone 11 (27, 16); the third 3 15 (24, 9),
// its value 19, 3.75 21 (26, 5) and about 4.16 three fewer (29, 32). The
// middle of each flat stretch is taken. With two steps (0.5 also for the
// noun after a modifier, the third for all the others) it gave 13 more
// (19, 6); with the steps measured before seenShare (0.5, 1.5, 3.33 and 5)
// 3 more (5, 2), as 한 after the modifying 는 now needs more than 1.5 (1.75
// is enough).
const (
	bonusSameLetters = 0.5
	bonusOverCount   = 2.0
)

// bonusShortTaken is what the rebuttal "one syllable long" alone takes
// from lambda: its verdict is 2/(1+forceShort)-1, so the value is about
// 3.33. A rule adds it where the candidate it prefers stands below
// another word by up to about that: a noun of one syllable the rebuttal
// held down (자+의 written onto a noun under 자의, gluedShortRule, which
// gives back what the rebuttal took), a stem of one syllable before a
// word that shows it is a predicate's (남+아 before 있다, 알+ㄹ before 수,
// 거+는 after a modifying form), the copula's piece (야 at a sentence's
// end, 이나 before a comma) and 말다 with its ㄹ dropped (마라 after -지).
var bonusShortTaken = -Lambda * math.Log10(2/(1+forceShort)-1)

// bonusOverWhole is what the rule "a dependent noun stands right after a
// modifying form" (boundAfterRule) adds: it carries the dependent noun
// read apart with its particle, or with the copula dropped and its
// closing, over the whole word of the same letters (수도 over 수+도, the
// noun 건지 over 거+(이)+ㄴ지 and 거지 over 거+(이)+지 written onto the
// modifying form: 있는건지, 먹는거지), where the dropped copula's piece
// costs more than bonusShortTaken carries. It was chosen as 5 among 0.5
// to 8 with the word table before it was counted again (4.5 gave three
// more right than 3.5, the noun 건지; 5 to 6 five more, the noun 거지; 8 one
// more wrong). With the table counted again and seenShare, any value from
// 3.33 to 7.5 gives the same words right on the answer sheet with the
// three steps above, but at 3.33 the words written onto the modifying
// form in TestBoundAfterRuleOnto stay the nouns.
const bonusOverWhole = 5.0

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

// tableNoun is the name of a noun of the word table (열쇠:체언), and
// nounNames the kinds of the names of the closed sets' nouns (의존명사:수,
// 대명사:나): the rule of a sentence gluedShortRule speaks of such a noun
// of one syllable, under any of the names its letters are read with.
var (
	tableNoun = keyKind + ":" + kindLabel[kindNominal]
	nounNames = []string{"의존명사", "대명사"}
)

// copulaAdnForms are the letters of the copula 이 with a modifying
// closing fused into it (이+ㄴ, 이+ㄹ): the rule of a sentence
// unmodifiedRule prefers a noun of the word table that ends in them where
// nothing an adnominal modifies follows.
var copulaAdnForms = []string{"인", "일"}

// topicNames are the names of 이란 and 란 as the topic particle, which a
// rule of a sentence prefers right after a closing sign (shutTopicRule).
var topicNames = []string{"보조사:이란", "보조사:란"}

// copulaEndNames are the names of the ending 야, which follows the
// copula 이 only, and of the ending 라고, and sentenceEnds the names of
// the last piece of the word right after a word that ends its sentence:
// none (nothing follows), a full stop or a question mark (copulaEndRule).
var (
	copulaEndNames = []string{"맺음:야", "맺음:라고"}
	sentenceEnds   = []string{"", "마침표", "물음표"}
)

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

// forceUnitWord is the force of the rebuttal "the stem begins right
// after a number, and its letters are a unit and a word written after a
// unit" against a known word (5년간 read as 5 and 년간, 6월경 as 6 and
// 월경): the number and its unit (5년, 6월) and the word after it (간,
// 경) come first, as article 43 of the orthography writes a unit with
// its number. It is chosen among a few values tried on the answer
// sheet: 0.25 righted 2 words of the kind, 0.5 and 0.98 righted 17 each
// (the same ones), so the least of those, which still keeps the known
// word as a candidate behind them (a verdict of about 1/3).
// Those were measured before the word table was counted again; with the
// table counted again, against no such rebuttal 0.25 rights 12 words and
// 0.375 to 1 the same 24.
const forceUnitWord = 0.5

// numberAdnNames are the names of the closing after the copula's stem
// that, right after a number, the dependent noun 인 (a count of people)
// is read before (numberAdnRebuttal): the modifying ㄴ (1인, 300인).
var numberAdnNames = []string{"꾸밈:ㄴ"}

// forceNumberAdn is the force of the rebuttal "the copula's modifying
// form right after a number" (1인 read as 1+이+ㄴ rather than 1 and 인).
// It is chosen among a few values tried on the answer sheet: 0.25 changed
// nothing, 0.5 and 0.98 righted the same 9 words and turned the same 2
// wrong; the least of those keeps the copula as the candidate right
// behind.
// Those were measured before the word table was counted again; with the
// table counted again 0.375 and less give 8 fewer than this value (1
// righted, 9 turned wrong), and 0.5 to 1 the same words.
const forceNumberAdn = 0.5

// measureName names the graph of a quotation mark right after a digit
// read as a unit of measure (measureGraph): 3" as inches, 5' as feet.
const measureName = "단위:숫자 뒤 따옴표"

// forceMeasureOpen is the force of the rebuttal "a sign of its kind
// stands open before it" against a quotation mark right after a digit
// read as a unit (measureGraph). It is chosen among a few values tried on
// the answer sheet: 0.1 changed nothing, 0.25 and 0.5 righted the same 5
// words; 0.25 leaves the unit (a verdict of 0.6) within about 0.6 of
// lambda of the closing, 0.5 (a verdict of 1/3, as forceUnitWord and
// forceNumberAdn give) sends it behind with room.
// Those were measured before the word table was counted again; with the
// table counted again 0 gives 5 fewer than this value, and 0.25 to 1 the
// same words.
const forceMeasureOpen = 0.5

// degreeNames are the names of the unit 도 (an angle, a temperature),
// which a rule of a sentence prefers right after a number at the end of
// its word to the particle 도 (degreeUnitRule).
var degreeNames = []string{"단위:도"}

// bonusOverUncounted is what the rule of a demonstrative determiner that
// is also an adjective's modifying form (determinerFormRule) adds: the
// word table counts the adjectives (다르다, 그렇다, 이렇다, 어떻다) one to
// two million times and the determiners of their letters (다른, 그런,
// 이런, 어떤) about once, as the counting read those letters as the
// adjective's form, so that the determiner stands about 5.2 of
// preference below and, rebutted as rarely seen, 3.2 more of lambda
// (8.35 for 다른 and 그런, 8.1 for 어떤, 7.3 for 이런), past each of the
// four steps above. Before a dependent noun (그런 것, 다른 데) the
// adjective's form also gains bonusShortTaken (boundBeforeRule), so that
// the gap is about 11.7. Measured on the answer sheet against no rule,
// 8.5 to 11.5 gave 73 or 74 more right words (96 righted, 23 turned
// wrong), and 11.75, 12, 15, 20 and 30 gave 76 (101, 25); 12 is the least
// round value of that stretch. When the word table counts these
// determiners, the value is to be measured again.
const bonusOverUncounted = 12.0

// bonusNewTaken is what the rebuttal "it is in no key table" (forceNew)
// takes from lambda: a word of markers alone, written apart from the
// word they belong to, is read with a new word's verdict (loneOf),
// 2/(1+forceNew)-1, so the value is about 7.16. The rule of a sentence
// apartParticleRule adds it where the word before shows what the
// particle belongs to. None of the four steps (bonusSameLetters to
// bonusOverWhole) is that far, and bonusOverUncounted is what another
// rule lacks for another reason: on the answer sheet bonusOverWhole (5)
// rights 8 words (none turned wrong) and this value 15 (one turned
// wrong); 이 written apart after a noun, read as the noun or the
// determiner, stands about 7.3 above the particle.
var bonusNewTaken = -Lambda * math.Log10(2/(1+forceNew)-1)

// dativeNames are the names of the case particle 게 that the contracts
// table puts under the pronoun forms 내, 네 and 제 (내게, 네게, 제게),
// dativeStems those forms, and causeStems the stems of the verbs that
// the ending -게 stands right before as what they bring about (-게 하다,
// -게 되다, -게 만들다): a rule of a sentence prefers the pronoun with the
// particle where none of them follows (dativeContractRule).
var (
	dativeNames = []string{"부사격:에게"}
	dativeStems = dativeStemsOf(contracts)
	causeStems  = []string{"하", "되", "만들"}
)

// negLinks are the names of the linking ending -지 (and of 진, -지 with
// the particle 는 fused) and negStems the stems of the verbs that deny
// what the -지 form says (-지 않다, -지 못하다, -지 말다): a rule of a
// sentence prefers the -지 form right before one of them
// (linkBeforeNegRule).
var (
	negLinks = []string{"맺음:지", "맺음:진(지는)"}
	negStems = []string{"않", "못하", "말"}
)

// predStemKinds are the names of a stem of the word table that a
// predicate is read on: a verb and an adjective (predStem; an auxiliary
// verb of the closed set is named by its own stem).
var predStemKinds = []string{keyKind + ":" + kindLabel[kindVerb], keyKind + ":" + kindLabel[kindAdj]}

// objectKinds are the kinds of the names of the object case particle
// (목적격:을, 목적격:를), after which a rule of a sentence
// prefers a predicate (predAfterObjectRule).
var objectKinds = []string{"목적격"}
