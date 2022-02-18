package db

import _ "embed"

var (
	//go:embed tlds.txt
	tlds string
	List []string
)
