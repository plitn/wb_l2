package dev10

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func Telnet() {
	timeout := flag.Duration("timeout", 10*time.Second, "Conn timeout")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("wrong usage")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))

	if err != nil {
		fmt.Fprintf(os.Stderr, "addr error: %s\n", err)
		os.Exit(1)
	}

	conn, err := net.DialTimeout("tcp", addr.String(), *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conn error: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	go serverWork(&conn)
	readServer(&conn)
}

func readServer(conn *net.Conn) {
	data := make([]byte, 1024)
	for {
		n, err := (*conn).Read(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "server reading err: %s\n", err)
			break
		}
		if n > 0 {
			fmt.Print(string(data[:n]))
		}
	}
}

func serverWork(conn *net.Conn) {
	for {
		data := make([]byte, 1024)
		n, err := os.Stdin.Read(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading error: %s\n", err)
			break
		}
		if n > 0 {
			_, err := (*conn).Write(data[:n])
			if err != nil {
				fmt.Fprintf(os.Stderr, "writing error: %s\n", err)
				break
			}
		}
	}
}
