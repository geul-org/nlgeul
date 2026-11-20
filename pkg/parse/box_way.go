//ff:type feature=parse type=model
//ff:what boxWay — the boxes of one way of cutting a stretch (a segment's nodes, in order), and once drawn at a width their lines stacked and the line of the title of the last box (where the way's inner bracket ends)
package parse

type boxWay struct {
	boxes []box
	lines []string
	last  int
}
