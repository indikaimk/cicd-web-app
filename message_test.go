package main

import (
	"testing"
)

// func TestHello(t *testing.T) {
// 	// want := "Hello world"
// 	// got := "rr"

// 	req, err := http.NewRequest("GET", "/hello", nil)

// 	// if got != want {
// 	// 	t.Errorf("got %q, wanted %q", got, want)
// 	// }
// 	if err != nil {
// 		t.Errorf("Error")
// 	} else {
// 		t.Errorf(req.Host)
// 	}
// }

func TestMessage(t *testing.T) {
	got, err := message("Hello")
	want := "Hello World"
	if got != want {
		t.Errorf("got %q, wanted %q: error %q", got, want, err.Error())
	}
}
