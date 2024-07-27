package main

import (
	benchmark "informations/test/Benchmark"
	"testing"
)
func main(){

}
func BenchmarkStringTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark.ConcatStringWithoutBuilder("Hello","golang")
	}
}
func BenchmarkStringTests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark.ConcatStringWithBuilder("Hello","Alireza")
	}
}