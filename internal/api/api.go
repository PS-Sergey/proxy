package api

import (
	"fmt"
	"net/http"
)

type Api struct {
	config     *Config
	router     *http.ServeMux
	cache      map[string][]byte
	cacheRange []string
}

func NewApi(conf *Config) *Api {
	return &Api{
		config: conf,
		router: http.NewServeMux(),
		cache:  make(map[string][]byte),
	}
}

func (api *Api) Start() error {
	api.configureRouter()
	fmt.Println("Server start on port " + api.config.Port)
	return http.ListenAndServe(api.config.Port, api.router)
}
