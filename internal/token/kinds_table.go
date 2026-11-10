package token

import "github.com/geul-org/koparse/pkg/rule"

// kindOrder is the order kinds are named in; kindLabel their names in
// a graph's name (열쇠:체언).
var kindOrder = []int{kindNominal, kindVerbal, kindAdverb, kindDeterminer, kindAlone}
var kindLabel = map[int]string{kindNominal: "체언", kindVerbal: "용언", kindAdverb: "부사", kindDeterminer: "관형사", kindAlone: "홀로"}

// kindTails says what may follow a stem of each kind: a nominal stands
// alone or takes particles; a predicate stem takes slots and a closing
// and never stands alone; an adverb stands alone or takes auxiliary
// particles; a determiner and the other words stand alone.
var kindTails = map[int]int{
	kindNominal:    rule.TailBare | rule.TailAux | rule.TailCase,
	kindVerbal:     rule.TailPred,
	kindAdverb:     rule.TailBare | rule.TailAux,
	kindDeterminer: rule.TailBare,
	kindAlone:      rule.TailBare,
}

// partKinds folds a word table's parts of speech into kinds; a part of
// speech not listed here stands alone.
var partKinds = map[string]int{
	"명사": kindNominal, "대명사": kindNominal, "수사": kindNominal, "식별자": kindNominal,
	"동사": kindVerbal, "형용사": kindVerbal, "보조용언": kindVerbal,
	"부사": kindAdverb, "관형사": kindDeterminer,
}

// closedKinds gives the kind of each closed set of the grammar tables.
var closedKinds = map[string]int{"의존명사": kindNominal, "보조 용언": kindVerbal, "관형사": kindDeterminer, "접속 부사": kindAdverb}
