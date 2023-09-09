package reverse_int

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
		{"569", 569, 965},
		{"222333222", 222333222, 222333222},
		{"212", 212, 212},
		{"313", 313, 313},
		{"22237222", 22237222, 22273222},
	}
	return tests
}

func TestReverseInt(t *testing.T) {
	tests := getTestData()
	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			result := ReverseInt(data.n)
			if result != data.expected {
				t.Errorf("The result is wrong expected : %d returned : %d ", data.expected, result)
			}
		})
	}
}

func BenchmarkReverseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseInt(333313333)
	}
}
