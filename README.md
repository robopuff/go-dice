[![GoDoc](https://godoc.org/github.com/robopuff/go-dice?status.svg)](https://godoc.org/github.com/robopuff/go-dice)

# go-dice

An implementation of [Dice coefficient](http://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient) algorithm in Go.
Code is based on [aceakash/string-similarity](https://github.com/aceakash/string-similarity)

## Installation

```bash
$ go get github.com/robopuff/go-dice
```

## Documentation

https://godoc.org/github.com/robopuff/go-dice

## Usage examples

```go
similarity := dice.ComparePair("healed", "sealed")
// `similarity` will be equal to 0.8
```

```go
bestIndex, results := dice.FindBest("healed", []string{"mailed", "edward", "sealed", "theatre"})
// Since highest score will have "sealed" `bestIndex` will be equal to `2`
// and results will provide an array of ordered float32 similarity scores
```

## License

This software is licensed under the BSD-3-Clause License. [View the license](LICENSE).
