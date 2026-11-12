//ff:func feature=token type=rule control=sequence
//ff:what bound — is a stem words bound by a slash or a hyphen (입력/출력, 원-달러): the name of the sign's reading and its verdict, no higher than that of the weakest word (0 when the stem holds no such sign, or a word of it is not known; the last word must take what follows the stem)
package token

func (sc *scan) bound(stem string, tail int) (string, float64) {
	_, name, v := sc.boundOf(stem, tail, false)
	return name, v
}
