package main

import (
	// "bytes"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	// "os/signal"
	// "strings"
	// "sync"
	// "syscall"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ch := make(chan struct{})
	go func(ch chan struct{}) {
		go func() {

		}()
		err := startTCPServer("localhost", "8888", ch)
		if err != nil {
			log.Println("TCP server error:", err)
		}
		
	}(ch)

	// Delay to allow the server to start
	time.Sleep(200 * time.Millisecond)

	// Set custom input
	input := "Hello World\nOpenAI GPT-3.5\n"
	
	tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        log.Fatal(err)
    }

    defer os.Remove(tmpfile.Name()) // clean up

    err = ioutil.WriteFile(tmpfile.Name(), []byte(input), 0644)
	if err != nil {
		t.Fatalf("Ошибка записи во временный файл: %s", err)
	}

    oldStdin := os.Stdin
    defer func() { os.Stdin = oldStdin }() // Restore original Stdin

    os.Stdin = tmpfile

	// Перенаправление вывода программы на буфер
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	if err != nil{
		log.Fatal("errora  ",err)
	}


	// Run the main function
	os.Args = []string{"cmd", "-timeout", "1s", "localhost", "8888"}
	main()

	// Capture and check the output
	expectedOutput := "Hello World\nOpenAI GPT-3.5\nConnection closed by the server.\n"
	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout
	if string(output) != expectedOutput {
		t.Errorf("Unexpected output. Expected:\n %s\n Got:\n %s qwe", expectedOutput, output)
		// stopTCPServer()
	}
	// log.Println("Passed")
	ch <- struct{}{}
}

func startTCPServer(host, port string, ch chan struct{}) error {
	listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("TCP server accept error:", err)
			continue
		}
		go func() {
			select {
			case <-ch:
				// log.Println("Server closing connection")
				return
			}
		}()
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
