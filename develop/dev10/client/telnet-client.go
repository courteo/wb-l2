package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)
var timeout time.Duration

func regTimeVar(p *time.Duration, name string, value time.Duration, usage string) {
    if flag.Lookup(name) == nil {
        flag.DurationVar(p, name, value, usage)
    }
}


func getTimeFlag(name string) time.Duration {
    return flag.Lookup(name).Value.(flag.Getter).Get().(time.Duration)
}


func init() {
    regTimeVar(&timeout, "timeout", 10*time.Second, "Connection timeout" )
}

func initFlags() {
    timeout = getTimeFlag("timeout")
}


func main() {
	flag.Parse()
	initFlags()
	host := flag.Arg(0)
	port := flag.Arg(1)
	log.Println(timeout)
	// check existing of host or port
	if host == "" || port == "" {
		fmt.Println("No host or port")
		os.Exit(1)
	}

	// For catch ctrl +C
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Connect
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
		// os.Exit(1)
	}
	defer conn.Close()

	// For close connection
	closeCh := make(chan struct{})

	// read from socket
	go func() {
		io.Copy(os.Stdout, conn)
		close(closeCh)
	}()

	// read from stdin to socket
	go func() {
		io.Copy(conn, os.Stdin)
		conn.Close()
	}()

	// wait close connection or ctrcl +C
	select {
	case <-signalCh:
		fmt.Println("Closing the connection...")
	case <-closeCh:
		fmt.Println("Connection closed by the server.")
	}

	<-closeCh
	timeout = 10*time.Second
}