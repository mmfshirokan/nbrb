package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mmfshirokan/nbrb/internal/config"
	"github.com/mmfshirokan/nbrb/internal/consumer"
	"github.com/mmfshirokan/nbrb/internal/handlers"
	"github.com/mmfshirokan/nbrb/internal/repository"
	"github.com/mmfshirokan/nbrb/internal/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cnf, err := config.New()
	if err != nil {
		log.Fatal("Fatal parse config error: ", err)
	}

	db, err := sql.Open("mysql", cnf.MysqlURL)
	if err != nil {
		log.Fatal("Fatal MySQl connection error: ", err)
	}
	defer db.Close()

	r := repository.New(db)
	s := service.New(r)
	h := handlers.New(s)
	c := consumer.New(s)

	go c.Consume(ctx, cnf.SourceURL)

	http.HandleFunc("/get", h.Get)
	http.HandleFunc("/getall", h.GetAll)

	server := &http.Server{
		Addr:    cnf.ServerPort,
		Handler: http.DefaultServeMux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Fatal server error: ", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	go func() {
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ch

		cancel()
		if err := server.Shutdown(nil); err != nil {
			log.Fatal("Fatal server shutdown error: ", err)
		}
	}()

	wg.Wait()
	log.Info("Server gracefully stopped")
}
