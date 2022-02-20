package main

import (
	"regexp"
	"strings"
)

func validate(domain string) (string, error) {
	domain = strings.ToLower(domain)

	d, _ := regexp.MatchString(`^-`, domain)
	if d {
		return domain, errInputDomain
	}

	r, _ := regexp.MatchString(`^[a-z0-9-]{1,63}$`, domain)
	if !r {
		return domain, errInputDomain
	}

	return domain, nil
}
