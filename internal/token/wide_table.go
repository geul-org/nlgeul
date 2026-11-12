package token

import (
	"github.com/geul-org/nlgeul/data"
)

// Fused letters that are common inside words too (decision 014-30).
// Their candidates are raised next to the others, never instead of them,
// and only with a word table: the word before must be a stem of the
// table, a predicate's or (for a shortened particle) a nominal.
//
// wideLetters hold the honorific 시 and what is fused under it: 신
// (시+ㄴ…), 실 (시+ㄹ…), 심 (시+ㅁ), 세 (시+어, before 요). The value
// is the name of the letter's graph.
var wideLetters = map[rune]string{'신': "합침:신", '실': "합침:실", '심': "합침:심", '세': "합침:세"}

const (
	nounHonor   = '심'
	vowelHonor  = '세'
	nounName    = "이름:ㅁ"
	vowelName   = "맺음:어(합침)"
	predGate    = "앞말:용언"
	nominalGate = "앞말:체언"
)

// shortClosings are closings a quotation or a particle is shortened
// into: 대 (다고 해), 래 (라고 해), 재 (자고 해), 단 (다는), 란 (라는),
// 잔 (자는), 냬 (냐고 해), and 진 (지는), 긴 (기는), 길 (기를). Each is one piece,
// named with the form it is short for. 대 and 단 follow an adjective's
// stem directly and a verb's as ㄴ대/는대, ㄴ단/는단; 래, 재, 란 and 잔
// follow a verb's.
var shortClosings = data.Must(data.Each("short-closings", 6, shortOf))

// shortParticles are particles shortened with 는: 선 (에서는) and, after
// a vowel or ㄹ, 론 (로는). They close a chain as 엔 and 에선 do.
var shortParticles = data.Must(data.Each("short-particles", 4, shortParticleOf))

// shortSlots are the same quotations shortened with the past fused in:
// 댔 (다고 했), 랬 (라고 했), 쟀 (자고 했). Each stands where a slot
// stands, before a closing (간댔어, 하랬는데, 가쟀다), after the stems
// 대, 래 and 재 follow. The dictionary lists none of them; they are
// read as 해 → 했 is.
var shortSlots = data.Must(data.Each("short-slots", 4, shortSlotOf))
