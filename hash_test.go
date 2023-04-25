package main

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	r, e := Hash("https://www.google.com")
	if e != nil {
		t.Errorf("%v", e)
	}
	fmt.Println(r, e)

	r, e = Hash("http://reddit.com/r/notfunny")
	if e != nil {
		t.Errorf("%v", e)
	}

	r, e = Hash("https://notavalidwebsite.no")
	fmt.Println(r, e)

}
