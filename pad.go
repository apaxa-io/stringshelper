package stringshelper

import (
	"golang.org/x/text/unicode/norm"
	"strings"
)

// Len returns number of glyph in UTF-8 encoded string.
func Len(s string) (l int) {
	var ia norm.Iter
	ia.InitString(norm.NFKD, s)
	for !ia.Done() {
		l++
		ia.Next()
	}
	return
}

// PadLeft returns original string padded on the left side with glyph to length l.
// Length for this function is in glyphs, not bytes or runes.
// Passed glyph should be single glyph (printable char) else result will be of wrong length. But where is no check for that.
func PadLeft(s, glyph string, l int) string {
	curLen := Len(s)
	if curLen >= l {
		return s
	}
	return strings.Repeat(glyph, l-curLen) + s
}

// PadRight returns original string padded on the right side with glyph to length l.
// Length for this function is in glyphs, not bytes or runes.
// Passed glyph should be single glyph (printable char) else result will be of wrong length. But where is no check for that.
func PadRight(s, glyph string, l int) string {
	curLen := Len(s)
	if curLen >= l {
		return s
	}
	return s + strings.Repeat(glyph, l-curLen)
}

// PadLeftWithBytes return original string padded on the left side with char b to length l.
// Length for this function is in bytes. So this function is good only for ascii strings.
func PadLeftWithByte(s string, b byte, l int) string {
	if len(s) >= l {
		return s
	}
	return strings.Repeat(string(b), l-len(s)) + s
}

// PadRightWithBytes return original string padded on the left side with char b to length l.
// Length for this function is in bytes. So this function is good only for ascii strings.
func PadRightWithByte(s string, b byte, l int) string {
	if len(s) >= l {
		return s
	}
	return s + strings.Repeat(string(b), l-len(s))
}
