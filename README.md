# hexcolor
Hex image/color implementation for go


[![Verify](https://github.com/g4s8/hexcolor/actions/workflows/go.yml/badge.svg)](https://github.com/g4s8/hexcolor/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/g4s8/hexcolor?status.svg)](https://godoc.org/github.com/g4s8/hexcolor)

// @todo #5:15min Add documentation for package: create doc.go file in root and explain package usage there.<br/>
//  Also add more docs for `Parse()` function.

## Usage

Download: `go get -u github.com/g4s8/hexcolor`

Example:
```go
import "github.com/g4s8/hexcolor"
import "image/color"

func main() {
  var c color.Color
  c, err := hexcolor.Parse("#fc12bd")
  if err != nil {
    panic(err)
  }
}
```

## Contribution
Fork repository, clone it, make changes,
check with `go build && go test && go vet`,
push to new branch and submit a pull request.
