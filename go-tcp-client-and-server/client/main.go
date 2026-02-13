package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(conn)
	output, err := r.ReadString(byte('\n'))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, output)
}
