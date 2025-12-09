package api

import (
	"log"
	"net/http"

	"github.com/irviner26/ecom/service/user"
	"github.com/jackc/pgx/v5"
	"github.com/julienschmidt/httprouter"
)

type APIServer struct {
	addr string
	db   *pgx.Conn
}

func NewAPIServer(addr string, db *pgx.Conn) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := httprouter.New()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	log.Println("Listening and serving on:", s.addr)

	return http.ListenAndServe(s.addr, router)
}
