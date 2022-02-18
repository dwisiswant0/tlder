package main

import (
	"fmt"
	"os"
	"strings"
)

func showUsage(u bool) {
	fmt.Fprintf(os.Stderr, "%s\n  ---\n  %s\n\n", banner, author)

	if u {
		fmt.Fprintf(os.Stderr, "%s\n\n", usage)
		os.Exit(2)
	}
}

func sliceContains(s string, subslc []string) bool {
	for _, slice := range subslc {
		if strings.Contains(s, slice) {
			return true
		}
	}

	return false
}
