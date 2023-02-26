// The MIT License (MIT)
// Copyright (c) 2019-2023 Kirill Che. <g4s8.public@gmail.com>
// https://github.com/g4s8/hexcolor/blob/master/LICENSE

package hexcolor

import (
	"errors"
	"image/color"
)

// ErrInvalidFormat indicates that format of hex color
// string is invalid.
var ErrInvalidFormat = errors.New("invalid hex color format")

// Parse hex string to image/color
// It takes #RGB, #RRGGBB, #ARGB, #AARRGGBB strings as input
// and parse it to RGBA color from `color` Go package.
// In case of invalid hex, it returns error
func Parse(hex string) (c color.RGBA, err error) {
	// got this parsing code from here https://stackoverflow.com/a/54200713/1723695

	c.A = 0xff

	if hex[0] != '#' {
		return c, ErrInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return 10 + b - 'a'
		case b >= 'A' && b <= 'F':
			return 10 + b - 'A'
		}
		err = ErrInvalidFormat
		return 0
	}

	switch len(hex) {
	case 9:
		c.R = (hexToByte(hex[1]) << 4) + hexToByte(hex[2])
		c.G = (hexToByte(hex[3]) << 4) + hexToByte(hex[4])
		c.B = (hexToByte(hex[5]) << 4) + hexToByte(hex[6])
		c.A = (hexToByte(hex[7]) << 4) + hexToByte(hex[8])
	case 7:
		c.R = (hexToByte(hex[1]) << 4) + hexToByte(hex[2])
		c.G = (hexToByte(hex[3]) << 4) + hexToByte(hex[4])
		c.B = (hexToByte(hex[5]) << 4) + hexToByte(hex[6])
	case 5:
		c.R = hexToByte(hex[1]) * 17
		c.G = hexToByte(hex[2]) * 17
		c.B = hexToByte(hex[3]) * 17
		c.A = hexToByte(hex[4]) * 17
	case 4:
		c.R = hexToByte(hex[1]) * 17
		c.G = hexToByte(hex[2]) * 17
		c.B = hexToByte(hex[3]) * 17
	default:
		err = ErrInvalidFormat
	}
	return
}
