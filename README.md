# wecty tutorial

[nobonobo/wecty](https://github.com/nobonobo/wecty) のチュートリアルをやってみる。

## 環境
* Windows 10 64bit
* go version go1.15 windows/amd64
* github.com/nobonobo/wecty v0.1.0

## 動かし方

* `$ go get github.com/nobonobo/wecty/cmd/wecty` でcliツールを取得
* `$ go generate` でレンダラーの自動生成
* `$ GOOS=js GOARCH=wasm go build -o ./src/main.js .` でwasmファイルを吐き出す
* `src/index.html` を開けばwasmが読み込まれて動く。

wecty/cmdでserverを起動してwasmのbuildとかもやってくれますが、  
wasmを他のプログラムに埋め込みたかったのでwasmを吐き出す方法にしています。
