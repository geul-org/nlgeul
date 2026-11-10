package token

import (
	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
)

// Fused letters that are common inside words too (decision 014-30).
// Their readings are raised next to the others, never instead of them,
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
var shortClosings = []short{
	{parse.Marker{Form: "대", Kind: rule.KindFinal}, "다고 해", true, rule.TailAdj},
	{parse.Marker{Form: "대", Lead: rule.JongN, Kind: rule.KindFinal}, "ㄴ다고 해", true, rule.TailVerb},
	{parse.Marker{Form: "는대", After: rule.AfterBatchim, Kind: rule.KindFinal}, "는다고 해", true, rule.TailVerb},
	{parse.Marker{Form: "래", Eu: true, Kind: rule.KindFinal}, "라고 해", false, rule.TailVerb},
	{parse.Marker{Form: "재", Kind: rule.KindFinal}, "자고 해", true, rule.TailVerb},
	{parse.Marker{Form: "단", Kind: rule.KindAdn}, "다는", true, rule.TailAdj},
	{parse.Marker{Form: "단", Lead: rule.JongN, Kind: rule.KindAdn}, "ㄴ다는", true, rule.TailVerb},
	{parse.Marker{Form: "는단", After: rule.AfterBatchim, Kind: rule.KindAdn}, "는다는", true, rule.TailVerb},
	{parse.Marker{Form: "란", Eu: true, Kind: rule.KindAdn}, "라는", false, rule.TailVerb},
	{parse.Marker{Form: "잔", Kind: rule.KindAdn}, "자는", true, rule.TailVerb},
	{parse.Marker{Form: "냬", Kind: rule.KindFinal}, "냐고 해", false, 0},
	{parse.Marker{Form: "진", Kind: rule.KindLink}, "지는", false, 0},
	{parse.Marker{Form: "긴", Kind: rule.KindNoun}, "기는", false, 0},
	{parse.Marker{Form: "길", Kind: rule.KindNoun}, "기를", false, 0},
}

// shortParticles are particles shortened with 는: 선 (에서는) and, after
// a vowel or ㄹ, 론 (로는). They close a chain as 엔 and 에선 do.
var shortParticles = []entry{
	{name: "부사격:선(에서는)", form: "선", class: classParticle, slot: rule.LinkTopic, kase: rule.CaseAdv},
	{name: "부사격:론(로는)", form: "론", class: classParticle, after: afterOpen, slot: rule.LinkTopic, kase: rule.CaseAdv},
}

// shortSlots are the same quotations shortened with the past fused in:
// 댔 (다고 했), 랬 (라고 했), 쟀 (자고 했). Each stands where a slot
// stands, before a closing (간댔어, 하랬는데, 가쟀다), after the stems
// 대, 래 and 재 follow. The dictionary lists none of them; they are
// read as 해 → 했 is.
var shortSlots = []entry{
	{name: "선어말:댔(다고 했)", form: "댔", slot: rule.SlotPast, tail: rule.TailAdj},
	{name: "선어말:ㄴ댔(ㄴ다고 했)", form: "댔", lead: rule.JongN, slot: rule.SlotPast, tail: rule.TailVerb},
	{name: "선어말:는댔(는다고 했)", form: "는댔", after: rule.AfterBatchim, slot: rule.SlotPast, tail: rule.TailVerb},
	{name: "선어말:랬(라고 했)", form: "랬", after: rule.AfterOpen, slot: rule.SlotPast, tail: rule.TailVerb},
	{name: "선어말:으랬(라고 했)", form: "으랬", after: rule.AfterBatchim, slot: rule.SlotPast, tail: rule.TailVerb},
	{name: "선어말:쟀(자고 했)", form: "쟀", slot: rule.SlotPast, tail: rule.TailVerb},
}
