//ff:type feature=rule type=model
//ff:what Shape — a criterion for the runs of digits a sign joins: how many there may be, how long each may be, the values each may have, and how the first must begin
package rule

// Of names the reading (날짜:-). Min and Max bound the number of runs
// (Max 0: no bound). Lens and Vals give, run by run, the least and the
// most digits and the least and the greatest value; the last pair
// holds for every run after it, a most of 0 is no bound, and an empty
// list asks nothing. Head is what the first run must begin with.
type Shape struct {
	Of   string
	Min  int
	Max  int
	Lens [][2]int
	Vals [][2]int
	Head string
}
