# pkg

밖에서 가져다 쓸 수 있는 패키지다. `internal/`에 기대지 않는다.

| 패키지 | 내용 |
|---|---|
| [token/](token/README.md) | 토큰 층. 낱말 표를 읽어 `Reader`를 만들고, 문장을 스트림으로 열어 구간마다 갈래를 λ 차례로 하나씩 받는다(`Next`·`Again`, 후보 수 한도를 거는 `NextHead`·`AgainHead`). `nlgeul token`이 이 패키지로 찍는다 |
| [parse/](parse/README.md) | 문법 표(조사, 어미, 닫힌 집합)를 읽어 내놓고, 어절 끝으로 문장성분을 판정한다. `token`이 문법 표를 여기서 읽는다 |
| [rule/](rule/README.md) | 여러 그래프가 함께 쓰는 toulmin 규칙 함수와 그 함수들이 읽는 사실 타입 |

## 옮긴 것

다른 모듈이 토큰 층을 하위 프로세스(`nlgeul token --format json`)로 부르지 않고 import해서 함수로 쓰도록 옮겼다. Go는 다른 모듈이 `internal/` 아래 패키지를 import하지 못하게 한다.

| 전 | 후 | 까닭 |
|---|---|---|
| `internal/token` | `pkg/token` | 밖에서 쓰는 것(낱말 표 읽기, `Reader`, 스트림, 갈래·마디·후보·조각 구조, JSON 구조)이 이 패키지에 있다 |
| `internal/parse` | `pkg/parse` | `token`이 문법 표를 여기서 읽는다(`Marker`, `MarkerOf`, `FormOf`, `ParticleMarkers`, `ClosingMarkers`, `ClosedWords`, `AfterNames`, `CaseNames`). `pkg/`는 `internal/`에 기대지 않으므로 함께 옮겼다 |

두 패키지는 통째로 옮겼고 패키지·타입·함수 이름과 동작은 그대로다. 밖에 쓸 일이 없는 `internal/cli`(명령)와 `internal/combo`(조합 나열, 명령 없음)는 `internal/`에 남았고 `pkg/token`을 쓴다. 새로 더한 공개 이름은 `Stream.NextHead`, `Stream.AgainHead`, `Part.Known`, `Branch.Lambda` 넷이다. 공개 API는 [token/README.md](token/README.md)와 [parse/README.md](parse/README.md)의 "공개 API" 절에 있다.

## 밖에서 쓰는 예

모듈은 `github.com/geul-org/nlgeul`이다. 낱말 표는 한 번 읽고, 문장마다 스트림을 연다.

```go
import "github.com/geul-org/nlgeul/pkg/token"

r, err := token.New(token.Options{Words: "words.tsv"}) // 낱말 표는 한 번 읽는다
if err != nil {
	return err
}
st := r.Stream("고양이가 방돌에")
for seg, ok := st.NextHead(2); ok; seg, ok = st.NextHead(2) { // 구간마다 1등 갈래, 마디마다 후보는 λ 차례 앞 2개
	for _, n := range seg.Nodes { // 마디(어절)
		best := n.Contest.Branches[0]       // λ 1위 후보. 후보는 모두 낱말 하나
		_ = best.Words[0].Parts[0].Known()  // 줄기가 새 낱말이면 false
		_ = best.Lambda()                   // 선호 + Lambda·log10(성립)
	}
	next, ok := st.AgainHead(seg.Region, 2) // 같은 구간의 다음 갈래. 없으면 ok가 false
	_, _ = next, ok
}
```

돌아가는 예는 `token/example_test.go`(`Example`)다: 시험용 낱말 표로 문장 둘을 읽어 1등 갈래를 받고, 1등 갈래에 새 낱말이 낀 구간에서만 아는 말로만 끊은 갈래가 나올 때까지 `AgainHead`로 다음 갈래를 받는다.
