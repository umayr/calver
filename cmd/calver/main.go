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
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `calver is a small utility to handle calender versioning:

Usage:
  --format string
		format to parse the provided version (default "YYYY.MM.DD")
  --modifier string
		modifier for prerelease versions (default "dev")
  --pre-release
		flag to create a prerelease

Example:
  $ calver 2020.12.20
  2020.12.20-1

  $ calver 2020.12.20-1
  2020.12.20-2

  $ calver --pre-release 2020.12.20-2
  2020.12.20-dev.3

  $ calver --format YYYY.0W 2019.01
  2020.52

  $ calver --format YY.MM 19.01
  20.12

For more information about Calender Versioning, please visit https://calver.org
`)
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
