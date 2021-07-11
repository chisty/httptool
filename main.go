package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	log := log.New(os.Stdout, "httptool", log.LstdFlags|log.Lshortfile)
	pFlag := flag.Int("parallel", 10, "parallel request limit")
	flag.Parse()

	log.Printf("pFlag= %d, args: %s", *pFlag, flag.Args())

	reqChan := make(chan string)

	for i := 1; i <= *pFlag; i++ {
		go start(i, reqChan, log)
	}

	for i := 0; i < len(flag.Args()); i++ {
		reqChan <- flag.Args()[i]
	}

	log.Println("work complete.")
}

func start(id int, reqChan <-chan string, log *log.Logger) {
	for add := range reqChan {
		log.Printf("Go: %d. Address: %s\n", id, add)
	}
}
