package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var usage = `Usage: %s [options] 
Options are:
    -p port     Port
    -d dir      Root dir
`
var (
	port int
	dir  string
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}

	flag.IntVar(&port, "p", 1985, "")
	flag.StringVar(&dir, "d", "", "")
	flag.Parse()

	flag.Usage()

	fmt.Printf("%v, %v", port, dir)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
