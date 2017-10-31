[![Build Status](https://travis-ci.org/ynozue/apichecker.svg?branch=master)](https://travis-ci.org/ynozue/apichecker)

# API Checker
API が正常に動いているかをチェックします。

## ビルド方法
```
make build
```

## コマンドでのビルド方法
**Linux 向け**
```
GOOS=linux GOARCH=amd64 go build -o build/apichecker
```
**Mac 向け**
```
GOOS=darwin GOARCH=amd64 go build -o build/apichecker
```
**Windows 向け**
```
GOOS=windows GOARCH=amd64 go build -o build/apichecker
```

## 実行方法
**開発環境**
```
go run apichecker.go -endpoint=${エンドポイントURI} -token=${Line Notify 用 Token}
```
**ビルドしたバイナリの実行**
```
apichecker -endpoint=${エンドポイントURI} -token=${Line Notify 用 Token}
```
