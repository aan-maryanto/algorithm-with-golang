package factorial

import "testing"

func getTestData() []struct {
	name     string
	n        int
	expected int
} {
	var tests = []struct {
		name     string
		n        int
		expected int
	}{
		{"10!", 10, 3628800},
		{"5!", 5, 120},
		{"8!", 8, 40320},
		{"13!", 11, 39916800},
		{"7!", 7, 5040},
	}
	return tests
}

func TestFactorial(t *testing.T) {
	tests := getTestData()
	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			result := Factorial(data.n)
			if result != data.expected {
				t.Errorf("The result is wrong expected : %d returned : %d ", data.expected, result)
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(20)
	}
}
