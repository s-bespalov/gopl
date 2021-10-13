// provides omdbapi to download a movie poster
package omdbapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	keyVar = "OMDBAPI_KEY"
	apiUrl = "http://www.omdbapi.com/"
)

type Movie struct {
	Poster string
}

func DownloadPoster(name string) (path string, e error) {

	// search for movie
	vs := make(url.Values)
	key, ok := os.LookupEnv(keyVar)
	if !ok {
		e = fmt.Errorf("download poster: %s enviroment variable not set", keyVar)
		return
	}
	vs.Add("apikey", key)
	name = url.QueryEscape(name)
	vs.Add("t", name)
	vs.Add("type", "movie")
	u := fmt.Sprintf("%s?%s", apiUrl, vs.Encode())
	resp, err := http.Get(u)
	if err != nil {
		e = fmt.Errorf("download poster: %v", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		e = fmt.Errorf("download poster: bad http status: %s", resp.Status)
		return
	}
	defer resp.Body.Close()
	var m Movie
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		e = fmt.Errorf("download poster: %v", err)
		return
	}
	resp.Body.Close()

	// download image
	resp, err = http.Get(m.Poster)
	if err != nil {
		e = fmt.Errorf("download poster: %v", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		e = fmt.Errorf("download poster: bad http status: %s", resp.Status)
		return
	}
	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e = fmt.Errorf("download poster: %v", err)
		return
	}

	// save to file
	path = fmt.Sprintf("%d_%s.jpg", time.Now().Unix(), path)
	err = ioutil.WriteFile(path, img, 0644)
	if err != nil {
		e = fmt.Errorf("download poster: %v", err)
		return
	}
	return
}
