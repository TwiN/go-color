package color

import "testing"

func BenchmarkInBlackOverGreen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InBlackOverGreen("hello, world")
	}
	b.ReportAllocs()
}
