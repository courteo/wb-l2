package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestSliceInt_removeDuplicate(t *testing.T) {
	slice := SliceInt{
		{1, 2, 3},
		{1, 2, 3},
		{4, 5, 6},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := SliceInt{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	result := slice.removeDuplicate(slice)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ошибка в removeDuplicate. Ожидаемый результат: %+v, Полученный результат: %+v", expected, result)
	}
}

func TestSliceInt_sortColumn(t *testing.T) {
	slice := SliceInt{
		{6, 5, 4},
		{9, 8, 7},
		{3, 2, 1},

	}

	expected := SliceInt{
		{3, 2, 1},
		{6, 5, 4},
		{9, 8, 7},
	}

	result := slice.sortColumn(slice, 0, false)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ошибка в sortColumn. Ожидаемый результат: %+v, Полученный результат: %+v", expected, result)
	}
}

func TestSliceString_removeDuplicate(t *testing.T) {
	slice := SliceString{
		{"hello", "world"},
		{"hello", "world"},
		{"foo", "bar"},
		{"foo", "bar"},
		{"baz", "qux"},
	}

	expected := SliceString{
		{"hello", "world"},
		{"foo", "bar"},
		{"baz", "qux"},
	}

	result := slice.removeDuplicate(slice)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ошибка в removeDuplicate. Ожидаемый результат: %+v, Полученный результат: %+v", expected, result)
	}
}

func TestSliceString_sortColumn(t *testing.T) {
	slice := SliceString{
		{"world", "hello"},
		{"bar", "foo"},
		{"qux", "baz"},
	}

	expected := SliceString{
		{"qux", "baz"},
		{"bar", "foo"},
		{"world", "hello"},
	}

	result := slice.sortColumn(slice, 1, false)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ошибка в sortColumn. Ожидаемый результат: %+v, Полученный результат: %+v", expected, result)
	}
}

func TestMakeSort_Integers(t *testing.T) {
	content := "3 8 1\n6 9 4\n9 1 7"
	expectedOutput := "9 1 7\n3 8 1\n6 9 4\n"

	tmpFile, err := ioutil.TempFile("", "test_file.txt")
	if err != nil {
		t.Fatalf("Ошибка создания временного файла: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	err = ioutil.WriteFile(tmpFile.Name(), []byte(content), 0644)
	if err != nil {
		t.Fatalf("Ошибка записи во временный файл: %s", err)
	}

	var slice SliceInt
	slice.MakeSort(bufio.NewScanner(tmpFile), 1, false, false)

	// Проверка содержимого файла
	fileContent, err := ioutil.ReadFile("outputInt.txt")
	if err != nil {
		t.Fatalf("Ошибка чтения файла: %s", err)
	}

	if !bytes.Equal(fileContent, []byte(expectedOutput)) {
		t.Errorf("Ошибка в MakeSort для чисел. Ожидаемый результат: %q, Полученный результат: %q", expectedOutput, fileContent)
	}
}

func TestMakeSort_Strings(t *testing.T) {
	content := "world hello\nbar foo\nqux baz\n"
	expectedOutput := "qux baz\nbar foo\nworld hello\n"

	tmpFile, err := ioutil.TempFile("", "test_file.txt")
	if err != nil {
		t.Fatalf("Ошибка создания временного файла: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	err = ioutil.WriteFile(tmpFile.Name(), []byte(content), 0644)
	if err != nil {
		t.Fatalf("Ошибка записи во временный файл: %s", err)
	}

	var slice SliceString
	slice.MakeSort(bufio.NewScanner(tmpFile), 1, false, false)

	// Проверка содержимого файла
	fileContent, err := ioutil.ReadFile("outputString.txt")
	if err != nil {
		t.Fatalf("Ошибка чтения файла: %s", err)
	}

	if !bytes.Equal(fileContent, []byte(expectedOutput)) {
		t.Errorf("Ошибка в MakeSort для строк. Ожидаемый результат: %q, Полученный результат: %q", expectedOutput, fileContent)
	}
}