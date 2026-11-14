//ff:type feature=token type=model
//ff:what netStats — what the nets of a stream count of themselves, for measuring (not sent up): how many cuts the window made, how wide the window was at each (tokens from the last cut to the place worked out), the places worked out when each net sent its first stretch, the most states a net held open at once, how many cuts had no state after the cut on the last word of the way (blocked), how many of those the states that end at the cut gave a way on (widened), and how many stretches were sent as they are for want of any reading (unread)
package token

type netStats struct {
	cuts    int
	widths  []int
	firsts  []int
	states  int
	blocked int
	widened int
	unread  int
}
