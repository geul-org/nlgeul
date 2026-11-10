# rule

toulmin 그래프의 규칙 함수(`func(Context, Specs) (float64, any)`) 가운데 기준(Spec)만 바꿔 여러 노드와 여러 그래프가 함께 쓰는 것을 모았다. 함수가 읽는 사실 타입과 컨텍스트 열쇠도 여기 있다. `internal/`에 기대지 않는다.

## 사실과 열쇠

- `KeyWord` → `*Word`(어절: 글, 조사 사슬 `Chain`, 꼬리 풀이들 `[]Tail`, 가장 긴 꼬리, 마지막 어절인지, 의존명사로 시작하는지, 문장 `*Sentence`).
- `KeyTail` → `Tail`(판정 중인 꼬리 풀이 하나).
- 읽는 쪽이 이 타입을 채워 컨텍스트에 넣는다. 규칙은 `WordOf`, `TailOf`로 꺼낸다.

## 기준

| 기준 | 쓰임 |
|---|---|
| `Claim{Name, Concludes}` | 주장 노드의 이름과 결론. 같은 함수를 여러 주장이 쓸 수 있게 하고, 평가 결과의 이름에서 `ClaimOf`로 되찾는다 |
| `Bits{Of, Mask}` | 격(`Case…`), 맺음 갈래(`Kind…`), 어절의 사실(`WordLast`, `WordBound`) 가운데 어느 것을 볼지 |
| `Versus{Mine, Rival}` | `Longer`가 견줄 두 길이(`SideChain`, `SideTail`, `SideLongest`) |
| toulmin `StringListSpec` | 낱말·조사·맺음의 목록 |

## 규칙 함수

값은 사실 하나를 보는 함수면 +1(적용)·-1(적용 안 됨)이고, 관측 횟수를 견주는 함수면 두 횟수 가운데 상대의 몫을 [-1, 1]에 옮긴 값이다.

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `ChainCase` | rule | 사슬의 가장 바깥 격조사가 기준의 격이다 | `Bits`(case) |
| `ChainNoCase` | rule | 격조사 없이 보조사만 붙었다 | 없음 |
| `ChainListed` | rule | 격조사 없는 사슬에 기준 목록의 조사가 있다 | 조사 목록 |
| `WordListed` | rule, except | 어절이 통째로 기준 목록의 낱말이다 | 낱말 목록 |
| `ClosingKind` | rule | 꼬리의 맺음이 기준의 갈래다 | `Bits`(kind) |
| `ClosingListed` | rule | 꼬리의 맺음이 기준 목록에 있다 | 맺음 목록 |
| `WordIs` | except | 어절이 기준의 사실을 갖는다(마지막 어절, 의존명사로 시작) | `Bits`(word) |
| `Longer` | except | 경쟁 풀이가 이 풀이보다 길게 읽는다 | `Versus` |
| `ChainSeenMore` | counter | 꼬리만큼 긴 조사가 관측됐다(횟수의 몫) | 없음 |
| `TailSeenMore` | counter | 사슬만큼 긴 꼬리가 관측됐다(횟수의 몫) | 없음 |

토큰 열 하나를 표지로 판정하는 그래프(컨텍스트 열쇠 `KeyPiece` → `Piece`)가 쓰는 함수:

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `PieceRead` | rule, counter | 색인에서 걸린 토큰 열이 이 자리에 있다 | 없음 |
| `PieceNoStem` | except | 표지 앞에 줄기가 없다 | 없음 |
| `PieceAfter` | except | 앞 토큰의 끝이 붙는 조건과 안 맞는다(`Fits`) | `Bits`(after) |
| `ChainLinks` | except | 조사 사슬에서 앞 조사의 갈래가 뒤 조사의 갈래를 받지 않는다(`Links`) | 없음 |
| `SlotOrder` | except | 선어말 칸의 순서가 어긋난다(`SlotBefore`) | 없음 |
| `PieceLast` | except | 이 표지 뒤에 이어지는 조각이 없다 | 없음 |
| `StemTail` | except | 줄기의 갈래가 뒤에 붙은 것과 맞지 않는다 | `Bits`(tail) |
| `SeenFew` | counter | 이 낱말은 관측이 적다(적은 정도) | `Bits`(enough) |
| `ShortKey` | counter | 이 낱말은 짧아 우연히 겹치기 쉽다 | `Bits`(syllables) |
| `PieceBase` | except | 합쳐진 조각이 기준 목록에 없는 음절 밑에 있다 | 음절 목록 |

도움 함수: `PieceOf`, `Applies`(예/아니오 → ±1), `Share`(횟수의 몫), `WordOf`, `TailOf`, `ClaimOf`.

기호 하나를 풀이별로 판정하는 그래프(컨텍스트 열쇠 `KeyMark` → `Mark`)가 쓰는 함수:

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `MarkSides` | except | 기호의 양쪽이 기준의 글자 갈래가 아니다 | `Bits`(sides) |
| `MarkEnclosed` | except | 점의 양쪽이 영숫자 열이다 | 없음 |
| `MarkAlone` | except | 점이 하나뿐이다 | 없음 |
| `MarkRepeated` | except | 점이 둘 이상 이어진다 | 없음 |
| `ExtUnlisted` | except | 점 오른쪽 마디가 기준 목록에 없다 | 마디 목록 |
| `LastUnlisted` | except | 이름의 맨 오른쪽 마디가 기준 목록에 없다 | 마디 목록 |
| `VersionShape` | except | 버전의 꼴이 아니다 | 없음 |
| `DecimalShape` | except | 소수의 꼴이 아니다 | 없음 |

## 조건식과 도움 함수 (toulmin 규칙이 아니다)

규칙 함수와 읽기 코드가 같은 문법 조건을 따로 갖지 않도록 조건식을 한 곳에 둔다. `//ff:what` 첫머리에 "조건식" 또는 "도움 함수(규칙 아님)"이라고 적혀 있다.

| 함수 | 뜻 | 쓰는 곳 |
|---|---|---|
| `Fits(prev, after)` | 앞 음절의 끝이 붙는 조건(`After…`)과 맞는가 | `PieceAfter`, 읽기 코드 |
| `Follows(left, right)` | 조사 사슬에서 자리 left 뒤에 자리 right가 올 수 있는가(자리 번호) | `kopas parse`의 읽기 코드 |
| `Links(left, right)` | 조사 사슬에서 갈래 left 뒤에 갈래 right가 올 수 있는가(갈래 표, `Link…`) | `ChainLinks` |
| `SlotOf(r)`, `SlotBefore(slot, next)` | 선어말 칸과 그 순서(시 < 었 < 겠 < 더 < 맺음, 었은 겹침) | `SlotOrder`, 읽기 코드 |
| `BareBases(lead)` | 받침 하나뿐인 맺음을 읽는 음절 | `PieceBase`의 기준, 읽기 코드 |
| `JongOf`, `JungOf` | 음절의 받침·모음 번호 | 읽기 코드 |
| `FusedVowel(r, more)` | 어/아가 합쳐진 음절인가 | 읽기 코드 |
| `FusedPast(r)` | 었/았이 합쳐진 음절인가(받침 ㅆ과 모음 조건) | 읽기 코드 |
| `LeadText(lead)` | 조각의 머리(받침, 합쳐지는 어)를 글자로 | 읽기 코드 |
| `RunRune(r)` | 영문·숫자 열을 잇는 글자인가(라틴 문자, 숫자, 밑줄) | 읽기 코드 |
| `Trim(field)` | 어절 앞뒤의 문장 부호 떼기 | 읽기 코드 |

상수 `After…`, `Jong…`, `LeadVowel`, `Slot…`도 여기 한 벌만 있다.
