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
	minWorker := *pFlag
	if len(args) < minWorker {
		minWorker = len(args)
	}

	reqChan := make(chan string, minWorker)
	resChan := make(chan string, minWorker)
	var wg sync.WaitGroup
	wg.Add(len(args))

	go startReceiver(resChan, &wg, log)

	for i := 1; i <= minWorker; i++ {
		go startWorker(i, reqChan, resChan, log)
	}

	for i := 0; i < len(args); i++ {
		reqChan <- args[i]
	}
	close(reqChan)

	wg.Wait()
	close(resChan)
	log.Println("work complete.")
}

//startReceiver receives data from all worker goroutine, and prints them
//if the url is invalid, it prints error line instead of md5 hash.
func startReceiver(resChan <-chan string, wg *sync.WaitGroup, log *log.Logger) {
	for res := range resChan {
		log.Println(res)
		wg.Done()
	}
}

//startWorker runs in a separate goroutine. Reads all incoming request from reqChan channel
//and try to fetch data using httpclient from the request address.
//On success, it creates md5 hash from response body. Otherwise, it will save the error
func startWorker(id int, reqChan <-chan string, resChan chan<- string, log *log.Logger) {
	for add := range reqChan {
		log.Printf("Go: %d. Address: %s\n", id, add)
		resp, err := fetch(add)
		if err != nil {
			resChan <- fmt.Sprintf("%s err: %s", add, err.Error())
			continue
		}

		hash := fmt.Sprintf("%x", md5.Sum(resp))
		resChan <- fmt.Sprintf("%s %s", add, hash)
	}
}

//fetch tries to do http Get request and read the response body.
func fetch(url string) ([]byte, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
