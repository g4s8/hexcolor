package hexcolor

import (
	"fmt"
	"image/color"
	"testing"
)

// @todo #1:30min Add tests to cover invalid color inputs
//  to check that error is correct.
func TestParseHex(t *testing.T) {
	cases := []struct {
		hex string
		rgb [3]uint8
	}{
		{"#f0f8ff", [...]uint8{0xF0, 0xF8, 0xFF}},
		{"#faebd7", [...]uint8{0xFA, 0xEB, 0xD7}},
		{"#00ffff", [...]uint8{0, 255, 255}},
		{"#7fffd4", [...]uint8{127, 255, 212}},
		{"#f0ffff", [...]uint8{240, 255, 255}},
		{"#f5f5dc", [...]uint8{245, 245, 220}},
		{"#ffe4c4", [...]uint8{255, 228, 196}},
		{"#000000", [...]uint8{0, 0, 0}},
		{"#abc", [...]uint8{0xAA, 0xBB, 0xCC}},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("parseHex(%s)", c.hex), checkHex(c.hex, c.rgb))
	}
}

func checkHex(hex string, expected [3]uint8) func(*testing.T) {
	return func(t *testing.T) {
		c, err := Parse(hex)
		if err != nil {
			t.Fatalf("failed to parse: %s", err)
		}
		exp := color.RGBA{expected[0], expected[1], expected[2], 0xFF}
		checkColor(t, exp, c)
	}
}

func checkColor(t *testing.T, expected, actual color.Color) {
	er, eg, eb, ea := expected.RGBA()
	ar, ag, ab, aa := actual.RGBA()
	if er != ar || eg != ag || eb != ab || ea != aa {
		t.Fatalf("expected %3v but was %3v", expected, actual)
	}
}

func benchmarkParse(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse(s)
	}
}

func BenchmarkParseLong(b *testing.B) {
	benchmarkParse("#7fffd4", b)
}

func BenchmarkParseShort(b *testing.B) {
	benchmarkParse("#123", b)
}

func BenchmarkParseErr(b *testing.B) {
	benchmarkParse("foobar", b)
}
