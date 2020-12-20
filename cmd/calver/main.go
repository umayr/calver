package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/umayr/calver"
)

var (
	flagFormat   = flag.String("format", "YYYY.MM.DD", "format to parse the provided version")
	flagPre      = flag.Bool("pre-release", false, "flag to create a prerelease")
	flagModifier = flag.String("modifier", "dev", "modifier for prerelease versions")
)

func init() {
	usage := flag.Usage

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "%s is a small utility to handle calender versioning:\n\n", os.Args[0])

		usage()

		_, _ = fmt.Fprint(os.Stderr, "\nFor more information on Calender Versioning: https://calver.org\n")
	}

	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	version := args[len(args)-1]

	c, err := calver.Parse(version, *flagFormat, *flagModifier)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var next string
	if *flagPre {
		next = c.PreRelease()
	} else {
		next = c.Release()
	}

	fmt.Println(next)
}
