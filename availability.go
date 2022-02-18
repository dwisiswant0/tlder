package main

import (
	"strings"

	"github.com/dwisiswant0/tlder/fingerprint"
)

// IsAvailable domain check from whois response with predefined fingerprint
func IsAvailable(domain string, extension string) (bool, error) {
	var s bool

	resp, err := whoIs(domain, extension)
	if err != nil {
		return s, err
	}

	if len([]byte(resp)) == 0 {
		return s, errNoResponses
	}

	if strings.Contains(resp, "RATE LIMIT EXCEEDED") {
		return s, errRateLimited
	}

	a := sliceContains(strings.ToLower(resp), fingerprint.Available)
	u := sliceContains(resp, fingerprint.Unregistered)
	// r := sliceContains(resp, fingerprint.Registered)

	if a || u {
		return !s, nil
	}

	return s, nil
}
