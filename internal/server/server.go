package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"mymod/internal/api"
	"mymod/internal/data"
)

func Server() error {
	r := http.NewServeMux()
	r.Handle("/", http.FileServer(http.Dir("./web")))
	r.HandleFunc("/api/", handleConnection)

	srv := &http.Server{
		Handler: r,
		Addr:    data.UrlServer,
	}

	go func() {
		log.Printf("Listening on :%s...", data.UrlServer)
		if err := srv.ListenAndServe(); err != nil {
			switch err {
			case http.ErrServerClosed:
				log.Println("server has been closed")
			default:
				log.Printf("ListenAndServe() error: %v", err)
			}
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ShutdownServer(srv)
	return nil
}

func ShutdownServer(srv *http.Server) {
	log.Println("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown server error: %v", err)
	}
	log.Println("shutdown server successfully")
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	get, err := api.GetApi()
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(get)
	if err != nil {
		log.Println(err)
	}
}
