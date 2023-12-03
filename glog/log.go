package glog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// KeyExecutionTime and KeyRepoLayer are predefined keys for log context.
const KeyExecutionTime = "execution_time"
const KeyRepoLayer = "repo_layer"

// Msg struct is a representation of a log message structure to be used in logging
type Msg struct {
	Level  string `json:"level"`
	Msg    string `json:"msg"`
	Err    any    `json:"err"`
	Caller string `json:"caller"`
}

// logMessage struct is an internal representation of a log message.
type logMessage struct {
	Msg  *Msg
	Data any

	key       string
	startTime time.Time
}

// NewLogMsg function creates and return a new logMessage object
// with time of object creation
func NewLogMsg(key string) logMessage {
	return logMessage{
		key:       key,
		startTime: time.Now(),
	}
}

// CallInfo struct represents information about method/function calls.
type CallInfo struct {
	PackageName string
	FileName    string
	FuncName    string
	Line        int
}

// CallerInfo function retrieves information about calling function,
// it can skip levels of call stack while collecting information.
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

// CallerInfoStr function retrieves caller information as string format
func CallerInfoStr() string {
	caller := CallerInfo(2)
	return caller.PackageName + " " + caller.FileName + " " +
		caller.FuncName + ":" + fmt.Sprintf("%d", caller.Line)
}
