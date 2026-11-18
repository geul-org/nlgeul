package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The names this layer's tables of the data package are written with
// (one tsv file a table). They are closed sets: a name that is not
// here is an error when the table is read.
//
// tailNames tell the stems a shortened closing or slot follows (- for
// any predicate's); roleNamesPair what a sign of a pair does; barNames
// what makes such a sign unsure or keeps it from being one.
var tailNames = map[string]int{"-": 0, "형용사": rule.TailAdj, "동사": rule.TailVerb}

var pairRoles = map[string]int{"열기": pairOpens, "닫기": pairShuts}

var pairBars = map[string]int{"헷갈림": pairUnsure, "낱말 속 안 됨": barWord, "띄어 쓴 자리 안 됨": barApart}

// What a sign of a pair does, as the table's column gives it, and the
// bit that marks an unsure sign among what keeps it from being one.
const (
	pairOpens = 1 << iota
	pairShuts
	pairUnsure = 1 << 8
)

// adnPlace is what a fused form read only right after a modifying
// closing is marked with.
const adnPlace = "꾸밈 뒤"

// The sets of the tables that hold more than one list: of fused forms
// (contracts.tsv), of this layer's own closed sets of words (the sets
// of closed-words.tsv that pkg/parse does not read) and of closings by
// what they take (closing-sets.tsv).
const (
	setContract    = "준말"
	setPronoun     = "한 음절 대명사"
	setNounCopula  = "의존명사+서술격"
	setTypo        = "오표기"
	setBoundAux    = "의존명사 보조 용언"
	setAuxMore     = "붙여 쓰는 보조 용언"
	setPronounWord = "대명사"
	setAskWord     = "의문 대명사"
	setExistStem   = "있다·없다·계시다"
	setVerbOnly    = "동사만"
	setDrop        = "서술격 빠짐"
	setDropKnown   = "서술격 빠짐(낱말 표)"
	setLinkTaker   = "이음 으"
	setHangul      = "한글 뒤만"
	setVerbExist   = "동사·있다·없다·계시다"
	setNegAdverb   = "부정 부사"
	setNumDet      = "수 관형사"
	setPointer     = "지시어"
	setPersonal    = "인칭 대명사"
)

var contractSets = map[string]int{setContract: 0, setPronoun: 0, setNounCopula: 0, setTypo: 0}

// The sets of the table of the words the spacing rules stand on
// (spacing-words.tsv): the words that link two nominals or close a
// list, the titles and offices that follow a name, the family names
// (marked, and read as a word of their own only right before a title,
// surnameCandidate), and the auxiliaries that are written apart.
const (
	setLink     = "이음·열거"
	setTitle    = "호칭·관직"
	setSurname  = "성"
	setAuxApart = "보조 용언"
)

// The names of the two columns of the dependent-noun table
// (bound-nouns.tsv) this layer reads: where a dependent noun stands (after
// a modifying closing, adnPlace; after a nominal; after a word of time
// only) and what else reads it (the rebuttal lifted when it is written
// right onto a modifying closing, adnOnto; read as the noun of the same
// letters where no modifying closing stands before it, adnBound; read
// as the noun where neither a modifier nor a nominal stands before it,
// unheldBound).
const (
	nominalPlace = "체언 뒤"
	timePlace    = "시간 뒤"
	useAdnOnto   = "꾸밈 맺음 뒤 붙여 씀"
	useAdnBound  = "꾸밈 뒤가 아니면 체언"
	useUnheld    = "관형어·체언 뒤가 아니면 체언"
)

var (
	boundPlaces = map[string]int{adnPlace: 1, nominalPlace: 2, timePlace: 4}
	boundUses   = map[string]int{useAdnOnto: 1, useAdnBound: 2, useUnheld: 4}
)

// slotLabel begins the name of a slot before a closing.
const slotLabel = "선어말:"

// The names of the case exception table (case-exceptions.tsv) for what
// a row does to the graph it attacks (지움: the reading the graph
// claims is dropped, the only one) and for the kind of its ground (a
// clause of the norms, or what was observed in the corpora).
const (
	caseDrop   = "지움"
	groundNorm = "규범"
	groundSeen = "관찰"
)

var caseStrengths = map[string]int{caseDrop: 0}

var caseGrounds = map[string]int{groundNorm: 0, groundSeen: 0}
