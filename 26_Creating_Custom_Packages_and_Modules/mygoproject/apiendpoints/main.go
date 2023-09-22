package apiendpoints

import (
	"fmt"
	"net/http"
)

type API struct {
	mux *http.ServeMux
}

func NewAPI() *API {
	return &API{mux: http.NewServeMux()}
}

func (api *API) Handle(pattern string, handler http.HandlerFunc) {
	api.mux.HandleFunc(pattern, handler)
}

func (api *API) StartServer() {
	fmt.Println("API server started on port 8080")
	http.ListenAndServe(":8080", api.mux)
}
