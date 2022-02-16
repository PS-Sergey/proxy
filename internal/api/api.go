package api

import (
	"fmt"
	"net/http"

	"github.com/PS-Sergey/proxy/cache"
)

type Api struct {
	config *Config
	router *http.ServeMux
	cache  *cache.Cache
}

func NewApi(conf *Config) *Api {
	return &Api{
		config: conf,
		router: http.NewServeMux(),
	}
}

func (api *Api) Start() error {
	api.configureRouter()
	api.configureCache()
	fmt.Println("Server start on port " + api.config.Port)
	return http.ListenAndServe(api.config.Port, api.router)
}
