package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCompare(t *testing.T) {
	text := "Hello, World!"
	word := "hello"
	ignoreCase := true
	invert := false

	result := compare(text, word, ignoreCase, invert)

	if result != true {
		t.Errorf("Ошибка в compare. Ожидаемый результат: true, Полученный результат: %v", result)
	}
}

func TestMain(t *testing.T) {
	content := "Hello, World!\nThis is a test.\nLine with word .\n"
	expectedOutput := "Line 1 - [This is a test.]\nLine 2 - [Line with word .]\n"

	tmpFile, err := ioutil.TempFile("", "test_file.txt")
	if err != nil {
		t.Fatalf("Ошибка создания временного файла: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	err = ioutil.WriteFile(tmpFile.Name(), []byte(content), 0644)
	if err != nil {
		t.Fatalf("Ошибка записи во временный файл: %s", err)
	}

	// Перенаправление вывода программы на буфер
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w


	// Замена аргументов командной строки для теста
	os.Args = []string{"cmd", "-b", "1", "word", tmpFile.Name()}

	// Запуск программы
	main()

	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	// Проверка вывода программы
	if string(output) != expectedOutput {
		t.Errorf(`Ошибка в Main. 
		Ожидаемый результат: 	%q 
		Полученный результат: 	%q`,
		 expectedOutput, string(output))
	}
}