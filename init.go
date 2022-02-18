package main

import (
	"flag"
	"log"
)

func init() {
	var e error

	opt = options{}

	flag.StringVar(&opt.domain, "d", "", "")
	flag.StringVar(&opt.domain, "domain", "", "")

	flag.BoolVar(&opt.silent, "s", false, "")
	flag.BoolVar(&opt.silent, "silent", false, "")

	flag.BoolVar(&opt.unreg, "unreg", false, "")

	flag.Usage = func() {
		showUsage(true)
	}
	flag.Parse()

	if opt.domain != "" {
		if !opt.silent {
			showUsage(false)
		}

		if opt.domain, e = validate(opt.domain); e != nil {
			log.Fatal(e)
		}
	} else {
		showUsage(true)
	}
}
