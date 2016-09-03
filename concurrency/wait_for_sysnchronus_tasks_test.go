package concurrency

import "testing"

func BenchmarkRunSynchronousTasks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunSynchronousTasks()
	}
}
