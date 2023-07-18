package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	// Set custom input
	input := "echo Hello\npwd\ncd ..\npwd\n:q\n"
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
	main()

	// Get the output
	expectedOutput := "$ Hello\n"					// echo Hello
	expectedOutput += `$ D:\wb-l2\develop\dev08`	//pwd
	expectedOutput += "\n" + `$ $ D:\wb-l2\develop` // cd + pwd
	expectedOutput += "\n$ Exiting\n"				//:q
	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	if !reflect.DeepEqual(string(output), expectedOutput) {
		t.Errorf("Expected output:\n%s\nBut got:\n%s ayta", expectedOutput, string(output))
	}
}
