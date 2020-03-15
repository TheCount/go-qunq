package qunq

import (
	"testing"
)

// quoteTestcases is a list of testcases for the quote functions.
var quoteTestcases = []Testcase{
	Testcase{``, `""`, nil},
	Testcase{`"`, `"\""`, nil},
	Testcase{`\`, `"\\"`, nil},
	Testcase{"\x00", `"`, ErrInvalidChar},
	Testcase{"\x7f", `"`, ErrInvalidChar},
	Testcase{"\x80", "\"\x80\"", nil},
	Testcase{"First good, then bad\nignored", `"First good, then bad`,
		ErrInvalidChar},
}

// TestQuoteString tests the QuoteString function.
func TestQuoteString(t *testing.T) {
	for i, tc := range quoteTestcases {
		result, err := QuoteString(tc.Input)
		if err != tc.ExpectedError {
			t.Errorf("Expected error '%s' in testcase %d, got '%s'", tc.ExpectedError,
				i, err)
		}
		if result != tc.Output {
			t.Errorf("Expected '%s' in testcase %d, got '%s'", tc.Output, i, result)
		}
	}
}

// TestQuoteBytes tests the QuoteString function.
func TestQuoteBytes(t *testing.T) {
	for i, tc := range quoteTestcases {
		result, err := QuoteBytes([]byte(tc.Input))
		if err != tc.ExpectedError {
			t.Errorf("Expected error '%s' in testcase %d, got '%s'", tc.ExpectedError,
				i, err)
		}
		if string(result) != tc.Output {
			t.Errorf("Expected '%s' in testcase %d, got '%s'", tc.Output, i, result)
		}
	}
}
