package qunq

// Testcase describes a testcase for the (un)quote functions.
type Testcase struct {
	// Input and Output are the input and expected output strings, respectively,
	// for the (un)quote functions.
	Input, Output string

	// ExpectedError is the error expected to be returned for this test case,
	// if any.
	ExpectedError error
}
