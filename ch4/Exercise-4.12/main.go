package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.12/xkcd"
)

func main() {
	var comics *[]xkcd.Comic
	var err error

	flag.Parse()

	switch flag.Arg(0) {
	case "download":
		_, comics, err = xkcd.DownloadAll(5)
		check(err)
		err = xkcd.SaveAll(comics)
		check(err)
		fmt.Printf("%d comics downloaded and saved\n", len(*comics))
	case "show":
		comics, err = xkcd.ReadAll()
		check(err)
		for _, c := range *comics {
			fmt.Printf("%d\t%s\n", c.Num, c.Title)
		}
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
