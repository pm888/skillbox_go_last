package server

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"mymod/internal/api"
	"mymod/internal/data"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Printf("starting server...")
		if err := http.ListenAndServe(data.UrlServer, r); err != nil {
			switch err {
			case http.ErrServerClosed:
				log.Printf("server has been closed")
			default:
				log.Printf("ListenAndServe() error: %v", err)
			}
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ShutdownServer(srv)
}

func ShutdownServer(srv *http.Server) {
	log.Printf("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown server error", err)
	}
	log.Println("shutdown server successfully")
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(api.GetApi())
}
