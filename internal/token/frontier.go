//ff:type feature=token type=model
//ff:what frontier — the walks not yet taken further, the greatest f first and, where f is equal, the one put in first, and the most walks it held at one time (peak, for measuring)
package token

type frontier struct {
	items []*walk
	peak  int
}
