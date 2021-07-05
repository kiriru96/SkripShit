package client

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	Host string
	Port string
	Type string
}

func (c Client) SendRequest(reqHeader string) {
	conn, err := net.Dial(c.Type, c.Host+":"+c.Port)

	if err != nil {
		fmt.Println("Request Error: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

}

func (c Client) ReadCorrection(paperid string) {

}
