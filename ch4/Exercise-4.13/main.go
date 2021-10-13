package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.13/omdbapi"
)

func main() {
	flag.Parse()
	m := strings.Join(flag.Args(), " ")
	p, e := omdbapi.DownloadPoster(m)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Printf("Your poster: %s\n", p)
}
