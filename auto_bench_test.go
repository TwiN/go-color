package color

import (
	"fmt"
	"testing"
)

func BenchmarkAutof(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Autof("hello, %s. HRU %s? I'm %02d%% cute.", "world", "today", 50)
	}
	b.ReportAllocs()
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello, %s. HRU %s? I'm %02d%% cute.", "world", "today", 50)
	}
	b.ReportAllocs()
}
