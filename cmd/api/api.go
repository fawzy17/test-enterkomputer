package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fawzy17/test-enterkomputer/service/order"
	"github.com/fawzy17/test-enterkomputer/service/product"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer  {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/example.com/api/v1").Subrouter()

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	orderHandler := order.NewHandler(orderStore, productStore)
	orderHandler.RegisterRoutes(subrouter)

	log.Println("Listening on:", s.addr)

	return http.ListenAndServe(s.addr, router)
}