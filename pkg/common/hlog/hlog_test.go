package hlog

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	defaultLog := DefaultLogger()
	systemLog := SystemLogger()

	assert.DeepEqual(t, logger, defaultLog)
	assert.DeepEqual(t, sysLogger, systemLog)
	assert.NotEqual(t, logger, systemLog)
	assert.NotEqual(t, sysLogger, defaultLog)
}

func TestSetLogger(t *testing.T) {
	setLog := &defaultLogger{
		stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		depth:  6,
	}
	assert.NotEqual(t, logger, setLog)
	SetLogger(setLog)
	assert.DeepEqual(t, logger, setLog)
}
