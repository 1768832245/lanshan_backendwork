package main

import "testing"

func TestJustDoIt(t *testing.T) {
	result := JustDoIt("/")
	result(1, 2)
}
