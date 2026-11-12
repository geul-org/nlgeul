//ff:type feature=hypo type=model
//ff:what Score — what a hypothesis is sorted by, and what is printed for it: its verdict and its preference (order says how they come from the readings picked)
package hypo

type Score struct {
	Verdict float64
	Pref    float64
}
