# gosymbol

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/gosymbol)

Symbol provide colored log icons

### How to install
```go
go get github.com/shamsher31/gosymbol
```

### How to use
```go
package main

import (
  "github.com/shamsher31/gosymbol"
)

func main() {

  // will give ℹ in blue color
  symbol.Info()
  
  // will give ✔ in green color
  symbol.Success()

  // will give ⚠ in yellow color
  symbol.Warning()

  // will give ✖ in red color
  symbol.Error()

}
```

### Related
[goisroot](https://github.com/shamsher31/goisroot)<br>
[gosudoblock](https://github.com/shamsher31/gosudoblock)<br>
[chalk](https://github.com/ttacon/chalk)<br>

### Why
This package is inspired by [log-symbols](https://www.npmjs.com/package/log-symbols) npm module.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
