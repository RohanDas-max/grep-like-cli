package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	filename := "test.txt"
	args := "kya"

	got := search(filename, args)
	var want int = 1

	if got < want {
		t.Error("got is not want", got, want)
	}
}

func TestOflag(t *testing.T) {

}
