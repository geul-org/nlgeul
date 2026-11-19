package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

// endings are the closings of the table closings.tsv in the data
// package. Its first 115 rows are from endings.md sections 1, 3 and 4
// and morphemes.md sections 2 and 3. Counts are the ones written in the
// table (0 where it gives none). The fourth field is the closing's kind
// as section 3 marks it (종, 연, 꾸, 이). Of the naming forms only ㅁ
// and 음 are read (a sentence may end in them); 기 and the rare
// closings 오, 소, 세, 마, 쇼, 담, 곤, ㄴ바, ㄴ즉 and ㄹ지언정 are left
// out. The rows after them, with no count, are endings the standard
// dictionary lists that the token layer read on top of these: 되,
// 어야지, 어야만, 어다가, the quoting endings with 면서, 며 and 니
// (다면서, ㄴ다면서, 는다면서, 라면서, 자면서, ㄴ다며, 는다며, 라며,
// 자며, 다니, ㄴ다니, 는다니, 라니), 길래, 든가, ㄴ들, ㄹ망정,
// ㄹ지라도, ㄹ뿐더러, ㄹ세라, 으라고, 다만, ㄴ다만, 는다만, 건만,
// 건대, 느냐고, ㄴ지라, 는지라, 은지라, 답니다, ㄴ답니다, 는답니다,
// 랍니다, and the modifying 려는 and 려던 (the dictionary's shortened
// 려고 하는, 려고 하던), and the quoting endings that are shortened
// (냐는, 느냐는, 자는, 냐며, 느냐며, 냐면, 느냐면, 냐면서, 느냐면서,
// 냐니까, 느냐니까, 자면, 자니, 자니까, 다니까, ㄴ다니까, 는다니까,
// 라니까, 다더라, ㄴ다더라, 는다더라), then 거라, 너라, 렴, 나이다,
// 나이까, 더이다, 소서, 노라고, 노라면, 거니와, 던지, 다거나, 로구나,
// 로다, ㄴ걸 and 오리다; then endings of the 하오체 and 하십시오체:
// 습디다, ㅂ디다, 습디까, ㅂ디까, 구려, 소이까, 오니까, 오리이까, 렷다,
// 리다, 리까 and 노니; then endings of intent and guess and other
// linking and final ones: 려나, 려니, 려다가, 려거든, 랴, 으되 (되
// has no form with 으), 건마는, ㄹ진대, 을진대, ㄹ는지, 을는지, ㄴ대야,
// 는대야, ㄴ담, ㄴ고, 다네 and the modifying 으라는; then the linking
// 요 of the copula's stem and 아니 (학문이요, 아니요), and the short 매,
// 람, 든 (the shortened 든지, next to the particle 든 of the same
// letters) and 어다; then 거니. Not listed, though the dictionary has them: 는걸
// and 은걸, which the token layer reads as a modifying closing with 거
// under 를; 리, which as a row stands after a stem that is not in the
// table; and 마, which the token layer reads after a predicate's stem
// of the word table only.
var endings = data.Must(data.Each(closingsFile, 5, endingOf))
