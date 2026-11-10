//ff:type feature=token type=model
//ff:what piece — one marker read at one place: its text, the names of the graphs that stand for it, their highest verdict and its preference, what kind of tail it makes after a stem, and whether it leaves no trace in the letters
package token

type piece struct {
	text    string
	names   string
	verdict float64
	pref    float64
	tail    int
	unseen  bool
}
