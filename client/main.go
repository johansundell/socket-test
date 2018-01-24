package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "/tmp/sudde.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	bufReader := bufio.NewReader(conn)
	fmt.Fprint(conn, "Testar detta\n")
	// TODO: This does not work
	str, err := bufReader.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(str)

}
