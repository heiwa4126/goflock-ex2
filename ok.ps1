# 30000になる
$bin = ".\goflock-ex2.exe"

# $INIT ; $INC & $INC & $INC
&$bin init
$inc = {&$bin flockinc}
Start-Job -ScriptBlock $inc
Start-Job -ScriptBlock $inc
Start-Job -ScriptBlock $inc
Get-Job | Wait-Job
