package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("I am a client application")

	net, err := net.Dial("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		fmt.Print("Please enter some text: ")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(text)

		if _, err = net.Write([]byte(text)); err != nil {
			log.Fatalln(err)
		}

		uppercaseStr, err := bufio.NewReader(net).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(uppercaseStr))
	}
}
