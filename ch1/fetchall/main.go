// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// func main() {
// 	start := time.Now()
// 	ch := make(chan string)
// 	for _, url := range os.Args[1:] {
// 		go fetch(url, ch) // start a goroutine
// 	}
// 	for range os.Args[1:] {
// 		fmt.Println(<-ch) // receive from channel ch
// 	}
// 	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
// }

// func fetch(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err) // send to channel ch
// 		return
// 	}
// 	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
// 	resp.Body.Close() // dont leak resources
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
// 		return
// 	}
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
// }

// 1.10 exercise
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	// set time format for file naming in case of usage
	// format := "2006-01-02T15:04:05"
	// timeNow := start.Format(format)
	// open file
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for range os.Args[1:] {
		output := <-ch
		fmt.Println(output) // receive from channel ch
		n, err := f.WriteString(output + "\n")
		if err != nil {
			fmt.Println(n, err)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // dont leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
