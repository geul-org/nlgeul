//ff:type feature=token type=model
//ff:what piece — one marker read at one place: its text, the names of the graphs that stand for it, their highest verdict and its preference
package token

type piece struct {
	text    string
	names   string
	verdict float64
	pref    float64
}
