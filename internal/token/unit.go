//ff:type feature=token type=model
//ff:what unit — one stretch of a sentence that is read on its own: its letters, how many of them at its head are a sign of a pair or a span of code (not read; a particle may be written onto it), and what that head is called
package token

// head is 0 and name "" for a plain unit, which is the text between
// two spaces or signs of a pair.
type unit struct {
	text string
	head int
	name string
}
