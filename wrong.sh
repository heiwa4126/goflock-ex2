#!/bin/sh
# 死ぬ。または30000にならない
BIN=./goflock-ex2
INIT="$BIN init"
INC="$BIN inc"

$INIT ; $INC & $INC & $INC
