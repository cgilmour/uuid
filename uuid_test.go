// Copyright 2015 Caleb Gilmour
// Use of this source code free and unencumbered software released into the public domain.
// For more information, refer to the LICENSE file or <http://unlicense.org/>

package uuid

import (
	"bytes"
	"testing"
)

func TestNew4(t *testing.T) {
	s, err := New4()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	checkValidUUIDv4(t, s)
}

func TestUpper(t *testing.T) {
	Upper()
	s, err := New4()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	checkValidUUIDv4(t, s)
	if !onlyDashAndUpperHex(s) {
		t.Errorf("Unexpected characters in UUID '%s'", s)
	}
}

func TestLower(t *testing.T) {
	Lower()
	s, err := New4()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	checkValidUUIDv4(t, s)
	if !onlyDashAndLowerHex(s) {
		t.Errorf("Unexpected characters in UUID '%s'", s)
	}
}

func TestCrappySource(t *testing.T) {
	// retain original source
	orig := source

	b := bytes.NewBufferString("1234567890123456")
	Source(b)
	s, err := New4()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	checkValidUUIDv4(t, s)
	s, err = New4()
	if err == nil {
		t.Errorf("Expected error from crappy source")
	}

	// restore original source
	source = orig
}

func checkValidUUIDv4(t *testing.T, s string) bool {
	if len(s) != 36 {
		t.Errorf("Unexpected result, len %d != 36", len(s))
		return false
	}
	for i, c := range s {
		switch i {
		// dashes
		case 8, 13, 18, 23:
			if c != '-' {
				t.Errorf("Unexpected character at pos %d: expected '-' in UUID '%s'", i+1, s)
				return false
			}
		case 14:
			if c != '4' {
				t.Errorf("Unexpected version '%c', expected '4' in UUID '%s'", s[14], s)
				return false
			}
		case 19:
			switch c {
			case '8', '9', 'a', 'b', 'A', 'B':
			default:
				t.Errorf("Unexpected reserved bits in '%c', expected 8/9/A/B/a/b in UUID '%s'", s[19], s)
				return false
			}
		default:
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
				'a', 'b', 'c', 'd', 'e', 'f',
				'A', 'B', 'C', 'D', 'E', 'F':
			default:
				t.Errorf("Unexpected character at pos %d: expected 0-9, a-f or A-F' in UUID '%s'", i+1, s)
				return false
			}
		}
	}
	return true
}

func onlyDashAndUpperHex(s string) bool {
	for _, c := range s {
		switch c {
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F':
		default:
			return false
		}
	}
	return true
}

func onlyDashAndLowerHex(s string) bool {
	for _, c := range s {
		switch c {
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f':
		default:
			return false
		}
	}
	return true
}

var s string

func BenchmarkNew4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, _ = New4()
		if len(s) != 36 {
			b.Fatalf("Benchmarking error: length check failed (%d != 36) at %d/%d", len(s), i+1, b.N)
		}
	}
}
