package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type response struct {
	address string
	hash    string
	err     error
}

const defaultParallel = 10
const maxParallel = 1000 //for now, added 1000 as max concurrent/parallel request limit.
const defaultTimeInSec = 5

func main() {
	log := log.New(os.Stdout, "", 0)
	pFlag := flag.Int("parallel", defaultParallel, "parallel request limit")
	flag.Parse()
	args := flag.Args()

	//minworker represents minimum required worker/goroutine to start.
	minWorker := *pFlag
	if minWorker < 0 {
		minWorker = defaultParallel
	} else if minWorker > maxParallel {
		minWorker = maxParallel
	}

	if len(args) < minWorker {
		minWorker = len(args)
	}

	reqChan := make(chan string, minWorker)
	resChan := make(chan response, minWorker)
	var wg sync.WaitGroup
	wg.Add(len(args))

	//start receiver routine.
	go startReceiver(resChan, &wg, log)

	//initialize minimum number of worker goroutine to fetch http requests & calculate md5 hash
	for i := 1; i <= minWorker; i++ {
		go startWorker(reqChan, resChan)
	}

	//send all arguments to reqChan, worker goroutines will read from this channel
	for i := 0; i < len(args); i++ {
		reqChan <- args[i]
	}
	close(reqChan) //no more req to send, close the channel

	wg.Wait()      //wait for receiver channel to complete
	close(resChan) //receiver channel processed all data.
}

//startReceiver receives data from all worker goroutine, and prints them
//if the url is invalid, it prints error line instead of md5 hash.
func startReceiver(resChan <-chan response, wg *sync.WaitGroup, log *log.Logger) {
	for res := range resChan {
		log.Println(res.string())
		wg.Done()
	}
}

//startWorker runs in a separate goroutine. Reads all incoming request from reqChan channel
//and try to fetch data using httpclient from the request address.
//On success, it creates md5 hash from response body. Otherwise, it will save the error
func startWorker(reqChan <-chan string, resChan chan<- response) {
	client := http.Client{Timeout: defaultTimeInSec * time.Second}
	for req := range reqChan {
		add, err := getFormattedAddress(req)
		if err != nil {
			resChan <- response{address: req, err: err}
			continue
		}
		resp, err := fetch(add, &client)
		if err != nil {
			resChan <- response{address: add, err: err}
			continue
		}

		hash := fmt.Sprintf("%x", md5.Sum(resp))
		resChan <- response{address: add, hash: hash}
	}
}

//fetch tries to do http Get request and read the response body.
func fetch(url string, client *http.Client) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

//getFormattedAddress returns formatted address if url scheme is missing
func getFormattedAddress(address string) (string, error) {
	u, err := url.Parse(address)
	if err != nil {
		return "", err
	}
	if !u.IsAbs() {
		u.Scheme = "http"
	}
	return u.String(), nil
}

func (r *response) string() string {
	if r.err != nil {
		return fmt.Sprintf("%s err: %s", r.address, r.err.Error())
	}
	return fmt.Sprintf("%s %s", r.address, r.hash)
}
