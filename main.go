package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	host := flag.String("host", "localhost", "Set this as the hostname you want to resolve")
	flag.Parse()

	fmt.Println("Attempting to resolve:", *host)
	result, err := net.LookupHost(*host)
	if err != nil {
		fmt.Println("Host did not resolve")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(*host, "resolves to", result[0])
}
