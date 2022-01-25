package main

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	filename := "test.txt"
	args := "kya"
	got := search(filename, args)
	want := true
	if got != want {
		t.Error("got is not want", got, want)
	}
}

func TestOflag(t *testing.T) {
	filename := "out.txt"
	data := "any thing"
	got := oflag(filename, data)
	var want error = nil
	if got != nil {
		t.Errorf("got %v instead %v", got, want)
	}
	os.Remove(filename)
}
