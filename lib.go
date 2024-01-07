package base64

import (
	"strconv"
	"strings"
)

type B64 struct {
	value string
}

var base64Chars = [64]byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
	'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3',
	'4', '5', '6', '7', '8', '9', '+', '/',
}

var binaryToInt = map[string]int{
	"000000": 0, "000001": 1, "000010": 2, "000011": 3,
	"000100": 4, "000101": 5, "000110": 6, "000111": 7,
	"001000": 8, "001001": 9, "001010": 10, "001011": 11,
	"001100": 12, "001101": 13, "001110": 14, "001111": 15,
	"010000": 16, "010001": 17, "010010": 18, "010011": 19,
	"010100": 20, "010101": 21, "010110": 22, "010111": 23,
	"011000": 24, "011001": 25, "011010": 26, "011011": 27,
	"011100": 28, "011101": 29, "011110": 30, "011111": 31,
	"100000": 32, "100001": 33, "100010": 34, "100011": 35,
	"100100": 36, "100101": 37, "100110": 38, "100111": 39,
	"101000": 40, "101001": 41, "101010": 42, "101011": 43,
	"101100": 44, "101101": 45, "101110": 46, "101111": 47,
	"110000": 48, "110001": 49, "110010": 50, "110011": 51,
	"110100": 52, "110101": 53, "110110": 54, "110111": 55,
	"111000": 56, "111001": 57, "111010": 58, "111011": 59,
	"111100": 60, "111101": 61, "111110": 62, "111111": 63,
}

func Encode(s string) B64 {
	out := getStringToBitString(s)
	sixes := getSixBitStrings(out)
	outStr := getB64StrFromSixes(sixes)
	outStr = addPadding(outStr)
	return B64{value: outStr}
}

func getB64StrFromSixes(sixes []string) string {
	outStr := ""
	for _, h := range sixes {
		value := binaryToInt[h]
		char := base64Chars[value]
		outStr += string(char)
	}
	return outStr
}

func getStringToBitString(s string) string {
	byteArr := []byte(s)
	var bits []bool
	for _, word := range byteArr {
		for i := 7; i >= 0; i-- {
			if (word >> uint(i) & 1) == 1 {
				bits = append(bits, true)
			} else {
				bits = append(bits, false)
			}
		}
	}
	var builder strings.Builder
	for _, bit := range bits {
		if bit {
			builder.WriteString("1")
		} else {
			builder.WriteString("0")
		}
	}
	out := builder.String()
	return out
}

func getSixBitStrings(out string) []string {
	var sixes []string
	h := ""
	for i := 0; i < len(out); i++ {
		h = h + out[i:i+1]
		if len(h) == 6 {
			sixes = append(sixes, h)
			h = ""
		}

	}
	if len(h) != 0 {
		for i := 0; i <= 6-len(h); i++ {
			h += "0"
		}
		sixes = append(sixes, h)
	}
	return sixes
}

func addPadding(outStr string) string {
	if len(outStr)%4 == 3 {
		outStr += "="
	} else if len(outStr)%4 == 2 {
		outStr += "=="
	}
	return outStr
}

func getBitString(b byte) string {
	value := ""
	for i := 7; i >= 0; i-- {
		bit := int((b >> uint(i)) & 1)
		value += strconv.Itoa(bit)
	}
	return value
}

func (b *B64) Decode() string {
	return b.value
}

func (b *B64) String() string {
	return b.value
}
