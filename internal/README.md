# internal

함수 하나에 파일 하나(filefunc)라 패키지 폴더의 파일 수는 20개를 넘는다.

| 패키지 | 내용 |
|---|---|
| [cli/](cli/README.md) | cobra 루트 명령 `nlgeul`과 하위 명령 `token` |
| [combo/](combo/README.md) | (명령 없음) 토큰 층의 후보를 경합마다 하나씩 골라 이은 조합을 세고 정렬해 나열(실험) |

`token`과 `parse`는 밖에서 import하도록 `pkg/`로 옮겼다([pkg/README.md](../pkg/README.md)). 여기 둘은 `pkg/token`을 쓴다.
