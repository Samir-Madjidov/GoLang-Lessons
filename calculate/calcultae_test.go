package calculate

import "testing"

func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -1, -2},
		{"mixed numbers", -5, 10, 5},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
