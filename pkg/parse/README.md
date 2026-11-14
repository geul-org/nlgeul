# parse

문장성분의 판단. 이를 부르는 명령은 없다. `Lines(문장)`이 어절마다 `어절:문장성분` 줄을 돌려준다. 어절은 조각 사이에 `+`를 넣어 찍는다(`학교+에서+는`, `곰+이+다`, `가+시+었+다+고+합+니다`).

밖에서 import할 수 있는 패키지다(`github.com/geul-org/nlgeul/pkg/parse`, 옛 자리 `internal/parse`). `pkg/token`이 문법 표를 여기서 읽어서 함께 옮겼다([pkg/README.md](../README.md)).

## 공개 API

| 이름 | 꼴 | 뜻 |
|---|---|---|
| `Lines` | `func Lines(sentence string) ([]string, error)` | 어절마다 `어절:문장성분` 줄 |
| `Marker` | 구조체 | 문법 표의 한 줄(조사나 맺음, 붙는 방식, 관측 횟수)을 다른 패키지가 읽는 꼴 |
| `ParticleMarkers`, `ClosingMarkers` | `func() []Marker` | 조사 표, 맺음 표의 줄들 |
| `ClosedWords` | `func() map[string][]string` | 표에 적힌 닫힌 집합의 낱말, 갈래별 |
| `MarkerOf` | `func(row data.Row, at int) (Marker, error)` | 표의 한 줄에서 맺음 하나를 읽는다 |
| `FormOf` | `func(row data.Row, col int) (int, string, error)` | 표에 적힌 꼴을 받침·모음 머리와 음절로 읽는다 |
| `AfterNames`, `KindNames`, `CaseNames`, `FlagNames` | `map[string]int` | 표의 칸에 쓰는 이름과 그 값 |

## 읽기 (일반 Go 코드, `type=reader`)

문장성분을 정하지 않고 무엇이 읽혔는지만 낸다.

- 표: `particles_table.go`(조사, 앞말 조건, 사슬 자리, 격, 관측 횟수), `endings_table.go`(맺음, 붙는 방식, 갈래 종·연·꾸, 붙는 자리 표시, 관측 횟수). 문서의 표를 옮긴 것이다. 두 변수는 이 파일에 있고 값은 `data/tables/particles.tsv`, `data/tables/closings.tsv`에서 읽는다(`data.Each`, 줄 읽기는 `particleOf`, `endingOf`). 표를 읽는 길과 줄을 더하는 법은 [data/README.md](../../data/README.md)에 있다.
- `readSentence` → 어절마다 `readWord`가 `word`를 낸다.
  - `chain`(`chainFact`): 어절 끝의 가장 긴 조사 사슬. 길이, 조사들(`에서+는`), 끝 조사와 그 관측 횟수, 사슬에서 가장 바깥 격조사의 격.
  - `tails`(`tailFact`): 어절 끝이 용언 꼬리로 읽히는 풀이 전부. 맺음 이름과 갈래, 줄기, 길이, 서술격 "이"가 읽혔는지, 관측 횟수, 그 풀이대로 조각낸 어절(`marked`).
  - `longest`(가장 긴 꼬리의 길이), `last`(마지막 어절인지), `text`.
  - `sent`(`sentence`): 마지막 어절의 줄기들, 앞 어절 가운데 이/가가 붙은 수.
- 길이는 반 음절 단위다. 조각이 든 음절 하나가 2, 줄기 끝 음절에 받침·모음으로 합쳐진 조각이 1이다.

닫힌 집합(`closed_table.go`, 값은 `data/tables/closed-words.tsv`, morphemes.md 8절을 옮긴 것)으로 읽는 것: 꾸밈 맺음 뒤에 붙여 쓴 의존명사(`아+는+거`, `있+는+거+지`)와 "것"의 줄임(`이+게`, `된+다+는+게`), 본용언의 어/아 꼴 뒤 보조 용언 줄기(`작성+해+준`, `해+주+시+고`), 문장 끝의 이름꼴 함·됨·임·음. 어절이 의존명사로 시작하는지(`bound`)와 다음 어절이 그런지(`nextBound`)도 사실로 낸다.

## 판정 (toulmin 그래프, `type=rule`)

`newGraph`가 그래프와 주장들을 돌려준다. 규칙 함수는 대부분 [pkg/rule](../rule/README.md)의 공통 함수이고 기준(Spec)으로 갈린다. 아래 표의 이름은 그래프 노드의 이름(주장은 `Claim` 기준의 이름)이다. 이 패키지에만 있는 규칙 함수는 `complementCase`, `complementShare`, `lastPredicate`, `beforeBound` 넷이다. 읽기가 낸 사실은 `facts`가 `rule.Word`로 옮겨 컨텍스트에 넣는다.  규칙 함수는 모두 `func(Context, Specs) (float64, any)`다. 값은 사실 하나를 보는 규칙이면 +1(적용)·-1(적용 안 됨)이고, 관측 횟수를 견주는 규칙이면 두 횟수 가운데 상대의 몫을 [-1, 1]에 옮긴 값이다.

| 규칙 | 갈래 | 문법 |
|---|---|---|
| `subjectCase` | 주장 → 주어 | 주격 조사(이/가, 께서, 단체의 에서, 사람 수의 서) |
| `objectCase` | 주장 → 목적어 | 목적격 조사(을/를) |
| `adnominalCase` | 주장 → 관형어 | 관형격 조사(의) |
| `adverbialCase` | 주장 → 부사어 | 부사격 조사(에, 에서, 에게, 으로, 과/와, 보다, 처럼, 라고 …) |
| `vocativeCase` | 주장 → 독립어 | 호격 조사(아/야, 이여, 이시여) |
| `complementCase` | 주장 → 보어 | 서술어가 되다·아니다이고 이/가가 붙었다 |
| `topicSubject`, `topicObject` | 주장 → 주어, 목적어 | 격조사 없이 보조사만 붙었다 |
| `closingPredicate` | 주장 → 서술어 | 종결·연결 맺음으로 끝난 용언 꼬리 |
| `lastPredicate` | 주장 → 서술어 | 마지막 어절의 용언 꼬리(입력 전제 1) |
| `adnominalClosing` | 주장 → 관형어 | 꾸밈 맺음(ㄴ/은, 는, ㄹ/을, 던) |
| `adverbialClosing` | 주장 → 부사어 | 맺음 게(기준 목록 `adverbialClosings`) |
| `rangeAdverbial` | 주장 → 부사어 | 격조사 없이 범위의 보조사(까지, 부터)가 붙었다 |
| `determinerWord` | 주장 → 관형어 | 지시·수 관형사가 혼자 섰다(기준 목록 `determiners`) |
| `conjunctiveWord` | 주장 → 독립어 | 접속 부사(기준 목록 `conjunctions`) |
| `wholeWord` | 예외 → 조사 쪽·꼬리 쪽 주장 전부 | 닫힌 집합의 낱말은 쪼개 읽지 않는다 |
| `beforeBound` | 예외 → 조사 쪽 주장 전부, `chainSeenMore` | 뒤 어절이 의존명사로 시작하고 이 어절이 꾸밈 맺음으로 읽힌다 |
| `boundStem` | 예외 → `adnominalClosing`, `tailSeenMore` | 의존명사는 용언 줄기가 아니다 |
| `finalWord` | 예외 → 조사 쪽 주장 전부, `adnominalClosing`, `adverbialClosing`, `chainReadsMore`, `chainSeenMore` | 마지막 어절은 서술어다 |
| `chainReadsMore` | 예외 → 꼬리 쪽 주장 전부 | 어절 끝이 조사 사슬로 더 길게 읽힌다 |
| `tailReadsMore` | 예외 → 조사 쪽 주장 전부 | 어절 끝이 용언 꼬리로 더 길게 읽힌다 |
| `tailCutShort` | 예외 → 꼬리 쪽 주장 전부 | 같은 끝에서 더 긴 꼬리가 읽힌다 |
| `chainSeenMore` | 반박 → 꼬리 쪽 주장 전부 | 같은 길이의 조사가 표에서 관측됐다(횟수의 몫) |
| `tailSeenMore` | 반박 → 조사 쪽 주장 전부 | 같은 길이의 꼬리가 표에서 관측됐다(횟수의 몫) |
| `complementShare` | 반박 → `subjectCase` | 되다·아니다 앞의 이/가 말 n개 가운데 하나는 보어다(몫 1/n) |

`judge`가 꼬리 풀이마다 그래프를 한 번 평가하고(`weigh`), 성분마다 가장 높은 verdict를 남겨 0보다 큰 것을 큰 순서로 `|`로 잇는다. 하나도 없으면 `미정`이다. 선 주장 가운데 조사 쪽 주장이 있으면 조사 사슬 앞과 조사 사이에, 없으면 선 주장이 기댄 가장 긴 꼬리 풀이의 조각대로 찍는다(`noteSplit`, `mark`).

꼬리의 조각(`pieces_table.go`, 값은 `data/tables/closing-pieces.tsv`, `stemMark`, `slotsMark`, `closingMark`, `affixMark`): 줄기, 선어말 칸 하나하나, 맺음, 요가 한 조각씩이다. 줄기와 조각이 한 음절에 합쳐지면 그 음절을 표면 꼴대로 두고(`갔+다`, `합+니다`, `해서`), 문법 조각끼리 합쳐진 음절은 온 음절로 풀 수 있을 때 풀어 찍는다(`셨` → `시+었`, `세요` → `시+어+요`, `죠` → `지+요`). 명사 뒤의 하·되·시키 앞에서 가르고(`지시+해서`), 그 앞말이 인용(다고·라고·자고·냐고)으로 읽히면 그것도 조각낸다.

예외를 더하려면 규칙 함수 파일 하나를 만들고 `newGraph`에 `g.Except(함수).Attacks(c["주장 이름"])` 한 줄을 더한다(`new_graph_test.go`가 테스트 안에서 그렇게 한다). 규칙이 볼 사실이 더 필요하면 `word`나 `sentence`에 칸을 더하고 읽기 코드에서 채운다.
