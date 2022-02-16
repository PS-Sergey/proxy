package api

import "github.com/PS-Sergey/proxy/cache"

func (api *Api) configureRouter() {
	api.router.HandleFunc("/", api.getProxy)
}

func (api *Api) configureCache() {
	api.cache = cache.NewCache(api.config.CacheSize)
}
