package sockets

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// RunSock метод инициализации сокета
func RunSock() {
	ln, err := net.Listen("tcp", os.Getenv("IP")+":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Failed listen socket")
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal("Failed accept socket")
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("Failed accept socket")
		}

		fmt.Print("Message received:", string(message))

		response := strings.ToUpper(message)

		conn.Write([]byte(response + "\n"))
	}
}

// ConnectSock метод подключения к сокету
func ConnectSock() {
	conn, err := net.Dial("tcp", os.Getenv("IP")+":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Failed dial socket", err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed read text.")
		}

		fmt.Fprintf(conn, text)

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("Failed read text.")
		}

		fmt.Print("Message from server: " + message)
	}
}
