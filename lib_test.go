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

// test getBitString
func TestGetBitString(t *testing.T) {
	tests := []struct {
		input  byte
		output string
	}{
		{0, "00000000"},
		{1, "00000001"},
		{2, "00000010"},
		{3, "00000011"},
		{4, "00000100"},
		{5, "00000101"},
		{6, "00000110"},
		{7, "00000111"},
		{8, "00001000"},
		{9, "00001001"},
		{10, "00001010"},
		{11, "00001011"},
		{12, "00001100"},
		{13, "00001101"},
		{14, "00001110"},
		{15, "00001111"},
		{16, "00010000"},
		{17, "00010001"},
		{18, "00010010"},
		{19, "00010011"},
		{20, "00010100"},
		{21, "00010101"},
		{22, "00010110"},
		{23, "00010111"},
		{24, "00011000"},
		{25, "00011001"},
		{26, "00011010"},
		{27, "00011011"},
		{28, "00011100"},
		{29, "00011101"},
		{30, "00011110"},
		{31, "00011111"},
		{32, "00100000"},
		{33, "00100001"},
		{34, "00100010"},
		{35, "00100011"},
		{36, "00100100"},
		{37, "00100101"},
		{38, "00100110"},
		{39, "00100111"},
		{40, "00101000"},
		{41, "00101001"},
		{42, "00101010"},
		{43, "00101011"},
		{44, "00101100"},
		{45, "00101101"},
		{46, "00101110"},
		{47, "00101111"},
		{48, "00110000"},
		{49, "00110001"},
		{50, "00110010"},
		{51, "00110011"},
		{52, "00110100"},
		{53, "00110101"},
		{54, "00110110"},
		{55, "00110111"},
		{56, "00111000"},
		{57, "00111001"},
		{58, "00111010"},
		{59, "00111011"},
		{60, "00111100"},
		{61, "00111101"},
		{62, "00111110"},
		{63, "00111111"},

		{64, "01000000"},
	}

	for _, test := range tests {
		result := getBitString(test.input)
		if result != test.output {
			t.Errorf("Input: %d, Expected value: %s, Got: %s", test.input, test.output, result)
		}
	}
}
