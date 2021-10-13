package omdbapi

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	keyVar = "OMDBAPI_KEY"
	apiUrl = "http://www.omdbapi.com/"
)

type Movie struct {
	poster string
}

func DownloadPoster(name string) (path string, e error) {
	vs := make(url.Values, 0)
	key, ok := os.LookupEnv(keyVar)
	if !ok {
		e = fmt.Errorf("download poster: %s enviroment variable not set", keyVar)
		return
	}
	vs.Add("apikey", key)
	name = url.QueryEscape(name)
	vs.Add("t", name)
	u := fmt.Sprintf("%s?%s", apiUrl, vs.Encode())
	resp, err := http.Get(u)
	if err != nil {
		e = fmt.Errorf("download poster: %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		e = fmt.Errorf("download poster: bad http status: %s", resp.Status)
	}
	return
}
