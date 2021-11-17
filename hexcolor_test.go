// The MIT License (MIT) Copyright (c) 2019-2021 artipie.com
// https://github.com/g4s8/hexcolor/blob/master/LICENSE

package hexcolor

import (
	"errors"
	"fmt"
	"image/color"
	"testing"
)

func TestParseHex(t *testing.T) {
	cases := []struct {
		hex string
		rgb [4]uint8
	}{
		{"#00000000", [...]uint8{0, 0, 0, 0}},
		{"#000000ff", [...]uint8{0, 0, 0, 0xFF}},
		{"#fa0000d7", [...]uint8{0xFA, 0, 0, 0xD7}},
		{"#f0fffff0", [...]uint8{240, 255, 255, 240}},
		{"#f0f8ff", [...]uint8{0xF0, 0xF8, 0xFF, 0xFF}},
		{"#faebd7", [...]uint8{0xFA, 0xEB, 0xD7, 0xFF}},
		{"#00ffff", [...]uint8{0, 255, 255, 255}},
		{"#7fffd4", [...]uint8{127, 255, 212, 255}},
		{"#f0ffff", [...]uint8{240, 255, 255, 255}},
		{"#f5f5dc", [...]uint8{245, 245, 220, 255}},
		{"#ffe4c4", [...]uint8{255, 228, 196, 255}},
		{"#000000", [...]uint8{0, 0, 0, 255}},
		{"#abcd", [...]uint8{0xAA, 0xBB, 0xCC, 0xDD}},
		{"#abc", [...]uint8{0xAA, 0xBB, 0xCC, 0xFF}},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("parseHex(%s)", c.hex), checkHex(c.hex, c.rgb))
	}
}

func TestInvalidHexWithoutInitSymbol(t *testing.T) {
	_, err := Parse("abc")
	checkErr(t, err)
}

func TestInvalidHexSymbols(t *testing.T) {
	_, err := Parse("#abx")
	checkErr(t, err)
}

func TestInvalidHexCount(t *testing.T) {
	_, err := Parse("#abcdef123")
	checkErr(t, err)
}

func checkErr(t *testing.T, err error) {
	if err == nil {
		t.Errorf("expected invalid error but got nil")
	}
	if !errors.Is(ErrInvalidFormat, err) {
		t.Errorf("expected invalid error but got %s", err)
	}
}

func checkHex(hex string, expected [4]uint8) func(*testing.T) {
	return func(t *testing.T) {
		c, err := Parse(hex)
		if err != nil {
			t.Fatalf("failed to parse: %s", err)
		}
		exp := color.RGBA{expected[0], expected[1], expected[2], expected[3]}
		checkColor(t, exp, c)
	}
}

func checkColor(t *testing.T, expected, actual color.Color) {
	er, eg, eb, ea := expected.RGBA()
	ar, ag, ab, aa := actual.RGBA()
	if er != ar || eg != ag || eb != ab || ea != aa {
		t.Fatalf("expected %4v but was %4v", expected, actual)
	}
}

func benchmarkParse(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse(s)
	}
}

func BenchmarkParseLongTransparent(b *testing.B) {
	benchmarkParse("#7fffd4fc", b)
}

func BenchmarkParseLong(b *testing.B) {
	benchmarkParse("#7fffd4", b)
}

func BenchmarkParseShortTransparent(b *testing.B) {
	benchmarkParse("#1234", b)
}

func BenchmarkParseShort(b *testing.B) {
	benchmarkParse("#123", b)
}

func BenchmarkParseErr(b *testing.B) {
	benchmarkParse("foobar", b)
}
