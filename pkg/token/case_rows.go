//ff:type feature=token type=model
//ff:what caseRows — the criterion a graph's case exception carries next to its list of words: the rows of the case exception table for that graph, so that a trace names the row that applied
package token

// The rows are left out of the rule's identifier, as a list's items are.
type caseRows struct {
	Rows []caseException `json:"-"`
}
