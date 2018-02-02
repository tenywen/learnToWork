package logger

import (
	"runtime"
)

func PANIC(v ...interface{}) {
	if x := recover(); x != nil {
		sugarLogger.DPanic(v...)
		for i := 1; i < 10; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if ok {
				f := runtime.FuncForPC(pc)
				sugarLogger.DPanicf("%v %v:%v", f.Name(), file, line)
			}
		}
	}
}
