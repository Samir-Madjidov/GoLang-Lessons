package calculate

import "testing"

func TestAdd(t *testing.T) {
	reult := add(2, 3)
	expected := 5

	if reult != expected {
		t.Errorf("add(2,3) was incorrect, got: %d, want: %d.", reult, expected)
	}

}
func TestMultiplay(t *testing.T) {
	result := Multiplay(3, 4)
	expected := 20

	if result != expected {
		t.Errorf("Multiplay(3,4) was incorrect, got: %d, want: %d.", result, expected)
	}
}
