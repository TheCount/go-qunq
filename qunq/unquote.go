package qunq

// UnquoteString returns the unquoted version of s.
// If an error occurs, a partial result ending right before where the error
// occurred is returned.
func UnquoteString(s string) (string, error) {
	b, err := UnquoteBytes([]byte(s))
	return string(b), err
}

// UnquoteBytes returns a new byte slice containing the unquoted version of s.
// If an error occurs, a partial result ending right before where the error
// occurred is returned.
func UnquoteBytes(s []byte) ([]byte, error) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return nil, ErrSyntax
	}
	s = s[1 : len(s)-1]
	result := make([]byte, 0, len(s))
	backslash := false
	for _, c := range s {
		switch {
		case !backslash && c == '"':
			return result, ErrSyntax
		case !backslash && c == '\\':
			backslash = true
		case c == '\t', c >= 0x20 && c <= 0x7e, c >= 0x80:
			result = append(result, c)
			backslash = false
		default:
			return result, ErrInvalidChar
		}
	}
	if backslash { // final double quote was escaped
		return result, ErrSyntax
	}
	return result, nil
}
