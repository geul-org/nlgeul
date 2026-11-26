//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what caseJointMark — whether the letters between the two numbers of a case number's shape end with a joint that hands the first number on to the second (caseJoints) and are no case mark: the particle of a branch number alone (11의2), or a piece read after a number and then a joint (the branch of an article 31조의4, a fraction 3000분의7, the next part of a provision 36조제1항)
package parse

import "strings"

func (sc *scan) caseJointMark(mark string) bool {
	for _, joint := range caseJoints {
		rest, ok := strings.CutSuffix(mark, joint)
		if ok && (sc.readsAfter(rest, 0) || mark == branchJoint) {
			return true
		}
	}
	return false
}
