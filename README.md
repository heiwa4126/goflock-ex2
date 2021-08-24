# goflock-ex1

GoLangで複数プロセスで排他制御を行うサンプルコード。

[heiwa4126/goflock-ex1: GoLangで複数プロセスでflockを使って排他制御を行うサンプルコード](https://github.com/heiwa4126/goflock-ex1)
が元で、
[gofrs/flock: Thread\-safe file locking library in Go \(originally github\.com/theckman/go\-flock\)](https://github.com/gofrs/flock)
を使って
WindowsとLinuxで動くバージョンを作ろうとしたら、
収取がつかなくなったので作り直したもの。

# 動かし方

まずビルド
```sh
go build
```

Goなのでサブコマンド式になってます。

```sh
BIN=./goflock-ex2
INIT="$BIN init"      # カウンターを0に初期化
INC="$BIN inc"        # カウンターを10000増やす(排他制御なし)。並行で複数動かすと死ぬ。
FINC="$BIN flockinc"  # カウンターを10000増やす。並行で動かしても死なない。

# 期待通りに動いて、10000になる
$INIT ; $INC

# 死ぬ。または30000にならない
$INIT ; $INC & $INC & $INC

# 30000になる
$INIT ; $FINC & $FINC & $FINC
```

# ex1,2 ...

[heiwa4126/goflock-ex1](https://github.com/heiwa4126/goflock-ex1)からの続き。

- ex5 - ex2を [gofrs/flock](https://github.com/gofrs/flock) にしてみたもの。ロックファイルがカウンタファイルと別


# そのほかメモ

Linuxは
終わったら
```sh
rm /tmp/goflock-ex1-count*
```
すること。
