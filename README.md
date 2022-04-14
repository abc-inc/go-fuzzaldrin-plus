# go-fuzzaldrin-plus - fuzzy filtering and string scoring

This is Go port of [fuzzaldrin](https://github.com/atom/fuzzaldrin),
a library used by Atom and so its focus is on scoring and filtering paths,
methods, and other things common when writing code.
It therefore specializes in handling common patterns in these types of strings
such as characters like `/`, `-`, and `_`, and also handling of camel case text.

## Usage

```sh
go get -u github.com/abc-inc/go-fuzzaldrin-plus
```

```go
import "github.com/abc-inc/go-fuzzaldrin-plus"
```

### Filter

`Filter` sorts and filters the given items by matching them against the query.
It returns a slice of items sorted by best match against the query.

```go
func ExampleFilter_string() {
	ss := []string{"Call", "Me", "Maybe"}
	res := fuzzaldrin.Filter(ss, "me", ident[string])
	fmt.Println(strings.Join(res, "\n"))
	// Output:
	// Me
	// Maybe
}
```

### Match

`Match` returns the indices of any query matches in the given string.

```go
func ExampleMatch() {
	fmt.Println(fuzzaldrin.Match("Call Me Maybe", "m m"))
	fmt.Println(fuzzaldrin.Match("Call Me Maybe", "eY"))
	// Output:
	// [5 8]
	// [6 10]
}
```

### Score

`Score` calculates the score of the given string against the query.

```go
func ExampleScore() {
	fmt.Println(fuzzaldrin.Score("Me", "me"))
	fmt.Println(fuzzaldrin.Score("Maybe", "me"))
	// Output:
	// 0.17099999999999999
	// 0.0693
}
```

## Similar Projects

- https://github.com/Sixeight/go-fuzzaldrin - supports only strings and no structs
