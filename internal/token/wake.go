//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wake — the readings of a unit, best first, parted in two: those awake (within wakeGap of the best reading's place, seven at most), which are sent up, and those asleep, which are kept and not sent up
package token

func wake(paths []*path) ([]*path, []*path) {
	var awake, asleep []*path
	for _, p := range paths {
		if len(awake) < branches && rank(paths[0])-rank(p) <= wakeGap {
			awake = append(awake, p)
			continue
		}
		asleep = append(asleep, p)
	}
	return awake, asleep
}
