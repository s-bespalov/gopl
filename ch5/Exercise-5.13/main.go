package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/s-bespalov/gopl/ch5/Exercise-5.13/links"
)

func main() {
	for _, arg := range os.Args[1:] {
		ok, err := downloadRoot(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		if ok {
			fmt.Printf("%s... Downloaded!\n", arg)
		}

		links := crawl(arg)
		for _, link := range links {
			ok, err = handleLink(link, arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error handling link %s: %v\n", link, err)
			} else if ok {
				fmt.Printf("%s... Downloaded!\n", link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func downloadRoot(url string) (bool, error) {
	dir, err := getDir(url)
	if err != nil {
		return false, err
	}
	f := fmt.Sprintf("%s/index.html", dir)
	err = download(url, f)
	if err != nil {
		return false, fmt.Errorf("downloading from %s to %s: %v", url, f, err)
	}
	return true, nil
}

func getDir(url string) (dir string, err error) {
	dir = strings.TrimPrefix(url, "http://")
	dir = strings.TrimPrefix(dir, "https://")
	slash := strings.Index(dir, "/")
	if slash >= 0 {
		dir = dir[:slash]
	}

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)
		if err != nil {
			err = fmt.Errorf("creating dir %q: %v", dir, err)
		}
	}
	return
}

func handleLink(url, root string) (bool, error) {
	if !strings.HasPrefix(url, root) {
		return false, nil
	}
	rdir, err := getDir(url)
	if err != nil {
		return false, err
	}

	fpath := strings.TrimPrefix(url, root)
	fpath = strings.TrimSuffix(fpath, "/")
	slash := strings.LastIndex(fpath, "/")
	if slash >= 0 {
		fpath = fpath[slash+1:]
	}

	q := strings.LastIndex(fpath, "?")
	if q >= 0 {
		fpath = fpath[:q]
	}

	fpath = fmt.Sprintf("%s/%s.html", rdir, fpath)
	err = download(url, fpath)
	if err != nil {
		return false, fmt.Errorf("downloading from %s to %s: %v", url, fpath, err)
	}

	return true, nil
}

func download(url, path string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		if err != nil {
			return fmt.Errorf("HTTP GET from %s: %v", url, err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("HTTP GET %s status code: %d", url, resp.StatusCode)
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("can't create file %s: %v", path, err)
	}

	body := bufio.NewReader(resp.Body)
	_, err = body.WriteTo(f)
	resp.Body.Close()
	f.Close()
	if err != nil {
		return fmt.Errorf("writing page %s to file %s: %v", url, path, err)
	}

	return nil
}
