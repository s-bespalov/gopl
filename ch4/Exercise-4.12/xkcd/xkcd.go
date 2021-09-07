// provides go api to xkcd.com
package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	JsonUrl    = "https://xkcd.com/%d/info.0.json"
	ComicUrl   = "https://xkcd.com/%d"
	ComicCount = 5
)

type Comic struct {
	Num        int
	Transcript string
	Title      string
	Alt        string
}

// DownloadAll downloads all comics from xkcd.com
// returns a number of downloaded comics,
// slice of Comic objects and error
func DownloadAll() (count int, result *[]Comic, e error) {
	cs := make([]Comic, ComicCount)
	result = &cs
	for i := range *result {
		url := fmt.Sprintf(JsonUrl, i+1)
		resp, err := http.Get(url)
		if err != nil {
			e = err
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			e = fmt.Errorf("download comics error: %s, URL: %s", resp.Status, url)
			return
		}
		if e = json.NewDecoder(resp.Body).Decode(&cs[i]); err != nil {
			return
		}
		count += 1
	}
	return
}


