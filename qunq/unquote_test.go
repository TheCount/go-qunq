package qunq

import (
	"testing"
)

// unquoteTestcases is a list of testcases for the unquote functions.
var unquoteTestcases = []Testcase{
	Testcase{``, ``, ErrSyntax},
	Testcase{`"`, ``, ErrSyntax},
	Testcase{`""`, ``, nil},
	Testcase{`foo`, ``, ErrSyntax},
	Testcase{`"foo"`, `foo`, nil},
	Testcase{"\"\t\"", "\t", nil},
	Testcase{"\"\x7f\"", ``, ErrInvalidChar},
	Testcase{`"""`, ``, ErrSyntax},
	Testcase{`"\"`, ``, ErrSyntax},
	Testcase{`"quote \"foo\" \\unquote"`, `quote "foo" \unquote`, nil},
}

// TestUnquoteString tests the UnquoteString function.
func TestUnquoteString(t *testing.T) {
	for i, tc := range unquoteTestcases {
		result, err := UnquoteString(tc.Input)
		if err != tc.ExpectedError {
			t.Errorf("Expected error '%s' in testcase %d, got '%s'", tc.ExpectedError,
				i, err)
		}
		if result != tc.Output {
			t.Errorf("Expected '%s' in testcase %d, got '%s'", tc.Output, i, result)
		}
	}
}
