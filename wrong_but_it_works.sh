#!/bin/sh
# 並行に動かさなければ、ちゃんと10000になる。
BIN=./goflock-ex2
INIT="$BIN init"
INC="$BIN inc"

$INIT ; $INC
