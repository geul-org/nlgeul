package combo

// Limit is how many combinations a sentence may have and still be listed.
// The combinations of a sentence are all held and sorted at once, so a
// sentence with more than this many is only counted, not listed. It is
// a guard against running out of memory, not a measured value: change
// this one number to move it.
const Limit = 1000000
