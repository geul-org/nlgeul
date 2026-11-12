//ff:type feature=hypo type=model
//ff:what Score — what a combination of candidates is sorted by, and what is printed for it: its place (sorted by), and its verdict and its preference (printed) (order says how they come from the candidates picked)
package hypo

type Score struct {
	Place   float64
	Verdict float64
	Pref    float64
}
