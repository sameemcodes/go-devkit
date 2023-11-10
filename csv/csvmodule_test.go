package csvmodule_test

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	// setup
	code := m.Run()
	// teardown
	os.Exit(code)
}

func TestWriteCSV(t *testing.T) {
	// your test code here
}
