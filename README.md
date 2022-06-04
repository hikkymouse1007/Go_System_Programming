# Go_System_Programming

https://www.lambdanote.com/products/go

- 解答つきの連載について
  - https://ascii.jp/serialarticles/1235262/
- Zenn のメモ
  - https://zenn.dev/mohira/scraps/c0aca378ac9fa7
  - https://zenn.dev/mohira/scraps/ea040641f2f122
- その他良さげなもの
  Go から学ぶ I/O
  - https://zenn.dev/hsaki/books/golang-io-package/viewer/file
    著者のその他の著書
  - https://zenn.dev/hsaki?tab=books
    Go 言語(golang) ファイルの読み書き、作成、存在確認、一行ずつ処理、コピー など
  - https://golang.hateblo.jp/entry/2018/11/09/163000

サンプルコード

- https://github.com/mikutas/go-system-programming

## Go by Example

https://gobyexample.com/

## デバッガのセットアップ

https://qiita.com/gs1068/items/1d89d4a9bed070782298

デバッガを起動するためには、go mod でモジュールにリモートリポジトリを指定する
必要がある。
https://stackoverflow.com/questions/67306638/go-test-results-in-go-cannot-find-main-module-but-found-git-config-in-users

```
git init
git remote add origin https://github.com/hikkymouse1007/Go_System_Programming.git
go mod init github.com/hikkymouse1007/Go_System_Programming
```

https://qiita.com/ymmy02/items/b4b38a57c024510c5210

### デバッグしたいファイルについて

ルートディレクトリにデバッグ対象のファイルを置いてファイルまたはデバッガを実行する

### Keyword

- システムコール
- カーネル
  カーネルは、シェルやデーモンといった、
  その他のソフトウェアを手足のように使うことによって、はじめて組織が回ります。
  https://wa3.i-3-i.info/word15.html
- ファイルディスクリプタ
  https://wa3.i-3-i.info/word14383.html
- プロセス起動
- 副作用
  https://python.ms/side-effect/#_1-%E5%89%AF%E4%BD%9C%E7%94%A8%E3%81%A8%E3%81%AF
  => オブジェクトの破壊的変更

```python
#
# 副作用あり
#
lst1 = [1, 0, 3, 2]
lst1.sort()
lst1
# [0, 1, 2, 3] -> オブジェクトが変化したので副作用がある
```

```python
#
# 副作用なし
#
lst2 = [1, 0, 3, 2]
lst3 = sorted(lst2)

lst2
# [1, 0, 3, 2] -> オブジェクトは変化していないので副作用はない

lst3
# [0, 1, 2, 3]
```

- バッファ

  > 何か（主にデータ）を一時的にプールしておく場所
  > https://wa3.i-3-i.info/word11609.html

- Go のインターフェース
  > インタフェースは、構造体と違って何かしら実体を持つものを表すのでは なく、「どんなことができるか」を宣言しているだけです

### main 関数について

https://jp.quora.com/Go%E8%A8%80%E8%AA%9E%E3%81%A71%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%A7%E5%AE%8C%E7%B5%90%E3%81%99%E3%82%8B%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92%E8%A4%87%E6%95%B0%E4%BD%9C%E6%88%90%E3%81%97%E3%81%9F
ファイル名を指定して実行する場合には、`package main`であれば同階層ディレクトリの複数ファイルで main 関数を使っても
実行できる。

### godoc のインストールについて

https://zenn.dev/mkosakana/articles/bb411e9d6b5ad9

### 練習問題　解答参考

https://github.com/yuyabu/go-sys

### Q 2-3 　補足

https://stackoverflow.com/questions/31622052/how-to-serve-up-a-json-response-using-go

```go
// You can set your content-type header so clients know to expect json

w.Header().Set("Content-Type", "application/json")

// Another way to marshal a struct to json is to build an encoder using the http.ResponseWriter
// get a payload p := Payload{d}
json.NewEncoder(w).Encode(p)
```

### 3.4.1 補足

バッファサイズ 5 ずつ表示する

```txt
$ go run stdin.go < stdin.go
size=5 input='packa
size=5 input='ge ma
size=5 input='in

i
size=5 input='mport
size=5 input=' (
        "
size=5 input='fmt"

size=5 input='  "io"
size=5 input='
        "os
size=5 input='"
)


size=5 input='func
size=5 input='main(
size=5 input=') {

size=5 input='for {
size=5 input='
                bu
size=5 input='ffer
size=5 input=':= ma
size=5 input='ke([]
size=5 input='byte,
size=5 input=' 5)

size=5 input='  size
size=5 input=', err
size=5 input=' := o
size=5 input='s.Std
size=5 input='in.Re
size=5 input='ad(bu
size=5 input='ffer)
size=5 input='
                if
size=5 input=' err
size=5 input='== io
size=5 input='.EOF
size=5 input='{

size=5 input='fmt.P
size=5 input='rintl
size=5 input='n("EO
size=5 input='F")

size=5 input='          bre
size=5 input='ak

size=5 input='}
                f
size=5 input='mt.Pr
size=5 input='intf(
size=5 input='"size
size=5 input='=%d i
size=5 input='nput=
size=5 input=''%s\n
size=5 input='", si
size=5 input='ze, s
size=5 input='tring
size=5 input='(buff
size=5 input='er))

size=5 input='          //
size=5 input='input
size=5 input=':
                /
size=5 input='/ hel
size=5 input='lo wo
size=5 input='rld

size=5 input='  // s
size=5 input='ize=5
size=5 input=' inpu
size=5 input='t='he
size=5 input='llo

size=5 input='  // s
size=5 input='ize=5
size=5 input=' inpu
size=5 input='t=' w
size=5 input='orl

size=5 input='  // s
size=5 input='ize=2
size=5 input=' inpu
size=5 input='t='d

size=5 input='  }
}

EOF
}
```

## Go の io パッケージ解説

https://christina04.hatenablog.com/entry/golang-io-package-diagrams

> io.Copy は io.Reader からデータを読み込み、io.Writer へと書き出します。
> 入力と出力をそのままセットで実行する

```go
func Copy(dst Writer, src Reader) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}
```

### エンディアンとは

https://qiita.com/tobira-code/items/a03f39a02678d80bbd26

https://wa3.i-3-i.info/word11426.html
http://tanehp.ec-net.jp/heppoko-lab/prog/zakki/byteorder/byteorder.html

https://rainbow-engine.com/little-endian-big-endian/

> TCP/IP プロトコルで使用されている事から、インターネット上でデータを転送する際には「ビッグエンディアン」に変換する必要があるという意味になります。
> 逆にリトルエンディアンは、前述の計算処理における利点がある事から Intel 等のプロセッサで利用されています

## png のチャンク

![スクリーンショット 2022-05-02 0 56 21](https://user-images.githubusercontent.com/54907440/166153977-daca9d14-6516-49a8-9e00-0791964f4c3b.png)
![スクリーンショット 2022-05-02 0 56 28](https://user-images.githubusercontent.com/54907440/166153978-06e34678-2dd6-4856-9fe6-9a344f330013.png)

## tee コマンド補足

https://atmarkit.itmedia.co.jp/ait/articles/1611/16/news022.html

```
ind /bin /usr/bin -type f | tee cmdlist | less

（findコマンドの結果をファイル「cmdlist」に保存し、さらにlessコマンドで閲覧する）
```

## Go routine & channel

並行実行と並列実行の違い
https://zenn.dev/hsaki/books/golang-concurrency/viewer/term
![スクリーンショット 2022-05-21 15 00 54](https://user-images.githubusercontent.com/54907440/169638147-b00bf9b3-f1a7-4591-b476-f20d57d35a0c.png)

goroutine とチャネル
https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

チャネルは、並行実行される goroutine 間を接続するパイプ(トンネル)のイメージ。
つまり、並行実行している関数から値を受信する。(ある goroutine から別の goroutine へ値を渡す。)
![スクリーンショット 2022-05-21 15 00 54](https://user-images.githubusercontent.com/54907440/169638377-2e23edd7-32d7-4eeb-8571-fb634f31fff8.png)

> バッファが詰まるとチャネルへの送信をブロックする。バッファが空のときは、チャネルの受信をブロックする。
