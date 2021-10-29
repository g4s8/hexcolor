package hexcolor

import (
	"errors"
	"image/color"
)

// Parse - parse hex string to image/color
func Parse(hex string) (c color.RGBA, err error) {
	// got this parsing code from here https://stackoverflow.com/a/54200713/1723695

	var errInvalidFormat = errors.New("invalid format")

	c.A = 0xff

	if hex[0] != '#' {
		return c, errInvalidFormat
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
		err = errInvalidFormat
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
		err = errInvalidFormat
	}
	return
}
