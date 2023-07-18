package main

import (
	// "bytes"
	"io/ioutil"
	"log"
	"os"
	// "strings"
	"testing"
)

func TestMainFlow(t *testing.T) {
	content := []byte("Hello World\nOpenAI GPT-3.5\n")
    tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        log.Fatal(err)
    }

    defer os.Remove(tmpfile.Name()) // clean up

    err = ioutil.WriteFile(tmpfile.Name(), []byte(content), 0644)
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
	main()


	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout
	// Capture and check the output
	expectedOutput := "Hello World \nOpenAI GPT-3.5 \n"
	if string(output) != expectedOutput {
		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, output)
	}
}

func TestFieldsOption(t *testing.T) {
	// Set custom input
	content := "Hello World\nOpenAI GPT-3.5\n"
	tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        log.Fatal(err)
    }

    defer os.Remove(tmpfile.Name()) // clean up

    err = ioutil.WriteFile(tmpfile.Name(), []byte(content), 0644)
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


	// Set custom command-line arguments
	os.Args = []string{"cmd", "-f", "2,1"}


	// Run the main function
	main()

	// Capture and check the output
	expectedOutput := "World Hello GPT-3.5 OpenAI "

	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	if string(output) != expectedOutput {
		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, output)
	}
}

func TestDelimiterOption(t *testing.T) {
	// Set custom input
	content := "Hello World\nOpenAI GPT-3.5\n"
	tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        log.Fatal(err)
    }

    defer os.Remove(tmpfile.Name()) // clean up

    err = ioutil.WriteFile(tmpfile.Name(), []byte(content), 0644)
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

	os.Args = []string{"cmd", "-d", `";"`}


	// Run the main function
	main()

	// Capture and check the output
	expectedOutput := "Hello World \nOpenAI GPT-3.5 \n"

	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	if string(output) != expectedOutput {
		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, output)
	}
}

// func TestSeparatedOption(t *testing.T) {
// 	// Set custom input
// 	content := "Hello World\nOpenAI GPT-3.5\n"
// 	tmpfile, err := ioutil.TempFile("", "example")
//     if err != nil {
//         log.Fatal(err)
//     }

//     defer os.Remove(tmpfile.Name()) // clean up

//     err = ioutil.WriteFile(tmpfile.Name(), []byte(content), 0644)
// 	if err != nil {
// 		t.Fatalf("Ошибка записи во временный файл: %s", err)
// 	}

//     oldStdin := os.Stdin
//     defer func() { os.Stdin = oldStdin }() // Restore original Stdin

//     os.Stdin = tmpfile

// 	// Перенаправление вывода программы на буфер
// 	oldStdout := os.Stdout
// 	r, w, err := os.Pipe()
// 	os.Stdout = w
// 	if err != nil{
// 		log.Fatal("errora  ",err)
// 	}

// 	os.Args = []string{"cmd", "-s"}

// 	// Run the main function
// 	main()

// 	// Capture and check the output
// 	expectedOutput := "Hello World OpenAI GPT-3.5 "

// 	w.Close()
// 	output, _ := ioutil.ReadAll(r)
// 	os.Stdout = oldStdout
// 	if  string(output) != expectedOutput {
// 		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, output)
// 	}
// }