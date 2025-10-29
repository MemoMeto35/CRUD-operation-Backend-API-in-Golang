package api
import (
    "database/sql"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/MemoMeto35/CRUD-operation-Backend-API-in-Golang/service/user"
)

type APIserver struct {
    addr string
    db   *sql.DB
}

func NewAPIserver(addr string, db *sql.DB) *APIserver {
    return &APIserver{
        addr: addr,
        db:   db,
    }
}

func (s *APIserver) Run() error {
    router := mux.NewRouter()
    subrouter := router.PathPrefix("/api/v1").Subrouter()

    userHandler := user.NewHandler()
    userHandler.RegisterRoutes(subrouter)

    log.Println("Listen On", s.addr)
    return http.ListenAndServe(s.addr, router)
}
