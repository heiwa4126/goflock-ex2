#!/bin/sh
# 30000になる
BIN=./goflock-ex2
INIT="$BIN init"
INC="$BIN flockinc"

$INIT ; $INC & $INC & $INC
