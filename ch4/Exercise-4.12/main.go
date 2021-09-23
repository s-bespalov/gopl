package main

import (
	"fmt"
	"log"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.12/xkcd"
)

func main() {
	_, comics, err := xkcd.DownloadAll(5)
	check(err)
	for _, c := range *comics {
		fmt.Printf("%d\t%s\n", c.Num, c.Title)
	}
	err = xkcd.SaveAll(comics)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
