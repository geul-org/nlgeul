//ff:type feature=token type=model
//ff:what wordScore — the verdict and the preference and the number of its making (born) of one word of a branch, kept beside the branch's words so that a branch can be cut into its words again
package token

type wordScore struct {
	verdict float64
	pref    float64
	born    int
}
