package main

import (
	"fmt"
	"log"
	"net/http"

	"a21hc3NpZ25tZW50/controllers/inbound"
	"a21hc3NpZ25tZW50/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	_, err := models.ConnectDatabase()

	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/inbound/list", inbound.ReadInbound).Methods("GET")
	r.HandleFunc("/v1/inbound/list/{id}", inbound.ShowInboundId).Methods("GET")
	r.HandleFunc("/v1/inbound/add", inbound.CreateInbound).Methods("POST")
	r.HandleFunc("/v1/inbound/update/{id}", inbound.UpdateInboundId).Methods("PUT")
	r.HandleFunc("/v1/inbound/delete/{id}", inbound.DeleteInboundId).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(r)
	fmt.Println("starting web server at http://localhost:5010")
	log.Fatal(http.ListenAndServe(":5010", handler))
}

type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "*" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
