package profiling

import "testing"

// This is a test function written to generate profiling
func TestProfiling(t *testing.T) {
	for i := 0; i < 1e8; i++ {
		_ = i * i
	}
}
