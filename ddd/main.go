package main

import (
	"fmt"
	"github.com/libp2p/go-reuseport"
	"html"
	"net/http"
	"os"
)

func main() {
	listener, err := reuseport.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(os.Getgid())
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	panic(server.Serve(listener))
}
