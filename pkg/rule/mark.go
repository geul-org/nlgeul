//ff:type feature=rule type=model
//ff:what Mark — the structural facts of one sign (a full stop) at one place: what stands on each side, the label to its right and the last label of the dotted name it is in, whether that name is made of numbers, how many of the sign stand in a row here, and the facts its reader tells of its place
package rule

// Left and Right are Side… flags (0 when nothing, or something else,
// stands there). Ext, Last and Numeric describe the name a dot between
// two runs is part of (task.py, a.b.com, v1.2.3): Ext is the run right
// of the dot and Last the name's last run, both in lower case, and
// Numeric tells that every run is digits (the first may begin with v).
// Run is the number of dots in a row. Parts are the runs the sign joins
// with the same sign next to it. Facts are Mark… flags that the letters
// around the sign do not show and its reader sets: where it stands in
// its unit and in the line, and whether a sign of its pair may stand
// open before it.
type Mark struct {
	Left    int
	Right   int
	Ext     string
	Last    string
	Numeric bool
	Run     int
	Parts   []string
	Facts   int
}
