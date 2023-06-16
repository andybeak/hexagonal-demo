package handlers

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type HttpServiceParams struct {
	port    string
	address string
}

func ProvideHttpServiceParams() *HttpServiceParams {
	return &HttpServiceParams{
		port:    os.Getenv("INTERNAL_HTTP_PORT"),
		address: os.Getenv("INTERNAL_HTTP_ADDRESS"),
	}
}

func ProvideHttpService(
	params *HttpServiceParams,
	r *mux.Router,
) *HttpService {
	return &HttpService{
		params: *params,
		router: r,
	}
}

type HttpService struct {
	params HttpServiceParams
	router *mux.Router
}

func (h *HttpService) StartHttpServer() {
	handler := handlers.CombinedLoggingHandler(os.Stdout, h.router)
	handler = handlers.RecoveryHandler()(handler)
	handler = handlers.CompressHandler(handler)
	handler = handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Authorization", "Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(handler)
	var listenAddr string
	if h.params.address == "" {
		listenAddr = ":"
	} else {
		listenAddr = h.params.address
	}
	listenAddr += h.params.port
	log.Printf("Starting http listener at [%s]...", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, handler))
}
