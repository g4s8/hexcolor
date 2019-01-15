# hexcolor
Hex image/color implementation for go

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

[![DevOps By Rultor.com](http://www.rultor.com/b/g4s8/hexcolor)](http://www.rultor.com/p/g4s8/hexcolor)

[![Build Status](https://img.shields.io/travis/g4s8/hexcolor.svg?style=flat-square)](https://travis-ci.org/g4s8/hexcolor)

[![PDD status](http://www.0pdd.com/svg?name=g4s8/hexcolor)](http://www.0pdd.com/p?name=g4s8/hexcolor)
[![License](https://img.shields.io/github/license/g4s8/hexcolor.svg?style=flat-square)](https://github.com/g4s8/hexcolor/blob/master/LICENSE)



## Contribution
Fork repository, clone it, make changes,
push to new branch and submit a pull request.



