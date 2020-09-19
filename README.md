# properties

A dead-simple file format for configuration and metadata `.properties` files. Basically just `key = value` pairs, one per line. This format is intended to be as easy as possible for anyone to implement a parser in their language of choice, and is an easy "least common denominator" for heterogeneous systems. JSON is awesome, but sometimes you just need a simple way to exchange keys/values between languages/systems. This is especially helpful when using Go, which does not deal well with dynamic data structures and types.

```
type = blueprints
orientation = above
size = huge
style = line
```

As you can see, the only reserved characters are newline (`\n`) and equals (`=`). Leading and trailing whitespace is stripped from all keys and values, leaving you with a familiar `map[string]string` type that can be used as such.

[More docs on pkg.go.dev](https://pkg.go.dev/github.com/gershwinlabs/properties)
