package token

import "github.com/geul-org/koparse/pkg/rule"

// Classes of index entry.
const (
	classParticle = iota // 조사
	classSlot            // 선어말 칸
	classClosing         // 맺음
	classKey             // 닫힌 집합의 낱말(열쇠)
)

// Final-consonant indexes, the lead of a fused 어/아 and the attachment
// conditions are the rule package's. keyVowel is the index key of 어/아
// fused into a stem's last vowel; a final consonant put under the
// stem's last syllable is keyed by its letter (rule.LeadText).
const (
	jongN  = rule.JongN
	jongM  = rule.JongM
	jongSS = rule.JongSS

	leadVowel = rule.LeadVowel
	keyVowel  = 'ㅓ'

	afterAny     = rule.AfterAny
	afterBatchim = rule.AfterBatchim
	afterOpen    = rule.AfterOpen
)

// beam is how many partial readings are kept at each place, and
// branches how many readings of a unit are sent up at most
// (stream-graph-algorithm.md section 12 names four to seven).
const (
	beam     = 16
	branches = 7
)

// syllables is the number of Hangul syllables: with no key table, a
// stem syllable may be any of them, so each costs log10(1/syllables)
// of preference.
const syllables = 11172

// newWord and noKey name the graph of a stem that is not in any table.
const (
	newWord = "새 낱말"
	noKey   = "열쇠 표에 없다"
)

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

// Layers of a word table's use: whether a stem is in the table; whether
// its kind fits what follows it; and the cost of structure.
const (
	layerNone = iota
	layerKnown
	layerKind
	layerCost
)

// The costs of structure, in the units of preference (log10). They are
// placeholders, not measured values: nothing in the tables says how
// much a boundary, an unseen piece or a piece is worth. Each is paid
// once: for every word after a unit's first, for every piece that
// leaves no trace in the letters (어/아 fused into ㅏ ㅓ ㅐ ㅔ with
// nothing after it), and for every stem and marker of a word.
// How strong the evidence is that a stem found in the word table is
// a word (the verdict of its kind's graph). These too are placeholders,
// not measured values. A stem observed seenEnough times or more is not
// rebutted as rarely seen; forceFew and forceShort are the forces of
// the rebuttals "rarely seen" and "a key of one syllable".
const (
	seenEnough = 10000
	forceFew   = 1.0
	forceShort = 0.25
	shortKey   = 1
)

const (
	costBoundary = 1.0
	costUnseen   = 1.0
	costPiece    = 1.0
)
