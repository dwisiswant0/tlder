package fingerprint

import _ "embed"

var (
	//go:embed available.txt
	ava string
	//go:embed registered.txt
	reg string
	//go:embed unregistered.txt
	unr string // Available but lowercase

	Available, Registered, Unregistered []string
)
