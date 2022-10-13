package hlog

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"log"
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	logger := &defaultLogger{
		stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		depth:  4,
	}

	logger.SetLevel(LevelTrace)
	assert.DeepEqual(t, LevelTrace, logger.level)
	assert.DeepEqual(t, LevelTrace.toString(), logger.level.toString())
	logger.SetLevel(LevelDebug)
	assert.DeepEqual(t, LevelDebug, logger.level)
	assert.DeepEqual(t, LevelDebug.toString(), logger.level.toString())
	logger.SetLevel(LevelInfo)
	assert.DeepEqual(t, LevelInfo, logger.level)
	assert.DeepEqual(t, LevelInfo.toString(), logger.level.toString())
	logger.SetLevel(LevelNotice)
	assert.DeepEqual(t, LevelNotice, logger.level)
	assert.DeepEqual(t, LevelNotice.toString(), logger.level.toString())
	logger.SetLevel(LevelWarn)
	assert.DeepEqual(t, LevelWarn, logger.level)
	assert.DeepEqual(t, LevelWarn.toString(), logger.level.toString())
	logger.SetLevel(LevelError)
	assert.DeepEqual(t, LevelError, logger.level)
	assert.DeepEqual(t, LevelError.toString(), logger.level.toString())
	logger.SetLevel(LevelFatal)
	assert.DeepEqual(t, LevelFatal, logger.level)
	assert.DeepEqual(t, LevelFatal.toString(), logger.level.toString())
	logger.SetLevel(7)
	assert.DeepEqual(t, 7, int(logger.level))
	assert.DeepEqual(t, "[?7] ", logger.level.toString())
}
