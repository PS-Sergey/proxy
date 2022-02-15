package api

func (api *Api) configureRouter() {
	api.router.HandleFunc("/", api.getProxy)
}
