package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func FileOpen(path string) []byte {
	data, err := ioutil.ReadFile(fmt.Sprintf("data/%s.json", path))
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func req(service, path string) {
	resp, err := http.Post(fmt.Sprintf("http://%s:8000/execute", service), "application/json", bytes.NewBuffer(FileOpen(path)))
	if err != nil {
		log.Println(fmt.Sprintf("Сервис '%s' недоступен", service))
	}
	fmt.Println(fmt.Sprintf("%s - %d", service, resp.StatusCode))
}

func worker(service, path200, path400 string, lag, lag400 int) {
	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			req(service, path400)
			randInt := rand.Intn(30-0) + lag400
			ticker = time.NewTicker(time.Second * time.Duration(randInt))
		default:
			req(service, path200)
		}
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(lag-0) + 0
		time.Sleep(time.Millisecond * time.Duration(randInt))
	}
}

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)

	var (
		service1 = "metrics-service-1"
		service2 = "metrics-service-2"
		json200  = "metrics-script-200"
		json400  = "metrics-script-400"
	)

	go worker(service1, json200, json400, 2000, 25)
	go worker(service2, json200, json400, 3000, 10)
}
