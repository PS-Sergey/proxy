package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (api *Api) getProxy(writer http.ResponseWriter, req *http.Request) {
	url := fmt.Sprintf("http://%s%s", api.config.ProxyUrl, req.RequestURI)
	cachedResp, ok := api.cache.GetValue(url)
	if ok {
		writer.Write(cachedResp)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error during getting response from server", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error during read response from server", err)
		return
	}
	api.cache.SetValue(url, body)
	writer.Write(body)
}
