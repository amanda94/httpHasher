package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	u := "https://www.google.com"
	_, e := Hash(u)
	if e != nil {
		t.Errorf("%s shouldn't fail, error: %v", u, e)
	}

	u = "http://reddit.com/r/notfunny"
	_, e = Hash(u)
	if e != nil {
		t.Errorf("%s shouldn't fail, error: %v", u, e)
	}
	u = "https://notavalidwebsite.no"
	_, e = Hash(u)
	if e == nil {
		t.Errorf("this path should fail: %s", u)
	}
}
