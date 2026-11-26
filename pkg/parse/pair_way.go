//ff:type feature=parse type=model
//ff:what pairWay — one reading of a sign of a pair in one way the signs before it may stand open: the candidate (no name when the sign is read as its letter alone there) and the way the signs stand open after it
package parse

type pairWay struct {
	read pairing
	way  string
}
