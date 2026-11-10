package mini

// particles is particles.md moved into code: sections 1 to 4, with the
// counts written there (1 where the table says observed without a
// number). Left out: forms with no observation, forms that are also a
// stem plus an ending (하고, 보고, 치고, 말고, 따라, 다가, 마는), the
// quoting 고 and 이다 (read as endings), and the one-consonant forms.
// The third field is the particle's place in a chain (section 5): 0 for
// the suffix-like 들, 씩, 쯤, 끼리, then 1 to 5. The fourth is the case
// the tables give the particle (0 for an auxiliary, a connective or a
// suffix); 에서 and 서 mark both the subject of a group and a place.
var particles = []particle{
	{"이", afterBatchim, 3, caseSubj | caseComp, 5530000}, {"가", afterVowel, 3, caseSubj | caseComp, 5530000},
	{"께서", afterAny, 3, caseSubj, 3955}, {"에서", afterAny, 1, caseSubj | caseAdv, 2820000}, {"서", afterAny, 3, caseSubj | caseAdv, 891},
	{"을", afterBatchim, 3, caseObj, 9690000}, {"를", afterVowel, 3, caseObj, 9690000},
	{"의", afterAny, 3, caseAdn, 8140000}, {"에", afterAny, 1, caseAdv, 5410000},
	{"에게", afterAny, 1, caseAdv, 440000}, {"한테", afterAny, 1, caseAdv, 2705}, {"께", afterAny, 1, caseAdv, 2939},
	{"에게서", afterAny, 1, caseAdv, 13243}, {"한테서", afterAny, 1, caseAdv, 17},
	{"으로부터", afterHard, 1, caseAdv, 130000}, {"로부터", afterOpen, 1, caseAdv, 130000},
	{"으로", afterHard, 1, caseAdv, 4770000}, {"로", afterOpen, 1, caseAdv, 4770000},
	{"으로서", afterHard, 1, caseAdv, 140000}, {"로서", afterOpen, 1, caseAdv, 140000},
	{"으로써", afterHard, 1, caseAdv, 59000}, {"로써", afterOpen, 1, caseAdv, 59000},
	{"과", afterBatchim, 1, caseAdv, 2890000}, {"와", afterVowel, 1, caseAdv, 2890000},
	{"이랑", afterBatchim, 1, caseAdv, 1}, {"랑", afterVowel, 1, caseAdv, 1},
	{"보다", afterAny, 1, caseAdv, 120000}, {"처럼", afterAny, 1, caseAdv, 62000},
	{"같이", afterAny, 1, caseAdv, 1394}, {"만큼", afterAny, 1, caseAdv, 8428},
	{"이라고", afterBatchim, 1, caseAdv, 100000}, {"라고", afterVowel, 1, caseAdv, 100000},
	{"에다", afterAny, 1, caseAdv, 1461}, {"에다가", afterAny, 1, caseAdv, 1461},
	{"아", afterBatchim, 4, caseVoc, 7037}, {"야", afterVowel, 4, caseVoc, 7037},
	{"이여", afterBatchim, 4, caseVoc, 978}, {"여", afterVowel, 4, caseVoc, 978},
	{"이시여", afterBatchim, 4, caseVoc, 163}, {"시여", afterVowel, 4, caseVoc, 163},
	{"이며", afterBatchim, 4, 0, 822}, {"며", afterVowel, 4, 0, 822},
	{"이나", afterBatchim, 4, 0, 270000}, {"나", afterVowel, 4, 0, 270000},
	{"은", afterBatchim, 4, 0, 8160000}, {"는", afterVowel, 4, 0, 8160000},
	{"도", afterAny, 4, 0, 940000}, {"만", afterAny, 2, 0, 230000},
	{"까지", afterAny, 2, 0, 220000}, {"부터", afterAny, 2, 0, 120000},
	{"마저", afterAny, 2, 0, 4613}, {"조차", afterAny, 2, 0, 7887},
	{"밖에", afterAny, 2, 0, 8433}, {"뿐", afterAny, 2, 0, 21000}, {"마다", afterAny, 2, 0, 27000},
	{"이나마", afterBatchim, 4, 0, 981}, {"나마", afterVowel, 4, 0, 981},
	{"이라도", afterBatchim, 4, 0, 2576}, {"라도", afterVowel, 4, 0, 2576},
	{"이든", afterBatchim, 4, 0, 9444}, {"든", afterVowel, 4, 0, 9444},
	{"이든지", afterBatchim, 4, 0, 3223}, {"든지", afterVowel, 4, 0, 3223},
	{"이든가", afterBatchim, 4, 0, 16}, {"든가", afterVowel, 4, 0, 16},
	{"이야", afterBatchim, 4, 0, 7533},
	{"이야말로", afterBatchim, 4, 0, 948}, {"야말로", afterVowel, 4, 0, 948},
	{"이란", afterBatchim, 4, 0, 7116}, {"란", afterVowel, 4, 0, 7116},
	{"커녕", afterAny, 4, 0, 66}, {"대로", afterAny, 2, 0, 17000}, {"요", afterAny, 5, 0, 2894},
	{"토록", afterAny, 2, 0, 518}, {"인들", afterAny, 4, 0, 25},
	{"들", afterAny, 0, 0, 1}, {"씩", afterAny, 0, 0, 1}, {"쯤", afterAny, 0, 0, 1}, {"끼리", afterAny, 0, 0, 1},
	{"엔", afterAny, 4, caseAdv, 305}, {"에선", afterAny, 4, caseAdv, 799}, {"에겐", afterAny, 4, caseAdv, 1},
	{"한텐", afterAny, 4, caseAdv, 1}, {"까진", afterAny, 4, 0, 1}, {"보단", afterAny, 4, caseAdv, 3},
}
