//ff:type feature=token type=model
//ff:what index — the token index: from a form's last token (or the key of a fused piece) to the entries that end there, the closed-set words with their kinds and what may follow them, whether a word table is given, the word table and how it is used, its misspelt forms (by the written form), the changed forms of its predicates' stems (each with the stems it stands for) and the last syllables of those that need the stem's own last syllable to be judged by, the graphs of stems by kind with the verdicts already worked out, the graph the word before the vocative 야 must pass, the graph of an auxiliary particle after a linking closing, the graph of a particle after a closing that asks, the forms of the closings, the graphs of the copula stem (said and dropped) and the closings it does not take, the graphs of the dependent nouns the copula stem is fused into, the graphs of the full stop's candidates, the graphs of the candidates of the signs of a pair and of the named signs at a unit's edge (each by the sign), the graphs of the signs that set two words side by side in one stem, the verdict of an unlisted run of Latin letters and digits, the graph of a new word and its verdict, the graph of a stem's restored last syllable, and the total observed counts, and the spacing rules that are on (by name), and the keys the tables hold that are longer than a new word's stem may be (farKeys, with each of them less its last letter in farBases and each of their beginnings in keyPrefixes, and keyMax, the longest key of any)
package token

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// layers is layerNone when there is no word table. memo keeps the
// verdicts of the graphs for what they were asked (judged).
type index struct {
	byKey       map[rune][]*entry
	words       map[string]*entry
	closed      map[string]int
	takes       map[string]int
	aux         []string
	lex         map[string]word
	typos       map[string][]typo
	lexTotal    float64
	alts        map[string][]alt
	irrEnds     map[rune][]alt
	layers      int
	lexed       bool
	kinds       map[int]*toulmin.Graph
	memo        map[judgeKey]float64
	barred      []string
	copula      *entry
	call        *entry
	linkAux     *entry
	askAux      *entry
	nounAux     *entry
	closingSet  map[string]bool
	dropped     *entry
	honor       *entry
	eu          *entry
	neg         *entry
	ji          *entry
	fixed       []*entry
	nouns       []*entry
	contracts   []contract
	wide        map[rune]*entry
	pred        *entry
	nominal     *entry
	heads       map[rune][]string
	signs       map[rune][]*entry
	fronts      map[rune]*entry
	after       map[rune][]*entry
	tilde       *entry
	number      float64
	afterPrefs  map[bool]float64
	joins       []*entry
	stops       []*entry
	links       map[rune][]*entry
	pairs       map[rune][]*entry
	marks       map[rune][]*entry
	run         float64
	stem        *toulmin.Graph
	restore     *toulmin.Graph
	unknown     float64
	total       float64
	counted     map[string]bool
	graphs      int
	edges       int
	spacing     map[string]bool
	farKeys     map[string]bool
	farBases    map[string]bool
	keyPrefixes map[string]bool
	keyMax      int
	ends        map[rune][]*entry
	noWindow    bool
}
