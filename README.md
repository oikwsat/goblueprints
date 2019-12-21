# Go言語によるWebアプリケーション開発

## 1章 WebSocketを使ったチャットアプリケーション

1. test
```
cd chapter01/trace
go test -cover
```

2. build
```
cd chapter01/chat
go get github.com/stretchr/gomniauth/
go build -o chat
./chat
```

3. access to http://localhost:8080/