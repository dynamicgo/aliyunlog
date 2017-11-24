package aliyunlog

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dynamicgo/config"
	"github.com/dynamicgo/slf4go"
)

var cnf *config.Config

func init() {
	cnf, _ = config.NewFromFile("./test.json")
}

func BenchmarkPutLog(b *testing.B) {
	factory, err := NewAliyunBackend(cnf)

	assert.NoError(b, err)

	slf4go.Backend(factory)

	logger := slf4go.Get("test")

	logger.Debug("1111")
}
