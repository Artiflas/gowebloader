package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var flagOutput = flag.String("o", "", "file as target")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("please just one URL")
		os.Exit(1)
	}
	url := args[0]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error loading data", url, err)
		os.Exit(2)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
