package token

import "github.com/geul-org/koparse/pkg/rule"

// kindOrder is the order kinds are named in; kindLabel their names in
// a graph's name (열쇠:체언).
var kindOrder = []int{kindNominal, kindVerb, kindAdj, kindAdverb, kindDeterminer, kindAlone}
var kindLabel = map[int]string{kindNominal: "체언", kindVerb: "동사", kindAdj: "형용사", kindAdverb: "부사", kindDeterminer: "관형사", kindAlone: "홀로"}

// kindTails says what may follow a stem of each kind: a nominal stands
// alone or takes particles or the copula; a verb stem takes slots and any closing,
// an adjective stem slots and the closings that are not the verbs'
// own, and neither stands alone; an adverb stands alone or takes
// auxiliary particles; a determiner and the other words stand alone.
var kindTails = map[int]int{
	kindNominal:    rule.TailBare | rule.TailAux | rule.TailCase | rule.TailCopula,
	kindVerb:       rule.TailPred | rule.TailVerb,
	kindAdj:        rule.TailPred | rule.TailAdj,
	kindAdverb:     rule.TailBare | rule.TailAux,
	kindDeterminer: rule.TailBare,
	kindAlone:      rule.TailBare,
}

// partKinds folds a word table's parts of speech into kinds; a part of
// speech not listed here stands alone.
var partKinds = map[string]int{
	"명사": kindNominal, "대명사": kindNominal, "수사": kindNominal, "식별자": kindNominal,
	"동사": kindVerb, "보조용언": kindVerb, "형용사": kindAdj,
	"부사": kindAdverb, "관형사": kindDeterminer,
}

// closedKinds gives the kind of each closed set of the grammar tables.
// The auxiliary verbs are taken as verbs and as adjectives (있다 is
// both).
var closedKinds = map[string]int{"의존명사": kindNominal, "보조 용언": kindVerb | kindAdj, "관형사": kindDeterminer, "접속 부사": kindAdverb}

// verbClosings are the closings only verbs take: the present ㄴ다/는다
// (also before a quotation's 고, 면, 는), 는구나, 는군, 느냐, 느라; the
// proposal 자, ㅂ시다/읍시다; purpose and intent 러, 려, 려고, 려면,
// 려다, 고자. The modifying 는 and 는데, 는지, 는가 are not here: 있다
// and 없다 take them.
var verbClosings = []string{
	"맺음:ㄴ다", "맺음:ㄴ다고", "맺음:ㄴ다면", "꾸밈:ㄴ다는", "맺음:는다", "맺음:는다고", "맺음:는다면", "꾸밈:는다는",
	"맺음:는구나", "맺음:는군", "맺음:느냐", "맺음:느라", "맺음:느라고",
	"맺음:자", "맺음:자마자", "맺음:자고", "맺음:ㅂ시다", "맺음:읍시다",
	"맺음:러", "맺음:으러", "맺음:려", "맺음:으려", "맺음:려고", "맺음:으려고", "맺음:려면", "맺음:으려면", "맺음:려다", "맺음:으려다", "맺음:고자",
}
