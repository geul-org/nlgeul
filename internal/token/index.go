//ff:type feature=token type=model
//ff:what index — the token index: from a form's last token (or the key of a fused piece) to the entries that end there, the closed-set words with their kinds and what may follow them, whether a word table is given, the word table and how it is used, the changed forms of its predicates' stems (each with the stems it stands for) and the last syllables of those that need the stem's own last syllable to be judged by, whether the readings that are asleep are printed, the graphs of stems by kind with the verdicts already worked out, the graphs of the copula stem (said and dropped) and the closings it does not take, the graphs of the full stop's readings, the verdict of an unlisted run of Latin letters and digits, the graph of a new word and its verdict, and the total observed counts
package token

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// layers is layerNone when there is no word table. seen keeps the
// verdict of a kind's graph for the facts it was last asked.
type index struct {
	byKey      map[rune][]*entry
	words      map[string]*entry
	closed     map[string]int
	takes      map[string]int
	lex        map[string]word
	lexTotal   float64
	alts       map[string][]alt
	irrEnds    map[rune][]alt
	layers     int
	lexed      bool
	asleep     bool
	kinds      map[int]*toulmin.Graph
	seen       map[[4]int]float64
	barred     []string
	copula     *entry
	dropped    *entry
	honor      *entry
	eu         *entry
	neg        *entry
	ji         *entry
	fixed      []*entry
	wide       map[rune]*entry
	pred       *entry
	nominal    *entry
	heads      map[rune][]string
	signs      map[rune][]*entry
	fronts     map[rune]*entry
	after      map[rune][]*entry
	tilde      *entry
	number     float64
	afterPrefs map[bool]float64
	joins      []*entry
	stops      []*entry
	run        float64
	stem       *toulmin.Graph
	unknown    float64
	total      float64
	counted    map[string]bool
	graphs     int
	edges      int
}
