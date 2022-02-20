package main

import (
	"errors"
	"sync"
)

var (
	opt options
	wg  sync.WaitGroup

	errInputDomain = errors.New("invalid input domain")
	errRateLimited = errors.New("whois query rate-limited, please wait & try again")
	errNoResponses = errors.New("no query responses from whois")
)
