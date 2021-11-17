// The MIT License (MIT) Copyright (c) 2019-2021 artipie.com
// https://github.com/g4s8/hexcolor/blob/master/LICENSE

/*
Package hexcolor allows to parse HEX representation of
color which is frequently used for web colors and in
other cases.

This package exposes single public entry point function:

	Parse(hex string) (c color.RGBA, err error)

It uses implementation suggested in StackOverflow
answer: https://stackoverflow.com/a/54200713/1723695
as "Fast Solution".

Supported formats:
 - RGB: #fc3 - red=ff, green=cc, blue=33
 - RRGGBB: #abcd02 - red=ab, green=cd, bluen=02
 - ARGB: #dfab - alpha=dd, RGB=fab
 - AARRGGBB: #c3a0b1c2 - alpha=c3 RRGGBB=a0b1c2

It supports uppercase and lowercase hex symbols.
All hex symbols should start with `#` char.
*/
package hexcolor
