package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/irviner26/ecom/service/user"
	"github.com/julienschmidt/httprouter"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := httprouter.New()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(router)

	log.Println("Listening and serving on:", s.addr)

	return http.ListenAndServe(s.addr, router)
}