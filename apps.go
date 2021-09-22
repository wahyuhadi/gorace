package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/gorace/models"
	"github.com/wahyuhadi/gorace/parser"
)

var file = flag.String("f", "", "http request file ")
var treq = flag.Int64("tr", 100, "total  request  ")
var stcode = flag.Int("stcode", 200, "Expect status code")

var client = &http.Client{}

const (
	fileNotFond = "File not found"
)

func ParseOptions() (opts *models.Options) {
	flag.Parse()
	return &models.Options{
		File:     *file,
		TotalReq: *treq,
	}
}

func CheckOptions(opts *models.Options) (err error) {
	if *file == "" {
		return errors.New(fileNotFond)
	}
	return nil
}

func main() {
	opts := ParseOptions()
	err := CheckOptions(opts)
	if err != nil {
		gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
		return
	}

	f, err := os.Open(opts.File)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	stream, err := parser.ReadHTTPFromFile(f)
	if err != nil {
		log.Fatalln(err)
	}

	host := stream.Request.Host
	url := stream.Request.URL
	gologger.Info().Str("state", "request").Str("URL", fmt.Sprintf("%v", url)).Msg("URL dest")
	gologger.Info().Str("state", "host").Str("URL", fmt.Sprintf("%v", host)).Msg("HOST dest")
	gologger.Info().Str("state", "proto").Str("URL", fmt.Sprintf("%v", stream.Request.Proto)).Msg("Proto dest")

	UriNew := fmt.Sprintf("http://%v%v", host, url)
	stream.Request.RequestURI = ""

	// Since the req.URL will not have all the information set,
	// such as protocol scheme and host, we create a new URL
	u, err := url.Parse(UriNew)
	if err != nil {
		panic(err)
	}
	stream.Request.URL = u
	ok := 0
	sendRequest := func() {
		response, err := client.Do(stream.Request)
		if err != nil {
			gologger.Error().Str("Error", "State").Str("Message", fmt.Sprintf("%v", err)).Msg("Error request")
			return
		}
		if response.StatusCode == *stcode {
			ok = ok + 1
		}
	}
	start := time.Now()
	var wg sync.WaitGroup
	ch := make(chan int, int(opts.TotalReq))
	for i := 0; i < int(opts.TotalReq); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sendRequest()
			ch <- n
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	gologger.Info().Str("state", "sequence").Str("start", start.Format("2006-02-01")).Msg("Start at")

	for routineNumber := range ch {
		gologger.Info().Str("state", "sequence").Str("completion sequence number", fmt.Sprintf("%v", routineNumber)).Msg("routine number")
	}
	gologger.Info().Str("state", "sequence").Str("finish", fmt.Sprintf("%v", time.Since(start))).Msg("Finish at")
	gologger.Info().Str("state", "success").Str("OK", fmt.Sprintf("%v", ok)).Msg("Expect http status code in responses ")

}
