package server

import (
	"fmt"
	"net"
	"os"
)

type Connection struct {
	Port string
	Host string
	Type string
}

func (c Connection) StartServer() {
	l, err := net.Listen(c.Type, c.Host+":"+c.Port)

	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening on " + c.Host + ":" + c.Port)

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go c.HandleRequest(conn)
	}
}

func (c Connection) HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}

	conn.Write([]byte(buf))

	conn.Close()
}
