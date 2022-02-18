package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dwisiswant0/tlder/db"
)

func main() {
	for _, ext := range db.List {
		ext = strings.ToLower(ext)

		wg.Add(1)
		go func(domain string, extension string) {
			defer wg.Done()

			a, e := IsAvailable(domain, extension)
			if e != nil && !opt.silent {
				fmt.Fprintln(os.Stderr, e)
				return
			}

			d := fmt.Sprint(domain, ".", extension)

			if !a {
				fmt.Println("-", d)
			} else if opt.unreg {
				fmt.Println("-", d)
			}
		}(opt.domain, ext)
	}

	wg.Wait()
}
