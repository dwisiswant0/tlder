# TLD:er

TLDs finder — check domain name availability across all valid top-level domains.

## Installation

- Get pre-built binary from [releases page](https://github.com/dwisiswant0/tlder/releases), or
- If you have [Go1.19+](https://go.dev/dl/) compiler installed & configured: `go install github.com/dwisiswant0/tlder@latest`.

## Usage

```console
> tlder -h
   ________  ___         
  /_  __/ / / _ \___ ____
   / / / /_/ // / -_) __/
  /_/ /___/____/\__/_/   
  ---
  @dwisiswant0


Usage:
  tlder -d NAME

Options:
  -d, --domain <NAME>   Domain NAME to permutate
  -s, --silent          Silent mode
      --unreg           Show unregistered TLDs only
```

## Data

TLDs data taken from [data.iana.org](https://data.iana.org/TLD/). Checks for data updates are carried out weekly by [GitHub workflow](https://github.com/dwisiswant0/tlder/actions/workflows/update.yaml).

## Library

You can use **TLDer** as library. For example:

```go
package main

import "github.com/dwisiswant0/tlder"

func main() {
	name := "dw1"
	ext := "io"

	avail, err := tlder.IsAvailable(name, ext)
	if err != nil {
		panic(err)
	}

	print(avail)
}
```

## License

**tlder** is distributed under MIT, contributions are welcome! See `LICENSE`.