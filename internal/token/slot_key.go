//ff:type feature=token type=model
//ff:what slotKey — a word at a place: a span and the word as it is printed, to put a word that stands at a place in many ways of cutting there once
package token

type slotKey struct {
	span Span
	text string
}
