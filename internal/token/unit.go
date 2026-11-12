//ff:type feature=token type=model
//ff:what unit — one stretch of a sentence that is read on its own: its letters, how many of them at its head are a sign of a pair or a span of code (not read; a particle may be written onto it), what that head is called, the candidates of a sign of a pair at its head (one for each depth it may open or close at) and whether that sign is also read as its letter alone, whether the unit begins an input line, whether a space stands right after it, whether its head is a sign of the unit table right after a digit (3" or 5': it may be read as a unit of measure), and the trace lines of the graphs that judged its head
package token

// head is 0 and name "" for a plain unit, which is the text between
// two spaces or signs of a pair.
type unit struct {
	text    string
	head    int
	name    string
	reads   []pairing
	bare    bool
	first   bool
	spaced  bool
	notes   []string
	measure bool
}
