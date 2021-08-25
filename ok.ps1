#!/usr/bin/env pwsh
# 30000になる
.\goflock-ex2.exe init
$job1 = & { .\goflock-ex2.exe flockinc } &
$job2 = & { .\goflock-ex2.exe flockinc } &
$job3 = & { .\goflock-ex2.exe flockinc } &
Get-Job | Wait-Job | Out-Null
Receive-Job -Name $job1.Name
Receive-Job -Name $job2.Name
Receive-Job -Name $job3.Name
