# NanoID

NanoID is a lightweight Go package for generating short, unique IDs. It's a great alternative to longer, more complex UUIDs, especially when you need something more user-friendly.

## Installation

```sh
go get github.com/9ssi7/nanoid
```

## Usage

```go
import "github.com/9ssi7/nanoid"

id, err := nanoid.New()
```

### Custom Length ID

```go
id, err := nanoid.NewWithLength(nanoid.LengthMedium)
if err != nil {
 log.Fatal(err)
}
```

### Validating ID

```go
if !nanoid.Validate(id) {
 log.Fatal("Invalid ID")
}
```

## Benchmark

NanoID is highly performant and compares favorably with Google's UUID package.

```txt
goos: linux
goarch: amd64
pkg: github.com/9ssi7/nanoid
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkNanoID
BenchmarkNanoID-20          1240796        933.8 ns/op
BenchmarkGoogleUUID
BenchmarkGoogleUUID-20      1312452        903.1 ns/op
```

As seen in the benchmark results, NanoID's performance is very close to Google UUID, making it a suitable choice for most use cases where short IDs are preferred.

## Security

NanoID uses `crypto/rand` for generating secure random IDs, reducing the likelihood of collision. However, it's crucial to consider the trade-off between ID length and collision probability. Shorter IDs are more user-friendly but have a higher chance of collision, especially in systems with a high number of generated IDs.
