//ff:type feature=rule type=model
//ff:what Mark — the structural facts of one sign (a full stop) at one place: what stands on each side, the label to its right and the last label of the dotted name it is in, whether that name is made of numbers, how many dots it holds, and how many dots stand in a row here
package rule

// Left and Right are Side… flags (0 when nothing, or something else,
// stands there). Ext, Last, Numeric, VHead and Dots describe the name
// a dot between two runs is part of (task.py, a.b.com, v1.2.3): Ext is
// the run right of the dot, Last the name's last run, Numeric tells
// that every run is digits (the first may begin with v), VHead that
// the first run begins with v. Run is the number of dots in a row.
// Parts are the runs the sign joins with the same sign next to it.
type Mark struct {
	Left    int
	Right   int
	Ext     string
	Last    string
	Numeric bool
	VHead   bool
	Dots    int
	Run     int
	Parts   []string
}
