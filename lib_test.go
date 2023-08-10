package base64

import (
	"testing"
)

func TestB64FromStr(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"hello", "aGVsbG8="},
		{"world", "d29ybGQ="},
		{"123", "MTIz"},
		{"", ""},
	}

	for _, test := range tests {
		b64Instance := Encode(test.input)
		if b64Instance.String() != test.output {
			t.Errorf("Input: %s, Expected value: %s, Got: %s", test.input, test.output, b64Instance.String())
		}
	}
}

func TestB64ToStr(t *testing.T) {
	input := "world"
	b64Instance := B64{value: input}
	result := b64Instance.Decode()
	if result != input {
		t.Errorf("Expected value: %s, got: %s", input, result)
	}
}

func TestGetValue(t *testing.T) {
	input := "testing"
	b64Instance := B64{value: input}
	result := b64Instance.String()
	if result != input {
		t.Errorf("Expected value: %s, got: %s", input, result)
	}
}
