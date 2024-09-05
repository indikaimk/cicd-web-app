//  from tutorial https://learning-cloud-native-go.github.io/docs/hello-world-server/

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", mux)
}
