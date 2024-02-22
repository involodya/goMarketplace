package handler

import (
	"context"
	"flag"
	"fullstack/backend/internal/pkg/auth"
	"fullstack/backend/internal/repository"
	reposqlite "fullstack/backend/internal/repository/sqlite"
	"fullstack/backend/internal/service"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer(url, dbName, signingKey string) *http.Server {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	logFile, err := os.OpenFile("log/backend.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	srv := InitializeServer(url, dbName, signingKey)

	StartPolling(srv)
	defer StopPolling(srv)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	return srv
}

func InitializeServer(url, dbName, defaultSigningKey string) *http.Server {
	db, err := reposqlite.NewSQLiteDB(dbName)
	if err != nil {
		log.Panicf("Failed to initialize database: %s\n", err.Error())
	} else {
		log.Println("Database is initialized")
	}

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)

	signingKey, ok := os.LookupEnv("AUTH_SIGNING_KEY")
	if !ok {
		log.Println("cannot get AUTH_SIGNING_KEY from ENV")
		signingKey = defaultSigningKey
	}
	authManager := auth.NewAuthManager([]byte(signingKey))

	h := NewHandler(serv, authManager)

	srv := &http.Server{
		Addr: url,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.NewRouter(), // Pass our instance of gorilla/mux in.
	}
	return srv
}

func StartPolling(srv *http.Server) {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

func StopPolling(srv *http.Server) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	_ = srv.Shutdown(ctx)
	log.Println("shutting down")
}
