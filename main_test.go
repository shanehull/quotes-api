package main

import "testing"

func Test_GenRandIndex(t *testing.T) {
	idx := genRandIndex()
	if idx < 0 || idx > len(quotes) {
		t.Fail()
	}
}
