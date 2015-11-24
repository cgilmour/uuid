// Copyright 2015 Caleb Gilmour
// Use of this source code free and unencumbered software released into the public domain.
// For more information, refer to the LICENSE file or <http://unlicense.org/>

package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

var source io.Reader = rand.Reader

var format = lowerFormat

const (
	upperFormat = "%08X-%04X-%04X-%04X-%012X"
	lowerFormat = "%08x-%04x-%04x-%04x-%012x"
)

func New4() (string, error) {
	// read 128 bits from random source.
	// note: only 122 bits get used, but this is easier
	b := make([]byte, 16)
	_, err := io.ReadFull(source, b)
	if err != nil {
		return "", err
	}
	// set version to 4
	b[6] = (b[6] | 0x40) & 0x4F
	b[8] = (b[8] | 0x80) & 0xBF
	return fmt.Sprintf(format, b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}

func Upper() {
	format = upperFormat
}

func Lower() {
	format = lowerFormat
}

func Source(r io.Reader) {
	source = r
}
