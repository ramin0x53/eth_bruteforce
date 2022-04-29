package main

import (
	"flag"
	"fmt"

	"github.com/ramin0x53/eth_bruteforce/config"
	"github.com/ramin0x53/eth_bruteforce/file"
	"github.com/ramin0x53/eth_bruteforce/key"
)

func options() {
	flag.StringVar(&config.Filename, "f", "", "wordlist")
	flag.StringVar(&config.Apikey, "api", "", "api key")
	flag.IntVar(&config.Threadnum, "t", 50, "threads number")
	flag.Parse()
	fmt.Println(config.Filename)
	fmt.Println(config.Apikey)
	fmt.Println(config.Threadnum)
}

func checkstring(line string) {
	sha := key.ShaConvert(line)
	address := key.AddrGenerator(sha)
	balance := key.GetBalance(address)

	if balance != 0 {
		fmt.Println("String: ", line)
		fmt.Printf("Private key: %x\n", sha)
		fmt.Printf("Address: %x\n", address)
		fmt.Println("Balance: ", balance)
		fmt.Println("-------------------------------------------------")
	}
}

func worker(txt <-chan string, results chan<- bool) {
	for j := range txt {
		checkstring(j)
		results <- true
	}
}

func main() {
	options()
	a := file.Readfile(config.Filename)
	results := make(chan bool, len(a))
	jobs := make(chan string, len(a))

	for w := 1; w <= config.Threadnum; w++ {
		go worker(jobs, results)
	}

	for _, i := range a {
		jobs <- i
	}

	close(jobs)

	for s := 1; s <= len(a); s++ {
		<-results
	}
}
