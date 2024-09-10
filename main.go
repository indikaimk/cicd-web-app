//  from tutorial https://learning-cloud-native-go.github.io/docs/hello-world-server/

package main

import (
	"errors"
	"io"
	"net/http"
)

func message(greeting string) (string, error) {
	if greeting == "" {
		return greeting, errors.New("empty string")
	} else {
		m := greeting + " World"
		return m, nil
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	m, err := message("Hello")
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
	} else {
		io.WriteString(w, m)
	}
}

func version(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "v1.0.18")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/version", version)

	http.ListenAndServe(":8080", mux)
}
