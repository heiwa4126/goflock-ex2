package ex5

import "os"

// const count_file = "/tmp/goflock-ex2-count5"
// const lock_file = "/tmp/goflock-ex2-count5.lock"
var count_file string = os.TmpDir() + "\\goflock-ex2-count5"
var lock_file string = os.TmpDir() + "\\goflock-ex2-count5.lock"
