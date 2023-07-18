package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDownloadSite(t *testing.T) {


	// Запускаем функцию `downloadSite` для загрузки содержимого тестовой страницы
	downloadSite("https://gzmland.ru/")

	// Проверяем, что файл был создан
	fileName := "gzmland.ru.html"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Errorf("File %s was not created", fileName)
	}

	// Читаем содержимое файла
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("Failed to read file %s: %v", fileName, err)
	}

	// Читаем содержимое файла
	expectedContent, err := ioutil.ReadFile("gzmland.ru2.html")
	if err != nil {
		t.Errorf("Failed to read file %s: %v", fileName, err)
	}

	if string(content) != string(expectedContent) {
		t.Errorf("File content does not match expected. Got: %s, Expected: %s", string(content), string(expectedContent))
	}

	// Удаляем временный файл
	err = os.Remove(fileName)
	if err != nil {
		t.Errorf("Failed to remove file %s: %v", fileName, err)
	}
}