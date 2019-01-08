# go-dice

An implementation of [Dice Coefficient](http://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient) algorithm in Go.
Code is based on [aceakash/string-similarity](https://github.com/aceakash/string-similarity)

## Installation

```bash
$ go get github.com/robopuff/go-dice
```

## Documentation

https://godoc.org/github.com/robopuff/go-dice

## Usage examples

```go
result := dice.ComparePair("iphone", "iphone x")
// result will be equal to 0.9090909090909091
```

```go
result := dice.FindBest("healed", []string{"mailed", "edward", "sealed", "theatre"})
// result will be an instance of dice.FindBestResults with "sealed" string scored highest (0.8)
```

## License

This software is licensed under the BSD-3-Clause License. [View the license](LICENSE).
