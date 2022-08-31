package goroutines

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Task2GracefulShutdown() {
	currentValue := 1
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
mainLoop:
	for {
		select {
		case <-exitChannel:
			fmt.Println("выхожу из программы")
			break mainLoop
		default:
			returnSquares(currentValue)
			currentValue++
		}
	}

}

func returnSquares(digit int) {
	fmt.Printf("%v^2=%v\n", digit, digit*digit)
	time.Sleep(1 * time.Second)
}
