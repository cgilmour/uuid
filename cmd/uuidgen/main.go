// Copyright 2015 Caleb Gilmour
// Use of this source code free and unencumbered software released into the public domain.
// For more information, refer to the LICENSE file or <http://unlicense.org/>

package main

import (
	"fmt"
	"flag"
	"github.com/cgilmour/uuid"
	"os"
)

var (
	file = flag.String("f", "", "file or device to read random data")
	upper = flag.Bool("u", false, "show the result in upper-case")
	n = flag.Uint("n", 1, "number of results to produce")
)

func main() {
	source := "(default)"

	flag.Parse()
	if *upper {
		uuid.Upper()
	}
	if *file != "" {
		if *file == "-" {
			// use stdin
			uuid.Source(os.Stdin)
			source = "(stdin)"
		} else {
			f, err := os.Open(*file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Cannot use '%s' as source: %s\n", *file, err)
			} else {
				defer f.Close()
				uuid.Source(f)
				source = fmt.Sprintf("'%s'", *file)
			}
		}
	}

	for i := uint(0); i < *n; i++ {
		s, err := uuid.New4()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating UUID #%d with source %s: %s\n", i+1, source, err)
			return
		}
		fmt.Println(s)
	}
}
