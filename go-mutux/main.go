package main

import (
	"fmt"
	"os"
	"time"
)

const (
	logFile = "./log2.txt"
)

func main() {

	mutex := make(chan bool, 1)
	f, _ := os.Create(logFile)
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile(logFile, os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " " + msg)
				defer f.Close()
			} else {
				break
			}
		}
	}()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}()
		}
	}
	fmt.Scanln()
}
