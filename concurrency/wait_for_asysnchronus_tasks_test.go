package concurrency

import "testing"

func BenchmarkRunAsynchronousTasks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunAsynchronousTasks()
	}
}
