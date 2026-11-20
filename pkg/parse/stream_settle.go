//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what settle — let the rules of a sentence judge the regions read and not yet judged, as one line of words (the best way of each): each word in order, with the words before it as they stand after their own judging and the words after it as they were read; what a region's candidates gain is added to them (region.lift) and its trace lines are kept with it; the words of the line then stand behind what is read next. Nothing is done with no rule or no region to judge
package parse

func (s *Stream) settle() {
	line := unjudged(s.queue)
	s.held = false
	if len(s.rules) == 0 || len(line) == 0 {
		return
	}
	seats := seatsOf(line)
	lifts := make([][]lift, len(line))
	for i := range seats {
		ls, notes := s.judgeSeat(seats, i)
		k := seats[i].region
		lifts[k], line[k].notes = append(lifts[k], ls...), append(line[k].notes, notes...)
	}
	for k, rg := range line {
		rg.lift(lifts[k])
		rg.judged = true
	}
	s.behind = nearsBack(seats, s.behind)
}
