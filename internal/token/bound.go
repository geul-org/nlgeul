//ff:func feature=token type=rule control=sequence
//ff:what bound — is a stem words bound by a slash or a hyphen (입력/출력, 원-달러): a candidate for each candidate of the sign that stands, with its verdict no higher than that of the weakest word (none when the stem holds no such sign, or a word of it is not known; the last word must take what follows the stem)
package token

func (sc *scan) bound(stem string, tail int) []markCandidate {
	_, candidates := sc.boundOf(stem, tail, false)
	return candidates
}
