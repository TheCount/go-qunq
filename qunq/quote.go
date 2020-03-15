package qunq

// QuoteString returns the quoted version of s.
// If an error occurs, a partial result ending right before where the error
// occurred is returned.
func QuoteString(s string) (string, error) {
	b, err := QuoteBytes([]byte(s))
	return string(b), err
}

// QuoteBytes returns a new byte slice containing the quoted version of s.
// If an error occurs, a partial result ending right before where the error
// occurred is returned.
func QuoteBytes(s []byte) ([]byte, error) {
	result := make([]byte, 0, len(s)+2)
	result = append(result, '"')
	for _, c := range s {
		switch {
		case c == '"', c == '\\':
			result = append(result, '\\', c)
		case c == '\t', c >= 0x20 && c <= 0x7e, c >= 0x80:
			result = append(result, c)
		default:
			return result, ErrInvalidChar
		}
	}
	result = append(result, '"')
	return result, nil
}
