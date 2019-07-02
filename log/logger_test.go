package logger

import (
	"os"
	"testing"
)

func Test_logger(t *testing.T) {
	os.Stderr.Write([]byte("sdfsdf"))
}
