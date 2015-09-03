# gosymbol

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/gosymbol)
[![Build Status](https://travis-ci.org/shamsher31/gosymbol.svg)](https://travis-ci.org/shamsher31/gosymbol)

Symbol provide unicode symbol for your Go apps

### How to install
```go
go get github.com/shamsher31/gosymbol
```

### How to use
```go
package main

import (
  "fmt"
  "github.com/shamsher31/gosymbol"
)

func main() {

  // will give ℹ in blue color
  fmt.Print(symbol.Info())
  
  // will give ✔ in green color
  fmt.Print(symbol.Success())

  // will give ⚠ in yellow color
  fmt.Print(symbol.Warning())

  // will give ✖ in red color
  fmt.Print(symbol.Error())

  // will give ©
  fmt.Println(symbol.Copyright())
  
  // will give ®
  fmt.Println(symbol.Registered())
  
  // will give β
  fmt.Println(symbol.Beta())
  
  // will give δ
  fmt.Println(symbol.Delta())

}
```

### Related
[goisroot](https://github.com/shamsher31/goisroot)<br>
[gosudoblock](https://github.com/shamsher31/gosudoblock)<br>
[chalk](https://github.com/ttacon/chalk)<br>

### Why
This package is inspired by [log-symbols](https://www.npmjs.com/package/log-symbols) npm module but with more unicode symbols.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
