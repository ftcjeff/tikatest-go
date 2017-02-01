package main

import (
	"flag"
	"log"

	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"io"
)

var tikaUrlOpt = flag.String("tika", "http://localhost:9998", "The URI to the Tika Server")
var fileToProcessOpt = flag.String("file", "", "The file to process")
var numIterationOpt = flag.Int("count", 1000, "The number of times to process the given file")

func main() {
	flag.Parse()

	log.Printf("Processing %s %d times\n", *fileToProcessOpt, *numIterationOpt)

	uri := *tikaUrlOpt + "/meta"

	buf, err := ioutil.ReadFile(*fileToProcessOpt)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client {}

	bufReader := strings.NewReader(string(buf))

	for i := 0; i < *numIterationOpt; i++ {
		bufReader.Seek(0, io.SeekStart)
		req, err := http.NewRequest("PUT", uri, bufReader)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Content-type", "application/octet-stream")
		req.Header.Add("Accept", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		contents, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		res.Body.Close()

		var data interface{}
		if err := json.Unmarshal(contents, &data); err != nil {
			log.Fatal(err)
		}

//		log.Printf("%v\n", data)
	}
}