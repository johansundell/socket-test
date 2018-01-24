package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ln, err := net.Listen("unix", "/tmp/sudde.sock")
	if err != nil {
		log.Fatal(err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(c chan os.Signal) {
		// Wait for a SIGINT or SIGKILL:
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		// Stop listening (and unlink the socket if unix type):
		ln.Close()
		// And we're done:
		os.Exit(0)
	}(sigc)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		go testService(conn)
	}
}

func testService(conn net.Conn) {
	timeoutDuration := 5 * time.Second
	bufReader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		str, err := bufReader.ReadString('\n')
		if err != nil {
			/*if err == io.EOF {
				break
			}*/
			log.Println(err)
			break
		}
		log.Println(str)
	}
	// TODO: This does not work
	log.Println("About to write")
	bufWriter := bufio.NewWriter(conn)
	n, err := bufWriter.WriteString("Tack så mycket")
	if err != nil {
		log.Println(err)
	}
	log.Println("Wrote", n)
}
