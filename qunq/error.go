package qunq

import (
	"errors"
)

const (
	// InvalidChars is the the set of characters which must not occur in quoted
	// strings.
	InvalidChars = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0a\x0b\x0c\x0d\x0e" +
		"\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x7f"
)

var (
	// ErrInvalidChar indicates an invalid character was encountered. Invalid
	// characters are the ASCII control characters other that horizontal tab
	// (i. e., all bytes from 0 to 31 except 9, and 127).
	ErrInvalidChar = errors.New("invalid character")

	// ErrSyntax indicates a syntax error was encountered in a quoted string.
	// A quoted string must begin with a double quote, end with an unescaped
	// double quote, and must otherwise not contain unescaped double quotes.
	ErrSyntax = errors.New("quoted-string syntax error")
)
