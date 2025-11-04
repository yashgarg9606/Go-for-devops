package calc

import "testing"

// BenchmarkCalculateSum benchmarks the CalculateSum function
func BenchmarkCalculateSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSum(100)
	}
}

//Run benchmark using command: "go test -bench="."
//  day_05/unit_testing/benchmarking/..."
