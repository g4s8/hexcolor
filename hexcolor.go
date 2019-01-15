package hexcolor

import (
	"fmt"
	"image/color"
)

// Parse - parse hex string to image/color
func Parse(hex string) (color.Color, error) {
	// got this parsing code from here https://stackoverflow.com/a/54200713/1723695
	var c color.RGBA
	var err error
	c.A = 0xff
	switch len(hex) {
	case 7:
		_, err = fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(hex, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	return c, err
}
