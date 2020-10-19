package github2

import (
	"fmt"
	"log"
)

func GetRepository(owner, repo string) (*Repository, error) {
	url := fmt.Sprintf("%s/%s/%s", reposUrl, owner, repo)
	log.Println(url)
	return nil, nil
}
