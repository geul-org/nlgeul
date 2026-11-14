//ff:type feature=token type=model
//ff:what unitStretch — one stretch of a unit as the stream takes it: its contest (the letters, and the branches when it is read the old way: the signs at the unit's edges, and a unit with a sign of a pair written onto) and, when the net read it, its cells, from which the region is made
package token

type unitStretch struct {
	contest Contest
	cells   *cells
}
