package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	// Создаем каналы для тестирования
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	// Запускаем функцию `or` в отдельной горутине и получаем объединенный канал
	combined := or(ch1, ch2, ch3)

	// Записываем значения в каждый из каналов
	go func() {
		ch1 <- "channel 1"
	}()

	go func() {
		ch2 <- "channel 2"
	}()

	go func() {
		ch3 <- "channel 3"
	}()

	// Ждем получения значения из объединенного канала
	select {
	case  <-combined:
		t.Log("ok")
	case <-time.After(10 * time.Second):
		t.Error("Timed out waiting for message")
	}
}

func TestMain(t *testing.T) {
	// Запускаем функцию `main` в отдельной горутине
	go main()

	// Ждем завершения выполнения функции `main`
	time.Sleep(3 * time.Second)
}