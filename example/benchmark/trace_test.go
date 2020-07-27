package benchmark

import (
	"os"
	"runtime/trace"
	"testing"
)

func TestTrace(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here
}
