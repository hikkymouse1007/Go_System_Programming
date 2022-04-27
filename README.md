# Go_System_Programming

https://www.lambdanote.com/products/go

- 解答つきの連載について
  - https://ascii.jp/serialarticles/1235262/
- Zenn のメモ
  - https://zenn.dev/mohira/scraps/c0aca378ac9fa7
  - https://zenn.dev/mohira/scraps/c0aca378ac9fa7

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
