package handler

import (
	"net/http"

	v2 "github.com/aniket-gupta/redis-cache-example/cmd/handler/v2"
)

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tracking", tracking)
	mux.HandleFunc("/search", search)

	//V2
	mux.HandleFunc("/v2/search", v2.SearchV2)
	mux.HandleFunc("/v2/cancel", v2.CancelRequest)
	return mux
}
