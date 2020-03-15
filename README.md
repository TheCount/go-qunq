# RFC 7230 quoted string support

![](https://github.com/TheCount/go-qunq/workflows/CI/badge.svg)
[![Documentation](https://godoc.org/github.com/TheCount/go-qunq/qunq?status.svg)](https://godoc.org/github.com/TheCount/go-qunq/qunq)

## Description

This package provides quoted string support similar to the `strconv.Quote` and
`strconv.Unquote` class of functions, but specifically tailored to the
[RFC 7230](https://tools.ietf.org/html/rfc7230) definition of quoted strings.

Please see the
[package documentation](https://godoc.org/github.com/TheCount/go-qunq/qunq)
for details.

## Example

```golang
package main

import "github.com/TheCount/go-qunq/qunq"

const testString = `quote "scary" unquote`

func main() {
  quoted, err := qunq.QuoteString(testString)
  if err != nil {
    panic(err)
  }
  println(quoted)
  unquoted, err := qunq.UnquoteString(quoted)
  if err != nil {
    panic(err)
  }
  println(unquoted)
}
```
