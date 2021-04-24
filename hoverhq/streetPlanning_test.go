package hoverhq

import "testing"

func TestStreetPlan(t *testing.T) {
	tests := [][]int{{1, 4}, {3, 25}, {4, 64}}
	for _, test := range tests {
		actual := streetPlan(test[0])
		if test[1] != actual {
			t.Errorf("expected %d got %d", actual, test[1])
		}
	}
}

func TestStreePlanFib(t *testing.T) {
	tests := [][]int{{1, 4}, {3, 25}, {4, 64}}
	for _, test := range tests {
		actual := streetPlanFib(test[0])
		if test[1] != actual {
			t.Errorf("expected %d got %d", actual, test[1])
		}
	}
}

func BenchmarkStreetPlanFib(b *testing.B) {
	//9.11 ns/op
	for i := 0; i < b.N; i++ {
		streetPlanFib(10)
	}
}

func BenchmarkStreetPlanFibRec(b *testing.B) {
	//279 ns/op
	for i := 0; i < b.N; i++ {
		streePlanFibRec(10)
	}
}

func Benchmark(b *testing.B) {
	//2049 ns/op
	for i := 0; i < b.N; i++ {
		fibGoRoutine(10)
	}
}
