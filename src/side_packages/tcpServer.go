package side_packages


import (
	"encoding/gob"
	"fmt"
	"net"
)


func server() {
	// сервер слушает порт 9999
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// если ошибок нет - принимается соединение с клиентом
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// обработка соединения (для удобства - в отдельном потоке)
		go handleServerConnection(c)
	}
}

// Функция получения и расшифровки сообщения
func handleServerConnection(c net.Conn) {
	// получение сообщения
	var msg string
	// важно, что здесь подключение используется как аргумент
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	// подключение к серверу по соответствующему порту
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// отправка сообщения
	msg := "Hello World"
	fmt.Println("Sending", msg)
	// важно, что здесь подключение используется как аргумент
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func TCPServer() {
	// для удобства сервер и клиент работают в разных потоках
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}