package helper

import (
	"runtime"
)

type Call struct {
	Skip int
	File string
	Line int
}

// Caller call link
func Caller() (list []*Call) {
	var file string
	var line int
	var ok bool
	for i := 0; ; i++ {
		_, file, line, ok = runtime.Caller(i)
		if !ok {
			break
		}
		list = append(list, &Call{
			Skip: i,
			File: file,
			Line: line,
		})
	}
	return
}
