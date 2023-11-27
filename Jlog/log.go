package Jlog

import (
	"path"
	"runtime"
	"strings"
)

type CallInfo struct {
	PackageName string
	FileName    string
	FuncName    string
	Line        int
}

func CallerInfo(skip ...int) (caller *CallInfo) {
	caller = &CallInfo{}
	skipCount := 1
	if len(skip) > 0 {
		skipCount = skip[0]
	}

	pc, file, line, ok := runtime.Caller(skipCount)
	if !ok {
		return
	}
	caller.Line = line
	_, caller.FileName = path.Split(file)

	parts := strings.Split(runtime.FuncForPC(pc).Name(), `.`)
	pl := len(parts)
	caller.FuncName = parts[pl-1]

	if parts[pl-2][0] == '(' {
		caller.FuncName = parts[pl-2] + `.` + caller.FuncName
		caller.PackageName = strings.Join(parts[0:pl-2], `.`)
	} else {
		caller.PackageName = strings.Join(parts[0:pl-1], `.`)
	}

	return
}
