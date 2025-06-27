package utils

import "testing"

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123456", true},
		{"abc", false},
		{"", false},
		{"mypassword", true},
	}

	for _, test := range tests {
		result := IsValidPassword(test.input)
		if result != test.expected {
			t.Errorf("For input %s,expected %v, got %v", test.input, test.expected, result)
		}
	}
}
