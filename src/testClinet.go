package main

import (
	"fmt"
	"net"
	"os"

	"bufio"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:11110")
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}
	var input string

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input,_, err = inputReader.ReadLine()

	conn.Write([]byte(input))

	fmt.Println("send msg")

	var msg [20]byte
	conn.Read(msg[0:])

	fmt.Println("msg is", string(msg[0:10]))
}
