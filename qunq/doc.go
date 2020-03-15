// Package qunq implements quoting and unquoting of strings according to the
// quoted-string definition of RFC 7230.
//
// The functionality of this package is similar to to the strconv.Quote and
// strconv.Unquote class of functions, except that RFC 7230 specifically
// disallows certain octets to appear in quoted or unquoted strings.
package qunq
