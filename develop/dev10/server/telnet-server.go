package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Установка обработчика сигнала для завершения при нажатии Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Запуск telnet-сервера на порту 23
	listener, err := net.Listen("tcp", ":23")
	if err != nil {
		log.Println("Ошибка запуска сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()

	log.Println("Сервер запущен, ожидание подключений...")

	// Ожидание подключений от клиентов
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка подключения клиента:", err)
			continue
		}

		log.Printf("Новое подключение от %s\n", conn.RemoteAddr())
		go func() {
			select {
			case <-signalChan:
				log.Println("Server closing connection")
				os.Exit(1)
			}
		}()
		// Запуск горутины для чтения из сокета и вывода в STDOUT
		go handleConnection(conn)
	}

	
}

// Обработка подключения клиента
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Копирование данных из сокета в STDOUT
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		log.Println("Ошибка чтения из сокета:", err)
		return
	}
}