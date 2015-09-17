package godis

import "testing"

// Test SDS strings and various methods in it
func TestSDS(t *testing.T) {
	str := "test string"
	s := NewSDS(str)
	if s.value != str {
		t.Errorf("SDS value: %s, want %s", s.value, str)
	}

	if s.String() != str {
		t.Errorf("SDS.String(): %s, want %s", s.String(), str)
	}

	if s.get() != str {
		t.Errorf("SDS.get(): %s, want %s", s.get(), str)
	}
}
