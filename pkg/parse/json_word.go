//ff:type feature=parse type=model
//ff:what JSONWord — one word (node) of a way as nlgeul parse --format json writes it: its place [start, end) over the letters of the sentence with its spaces left out, its letters, whether no space stood before it, and all its candidates in the order of their lambda
package parse

type JSONWord struct {
	Span       [2]int          `json:"span"`
	Text       string          `json:"text"`
	Glued      bool            `json:"glued"`
	Candidates []JSONCandidate `json:"candidates"`
}
