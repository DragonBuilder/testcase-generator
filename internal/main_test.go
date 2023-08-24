package internal

import (
	"os"
	"testcase-generator/internal/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.Init()
	exitVal := m.Run()
	os.Exit(exitVal)
}
