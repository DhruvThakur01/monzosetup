package main

import (
	"runtime"
	"bufio"
	"fmt"
	"net/http"
	"time"
	"os"
	"strconv"
)

var exit = make(chan bool)
/*
var client = http.Client{
    Timeout: 2 * time.Second,
}
*/


func makeRequest(url string, isEgress bool, nThreads int) {

	befReq := time.Now()

	res, err := http.Get(url)

	// res, err := http.Get("http://20.207.64.199/")

	aftRes := time.Now()
	if err != nil {
		fmt.Printf("Error 1 | IsEgress: %t;Response Time: %v; Error: %s\n", isEgress, aftRes.Sub(befReq).Nanoseconds(),err.Error())
		runtime.GC()
		return
	}
	defer res.Body.Close()

    scanner := bufio.NewScanner(res.Body)
    for i := 0; scanner.Scan() && i < 1; i++ {
		fmt.Printf("Success | IsEgress: %t;Response Time: %v;nThreads/Pod: %d;Response Status: %s;Response Body: %s\n", isEgress, aftRes.Sub(befReq).Nanoseconds(), nThreads, res.Status, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
		panic(err)
    }

	runtime.GC()
}

func makeMultipleRequests(url string, isEgress bool, nThreads int) {
	for {
		makeRequest(url, isEgress, nThreads)
	}

	exit<-true
}

func main ()  {
	//fmt.Println("hello")

	var url string
	var isEgress bool

	var nThreads, i int

	url = os.Args[1]
	secondArg := os.Args[2]
	nThreads, e := strconv.Atoi(os.Args[3])

	if(e != nil) {}

	if (secondArg == "true") {
		isEgress = true
	} else {
		isEgress = false
	}

	for i = 0; i < nThreads; i++ {
		go makeMultipleRequests(url, isEgress, nThreads)
	}

	<-exit

}
