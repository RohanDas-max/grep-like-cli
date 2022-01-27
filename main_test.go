package main

import (
	"os"
	"testing"
)

func TestSearchstdin(t *testing.T) {
	arg := "foo"
	input := []string{
		"foo", "bar", "foobaz", "barbazfoo", "food",
	}
	got := Searchstdin(input, arg)
	var want error = nil

	if got != want {
		t.Errorf("got %v instead of %v", got, want)
	}
}

func TestSearch(t *testing.T) {
	os.Create("test.txt")
	os.WriteFile("test.txt", []byte("Hello World"), 0400)
	filename := "test.txt"
	args := "Hello"
	defer os.Remove(filename)
	
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
