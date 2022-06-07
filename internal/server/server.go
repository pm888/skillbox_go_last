package server

import (
	"context"
	"fmt"
	"log"
	"mymod/internal/api"
	"mymod/internal/data"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Server() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./web")))
	mux.HandleFunc("/api/", handleConnection)

	srv := &http.Server{
		Handler: mux,
		Addr:    data.UrlServer}

	go func() {
		log.Printf("Listening on :%s...", data.UrlServer)
		if err := srv.ListenAndServe(); err != nil {
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
	get, err := api.GetApi()
	if err != nil {
		fmt.Println(err)
	}
	w.Write(get)

}
