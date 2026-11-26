package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// A particle of a pair that the final consonant of the word before it
// chooses (은/는, 을/를, 이/가, 과/와, 으로/로 …) is now and then written
// with the other form of its pair (민주당는, 골드은): a spelling against
// the norm, not another word. With a word table each such particle is
// read once more where its own form does not fit (addMisfit), as a
// candidate a rebuttal lowers (forceMisfit) and only on a nominal of the
// word table: a wrong particle is told only against a word that is
// known. misfitAfter gives, for the way a particle of the table
// attaches, the way the other form of its pair attaches, which is where
// the particle is read against the norm; a particle that attaches
// anywhere has no pair and no such reading. The pairs are read from the
// particle table (every row with a condition on the word before it), not
// listed here.
var misfitAfter = map[int]int{
	rule.AfterBatchim: rule.AfterVowel,
	rule.AfterVowel:   rule.AfterBatchim,
	rule.AfterHard:    rule.AfterOpen,
	rule.AfterOpen:    rule.AfterHard,
}

// misfitOwn are the particles of ownParticles that are one form of such
// a pair with a row of the particle table (야 with 이야; 은커녕 and
// 는커녕 with each other), read against the norm the same way.
var misfitOwn = []entry{
	{name: "보조사:야", form: "야", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkTopic, misfit: true},
	{name: "보조사:은커녕", form: "은커녕", class: classParticle, after: rule.AfterVowel, slot: rule.LinkTopic, count: 66, misfit: true},
	{name: "보조사:는커녕", form: "는커녕", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkTopic, count: 66, misfit: true},
}

// misfitKey sets a particle read against the norm apart from the other
// entries of its name and form where candidates are filed (groupBase):
// the 야 read only right after a closing has the same name and form and
// passes at the same place, and would hide it.
const misfitKey = "어긋남"

// misfitGate names the graph the word before a particle read against the
// norm must pass: a nominal of the word table (gateOf). It is a graph of
// its own, so that a tail that holds such a particle is told from the
// others by its gate: its letters do not count as markers a new word
// after a boundary would read as (markedFrom), for the reading is itself
// one against the norm.
const misfitGate = "앞말:체언(어긋난 조사 앞)"
