package token

import (
	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// gatedClosings are closings of this layer read only with a word
// table, after a predicate's stem of the table: the naming closing 기
// (먹기, 있기), which without that would be read at the end of every
// word that ends in 기. nounHeads are the names of the gated closings'
// entries: the naming closings a dropped copula stem stands before
// (dropNaming, dropAdn, droppableKnown: 문제+(이)+기+도), 기 only.
var gatedClosings = []parse.Marker{
	{Form: "기", Kind: rule.KindNoun, NoYo: true},
}

var nounHeads = closingNames(gatedClosings...)

// chainHeads are the naming closings a particle chain is read right
// after (있기+는, 받기+도, 있음+을, 됨+을): the word so made is a
// nominal. They are nounHeads with the naming closings 음 and ㅁ of the
// closing table, and the graph of that chain and namingEnd are made
// from this list. The copula's stem is not dropped before 음 or ㅁ
// (이다 with ㅁ is written 임), so the three lists of a dropped stem
// keep nounHeads.
var chainHeads = append(append([]string{}, nounHeads...), namingPrefix+"음", nounName)

// haoClosings are the closings of the 하오체 (소, 오, 으오, 올시다),
// read only with a word table, after a predicate's stem of the table,
// as gatedClosings are: 소 and 오 end many words (장소, 연구소, 라디오),
// and as table rows they would stand after a stem that is not in the
// table (접골시술소). They are not in gatedClosings, whose names are
// those of the naming closings. 오 stands after a vowel only: a stem's
// ㄹ goes before it (노오 for 놀+오, clause 18), which is not read here.
// 올시다 follows the copula's stem 이 or 아니 (아니올시다, 것이올시다).
var haoClosings = []parse.Marker{
	{Form: "소", After: rule.AfterBatchim, Kind: rule.KindFinal, NoIda: true, NoYo: true},
	{Form: "오", After: rule.AfterVowel, Kind: rule.KindFinal, NoYo: true},
	{Form: "으오", After: rule.AfterHard, Kind: rule.KindFinal, NoIda: true, NoYo: true},
	{Form: "올시다", Kind: rule.KindFinal, NoYo: true, NeedIda: true},
}

// looseClosings are the modifying closings 은 and 을 written onto a
// stem that ends in ㄹ without dropping the ㄹ (줄은 for 준, 줄을 for
// 줄), against clause 18 of the orthography (a stem's ㄹ goes before
// ㄴ and ㄹ, and such a stem takes no 으). They are read only with a
// word table, after a predicate's stem of the table that ends in ㄹ,
// next to the candidates the tables give, and are rebutted as a form
// that does not fit what stands before it is (forceClash).
var looseClosings = []parse.Marker{
	{Form: "은", After: rule.AfterLiquid, Kind: rule.KindAdn, NoIda: true, NoYo: true},
	{Form: "을", After: rule.AfterLiquid, Kind: rule.KindAdn, NoIda: true, NoYo: true},
}

const nounAuxName = "사슬:이름 꼴 뒤 조사"

// hangulBound are the closings that are written only onto a Hangul
// syllable: the modifying closings 은, 는 and 을 and the closings 아, 어,
// 나 and 여. After a run of Latin letters and digits, a number, a span
// of code, a path or a file name the same letters are a particle
// (push는, 3은, v2나), and a predicate made of such a run has 하 or the
// copula's stem between the two (push하는, API인, API여서). The
// closings 다, 고, 면, 지 and 네 are not here: before them the copula's
// stem may be dropped after such a run. A closed table: only the
// closings listed are held back.
var hangulBound = closingSets[setHangul]
