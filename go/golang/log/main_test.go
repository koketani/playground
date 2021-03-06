package main

import (
	// "io"
	"log"
	"os"
	"testing"
)

// package log
//
// const (
// 	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
// 	Ltime                         // the time in the local time zone: 01:23:23
// 	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
// 	Llongfile                     // full file name and line number: /a/b/c/d.go:23
// 	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
// 	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
// 	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
// 	LstdFlags     = Ldate | Ltime // initial values for the standard logger
// )

func Test_flags_default(t *testing.T) {
	if initFlag := log.Flags(); initFlag != (log.Ldate | log.Ltime) {
		t.Errorf("expected 3, but %v", initFlag)
	}
}

func Example_flags_customize() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Lmsgprefix)
	log.Println("hello")

	// Output:
	// main_test.go:32: hello
}

func Example_flags_zero() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.Println("hello")

	// Output:
	// hello
}
