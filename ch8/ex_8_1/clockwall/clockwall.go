package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// ./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
func main() {

	//clkSrvs := os.Args[1:] // clock servers

	zones := make(map[string]string)

	for _, arg := range os.Args[1:] {
		tokens := strings.Split(arg, "=")
		zones[tokens[0]] = tokens[1]
	}

	for k, v := range zones {
		time.Sleep(300 * time.Millisecond)
		go dial(k, v)
	}

	for {
		time.Sleep(time.Minute * 3)
	}
}

func dial(city string, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	copyTime(os.Stdout, conn, city)
}

func copyTime(dst io.Writer, src io.Reader, city string) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(dst, "%10s: %s\n", city, s.Text())
	}
}
