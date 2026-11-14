//ff:type feature=token type=model
//ff:what predKey — one edge of the net: a state and a state before it, so that a state keeps each state before it once without looking through its preds
package token

type predKey struct {
	state  *path
	before *path
}
