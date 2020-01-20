// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	// NOTE: ignoring potential errors from input.Err()
	defer c.Close()

	ch := make(chan string, 3)
	input := bufio.NewScanner(c)

	go func() {
		for input.Scan() {
			ch <- input.Text()
		}
	}()

	for {
		select {
		case <- time.After(10 * time.Second):
			fmt.Println("client idle for 10 seconds. server quit!")
			os.Exit(-1)
		case msg := <-ch:
			go echo(c, msg, 1*time.Second)
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
