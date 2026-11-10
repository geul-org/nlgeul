package token

// Classes of index entry.
const (
	classParticle = iota // 조사
	classSlot            // 선어말 칸
	classClosing         // 맺음
	classKey             // 닫힌 집합의 낱말(열쇠)
)

// Final-consonant indexes, and the index keys of pieces that are not a
// syllable of their own: a final consonant put under the stem's last
// syllable, and 어/아 fused into its vowel.
const (
	jongN  = 4
	jongL  = 8
	jongM  = 16
	jongB  = 17
	jongSS = 20

	leadVowel = -1 // 어/아 fused into the stem's last vowel
	keyVowel  = 'ㅓ'
)

// jamoKeys maps a final-consonant index to its index key.
var jamoKeys = map[int]rune{jongN: 'ㄴ', jongL: 'ㄹ', jongM: 'ㅁ', jongB: 'ㅂ', jongSS: 'ㅆ'}

// Attachment conditions of the grammar tables (the Marker's After).
const (
	afterAny = iota
	afterBatchim
	afterHard
	afterOpen
	afterVowel
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
