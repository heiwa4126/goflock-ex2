//go:build linux || darwin || aix

package ex5

// const count_file = "/tmp/goflock-ex2-count5"
// const lock_file = "/tmp/goflock-ex2-count5.lock"
var count_file string = "/tmp/goflock-ex2-count5"
var lock_file string = "/tmp/goflock-ex2-count5.lock"
