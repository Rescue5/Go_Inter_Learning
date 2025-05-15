package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"pinger/extra"
	"strings"
)

func ping(url string, respCh chan string, errCh chan error) {
	res, err := http.Get(url)
	if err != nil {
		errCh <- extra.WrapError("Ошибка в GET запросе", err)
		return
	}

	respCh <- fmt.Sprintf("%v Status Code: %v", url, res.StatusCode)
}

func main() {
	path := flag.String("path", "urls.txt", "Путь ка файлу с urls")
	flag.Parse()

	file, err := os.ReadFile(*path)
	if err != nil {
		panic(extra.WrapError("Ошибка парсинга файла", err))
	}
	urlSlice := strings.Split(string(file), "\n")

	respCh := make(chan string)
	errCh := make(chan error)

	for _, url := range urlSlice {
		go ping(url, respCh, errCh)
	}

	for range urlSlice {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case res := <-respCh:
			fmt.Println(res)
		}
	}
}
