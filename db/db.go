package db

import "strings"

func init() {
	List = strings.Split(tlds, "\n")
	List = List[1 : len(List)-1]
}
