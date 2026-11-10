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
