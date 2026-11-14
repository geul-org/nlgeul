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

var pairBars = map[string]int{"헷갈림": pairUnsure, "낱말 속 안 됨": barWord, "숫자 뒤 안 됨": barDigit, "띄어 쓴 자리 안 됨": barApart}

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
// (contracts.tsv), of this layer's own closed sets of words
// (closed-words-more.tsv) and of closings by what they take
// (closing-sets.tsv).
const (
	setContract    = "준말"
	setPronoun     = "한 음절 대명사"
	setNounCopula  = "의존명사+서술격"
	setTypo        = "오표기"
	setBoundAux    = "의존명사 보조 용언"
	setAuxMore     = "붙여 쓰는 보조 용언"
	setPronounWord = "대명사"
	setSpoken      = "구어 꼴"
	setVerbOnly    = "동사만"
	setDrop        = "서술격 빠짐"
	setDropKnown   = "서술격 빠짐(낱말 표)"
	setLinkTaker   = "이음 으"
	setHangul      = "한글 뒤만"
)

var contractSets = map[string]int{setContract: 0, setPronoun: 0, setNounCopula: 0, setTypo: 0}

// The sets of the table of the words the spacing rules stand on
// (spacing-words.tsv): the words that link two nominals or close a
// list, the titles and offices that follow a name, the family names
// (only marked, never made to stand), the dependent nouns that follow
// a nominal and those that follow a modifying closing, and the
// auxiliaries that are written apart.
const (
	setLink      = "이음·열거"
	setTitle     = "호칭·관직"
	setSurname   = "성"
	setBoundNoun = "의존 명사(체언 뒤)"
	setBoundAdn  = "의존 명사(꾸밈 뒤)"
	setAuxApart  = "보조 용언"
)

// slotLabel begins the name of a slot before a closing.
const slotLabel = "선어말:"
