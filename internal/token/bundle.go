//ff:type feature=token type=model
//ff:what bundle — what is sent up for one stretch between agreed boundaries: its text and its branches, best first; more than one branch is a contest (one of them holds)
package token

type bundle struct {
	text     string
	branches []branch
}
