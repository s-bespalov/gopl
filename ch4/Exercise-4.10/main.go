package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("less than a month old:")
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours() < 720 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("less than a year old:")
	for _, item := range result.Items {
		sc := time.Since(item.CreatedAt).Hours()
		if sc >= 720 && sc < 8760 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("more than a year old:")
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours() >= 8760 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
