package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("server is running")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		reader := bufio.NewReader(conn)
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(str)
		uppercaseStr := strings.ToUpper(str)
		fmt.Printf("Your text: %v\n", uppercaseStr)
		if _, err := conn.Write([]byte(uppercaseStr)); err != nil {
			log.Fatalln(err)
		}
	}
}
