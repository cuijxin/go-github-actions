package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; Hello GitHub Actions", result)
	}
}