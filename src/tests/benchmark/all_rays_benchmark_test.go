package benchmark

import (
	"chess/chess_engine/move_gen"
	"testing"
)

func BenchmarkAllRays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		move_gen.Gen_attack_rays(false)
	}
	b.ReportAllocs()
}

func BenchmarkAllRaysDiag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		move_gen.Gen_attack_rays(true)
	}
	b.ReportAllocs()
}