# Hello Cafe

## 사용언어 및 라이브러리
* Golang:Gin

## 빌드
```shell
  make -f build/Makefile build
```

## 실행
```shell
  make -f build/Makefile run
```

## 아키텍쳐
### 관심사 분리
* `handler` / `service` / `repository` layer 로 관심사를 구분하여 단방향으로 의존 하도록 작성
* 각 Layer 별로 독립적으로 개발이 가능 하기 때문에 기능 확장 변경시 대응에 유리하다
  * 코드의 재사용성 및 유지보수성을 향상
### Handler
* client 와 직접 연결되어 http request 를 요청 받고 client 에 response 해주는 부분을 담당
### Service
* Handler 를 통해 전달받은 Parameter 로 비지니스 로직을 구현 하는 부분
### Repository
* DB, Redis, ElasticSearch 와 상호 작용 하며 데이터의 조회, 생성, 수정, 삭제 를 담당한다

## Restful API
* URI 를 통해 요청할 자원에 대해 표시하며, Http Method 를 통해 자원의 CRUD 를 컨트롤 할 수 있도록 처리
