package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	log := log.New(os.Stdout, "", 0)
	pFlag := flag.Int("parallel", 10, "parallel request limit")
	flag.Parse()
	args := flag.Args()

	reqChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(args))

	for i := 1; i <= *pFlag; i++ {
		go start(i, reqChan, &wg, log)
	}

	for i := 0; i < len(args); i++ {
		reqChan <- args[i]
	}

	wg.Wait()
	log.Println("work complete.")
}

func start(id int, reqChan <-chan string, wg *sync.WaitGroup, log *log.Logger) {
	for add := range reqChan {
		log.Printf("Go: %d. Address: %s\n", id, add)
		resp, err := fetch(add)
		if err != nil {
			log.Printf("%s ,err: %s\n", add, err.Error())
			wg.Done()
			continue
		}

		hash := fmt.Sprintf("%x", md5.Sum(resp))
		log.Printf("%s %s\n", add, hash)

		wg.Done()
	}
}

func fetch(url string) ([]byte, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
