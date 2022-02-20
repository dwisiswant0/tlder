package fingerprint

import "strings"

func init() {
	Available, Registered, Unregistered = strings.Split(ava, "\n"), strings.Split(reg, "\n"), strings.Split(unr, "\n")
}
