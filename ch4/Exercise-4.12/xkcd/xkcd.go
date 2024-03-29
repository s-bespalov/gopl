// provides go api to xkcd.com
package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	JsonUrl  = "https://xkcd.com/%d/info.0.json"
	ComicUrl = "https://xkcd.com/%d"
	ComicDir = "comics"
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
func DownloadAll(comicCount int) (count int, result *[]Comic, e []error) {
	cs := make([]Comic, comicCount)
	e = make([]error, 0)
	result = &cs
	for i := range *result {
		url := fmt.Sprintf(JsonUrl, i+1)
		resp, err := http.Get(url)
		if err != nil {
			e = append(e, err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("download comics error: %s, URL: %s", resp.Status, url)
			e = append(e, err)
			continue
		}
		if err := json.NewDecoder(resp.Body).Decode(&cs[i]); err != nil {
			e = append(e, err)
			continue
		}
		count += 1
	}
	return
}

func Save(c *Comic, p string) error {
	data, err := json.MarshalIndent(c, " ", "")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(p, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func SaveAll(cs *[]Comic) error {
	err := os.RemoveAll(ComicDir)
	if err != nil {
		return err
	}
	err = os.Mkdir(ComicDir, 0777)
	if err != nil {
		return err
	}
	for _, c := range *cs {
		p := fmt.Sprintf("%s/%d.json", ComicDir, c.Num)
		if c.Num == 0 && c.Title == "" && c.Transcript == "" {
			continue
		}
		err = Save(&c, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadAll() (*[]Comic, error) {
	files, err := os.ReadDir(ComicDir)
	if err != nil {
		return nil, err
	}
	result := make([]Comic, 0, len(files))
	for _, e := range files {
		if e.Type().IsRegular() && strings.Contains(e.Name(), "json") {
			p := fmt.Sprintf("%s/%s", ComicDir, e.Name())
			data, err := ioutil.ReadFile(p)
			if err != nil {
				return nil, err
			}
			var c Comic
			err = json.Unmarshal(data, &c)
			if err != nil {
				return nil, err
			}
			result = append(result, c)
		}
	}
	return &result, nil
}

func Search(cs *[]Comic, num int, title, content string) *[]Comic {
	results := make([]Comic, 0)
	for _, c := range *cs {
		f := c.Num == num
		f = f || title != "" && strings.Contains(strings.ToLower(c.Title), title)
		f = f || content != "" && strings.Contains(strings.ToLower(c.Transcript), content)
		if f {
			results = append(results, c)
		}
	}
	return &results
}
