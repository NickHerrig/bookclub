package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net/http"
	"os"
)

//go:embed index.html
var index []byte

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write(index)
}

// Let's Go: Chapter 6.3 Request Logging
func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func run() int {
	addr := flag.String("addr", "127.0.0.1:5000", "listen address (host:port)")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleIndex)

	// apply middleware to router
	app := AccessLogMiddleware(mux)

	fmt.Printf("listening on %s\n", *addr)
	err := http.ListenAndServe(*addr, app)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
