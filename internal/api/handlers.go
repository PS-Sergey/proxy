package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (api *Api) getProxy(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("len of cache is", len(api.cache))
	url := fmt.Sprintf("http://%s%s", api.config.ProxyUrl, req.RequestURI)
	cachedResp, ok := api.cache[url]
	if ok {
		writer.Write(cachedResp)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error during getting response from server", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error during read response from server", err)
	}
	cacheSize, _ := strconv.Atoi(api.config.CacheSize)
	if len(api.cacheRange) == cacheSize {
		theOldestUrl := api.cacheRange[len(api.cacheRange)-1]
		delete(api.cache, theOldestUrl)
		api.cacheRange = api.cacheRange[:len(api.cacheRange)-1]
	}
	api.cache[url] = body
	api.cacheRange = append(api.cacheRange, url)
	writer.Write(body)
}
