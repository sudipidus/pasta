package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
)

var clipboardHistory []string

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go monitorClipboard(&wg)
	defer wg.Wait()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGUSR1)

	for {
		<-signalChan
		fmt.Println(strings.Join(clipboardHistory, ","))
	}

}

func monitorClipboard(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		content, err := clipboard.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		if len(clipboardHistory) == 0 {
			clipboardHistory = append(clipboardHistory, content)
		}
		if len(clipboardHistory) > 0 && clipboardHistory[len(clipboardHistory)-1] != content {
			clipboardHistory = append(clipboardHistory, content)
		}
		time.Sleep(1 * time.Second)
	}

}
