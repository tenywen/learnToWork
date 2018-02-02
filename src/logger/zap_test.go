package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	StartLogger("gs.log")
	DEBUG("这是一个测试!!!")
}

func TestError(t *testing.T) {
	StartLogger("test.log")
	ERROR("这是一个测试!!!")
}

func TestInfo(t *testing.T) {
	StartLogger("test.log")
	INFO("这是一个测试!!!")
}

func TestWarn(t *testing.T) {
	StartLogger("test.log")
	WARN("这是一个测试!!!")
}
