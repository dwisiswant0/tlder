package main

import (
	"fmt"
	"strings"

	"github.com/domainr/whois"
)

func whoIs(domain string, extension string) (string, error) {
	extension = strings.ToLower(extension)
	query := fmt.Sprint(domain, ".", extension)

	req, err := whois.NewRequest(query)
	if err != nil {
		return "", err
	}

	res, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		return "", err
	}

	return string(res.Body), nil
}
