//ff:func feature=rule type=helper control=sequence
//ff:what Links — 조건식(규칙 ChainLinks가 쓴다). 조사 사슬에서 갈래 left의 조사 뒤에 갈래 right의 조사가 올 수 있는가(표준 문법의 일반 규칙; 표에 없는 짝은 올 수 없다)
package rule

// linkable[left] is the mask of the kinds that may follow left. A
// suffix takes any particle. A case particle and a topic particle
// take only 요. An adverbial particle takes a range or topic particle,
// 의, a comparing particle and 과/와; a comparing particle and 과/와
// take a range or topic particle and 의. A range particle takes a case
// particle, 의, an adverbial one and the rest. Nothing follows a
// vocative particle or 요. 으로/로 also follows an adverbial particle
// (에게로, 께로부터) and is otherwise as one; the honorific 께서 takes a
// range or topic particle (께서는). The emphatic 가 follows an
// adverbial particle and nothing else, and nothing follows it.
var linkable = map[int]int{
	LinkSuffix: 1<<LinkSuffix | 1<<LinkCase | 1<<LinkAdn | 1<<LinkAdv | 1<<LinkComp | 1<<LinkWith | 1<<LinkRange | 1<<LinkTopic | 1<<LinkVoc | 1<<LinkYo | 1<<LinkDir | 1<<LinkHonor,
	LinkCase:   1 << LinkYo,
	LinkAdn:    1 << LinkYo,
	LinkAdv:    1<<LinkRange | 1<<LinkTopic | 1<<LinkAdn | 1<<LinkComp | 1<<LinkWith | 1<<LinkYo | 1<<LinkDir | 1<<LinkEmph,
	LinkDir:    1<<LinkRange | 1<<LinkTopic | 1<<LinkAdn | 1<<LinkComp | 1<<LinkWith | 1<<LinkYo | 1<<LinkEmph,
	LinkHonor:  1<<LinkRange | 1<<LinkTopic | 1<<LinkYo,
	LinkComp:   1<<LinkRange | 1<<LinkTopic | 1<<LinkAdn | 1<<LinkYo,
	LinkWith:   1<<LinkRange | 1<<LinkTopic | 1<<LinkAdn | 1<<LinkYo,
	LinkRange:  1<<LinkCase | 1<<LinkAdn | 1<<LinkAdv | 1<<LinkComp | 1<<LinkWith | 1<<LinkRange | 1<<LinkTopic | 1<<LinkYo | 1<<LinkDir,
	LinkTopic:  1 << LinkYo,
}

func Links(left, right int) bool {
	return linkable[left]&(1<<right) != 0
}
