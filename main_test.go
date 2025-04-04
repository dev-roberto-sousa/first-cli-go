package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to the coutn words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("palavra1 palavra2 palavra3 palavra4\n")

	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("expected %d, got %d instead. \n", exp, res)
	}

}

// TestCountLines tests the count function set to the coutn lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("palavra1 palavra2 palavra3\nline2\nline3\nline4")

	exp := 4

	res := count(b, true)

	if res != exp {
		t.Errorf("expected %d, got %d instead. \n", exp, res)
	}

}
