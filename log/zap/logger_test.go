package zap

import (
	"log"
	"testing"
)

func TestInfof(t *testing.T) {
	Info("info")
	Error("error")
	Warn("warn")
	Debug("debug")

	Error("error")

	Error("error")

	Error("error")
	Error("error")
}

type tmp int

func (t *tmp) add() {
	log.Printf("hahahahhahah")
}

var tp *tmp

func Test_pointer(t *testing.T) {
	tp.add()
}
