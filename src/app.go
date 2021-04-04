package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Hofsiedge/Organizer/src/tasktracker"
	"github.com/Hofsiedge/Organizer/src/tasktracker/middleware"
	"github.com/Hofsiedge/Organizer/src/tasktracker/repository/postgres"
	transport "github.com/Hofsiedge/Organizer/src/tasktracker/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	// Logging
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			TimeKey:        "time",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}.Build()
	defer logger.Sync()

	logger.Info("logger initialized",
		zap.String("logger", "go.uber.org/zap"),
	)

	// DB connection
	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %w\n", err)
		logger.Error(err.Error())
		panic(err)
	}
	defer dbPool.Close()

	logger.Info("DB connected")

	// Task tracker service
	var svc tasktracker.IService
	svc = tasktracker.Service{
		Repository: postgres.TaskRepository{
			Db: dbPool,
		},
	}
	svc = middleware.LoggingMiddleware{
		Logger: logger,
		Next:   svc,
	}

	// handlers
	createTaskHandler := httptransport.NewServer(
		transport.MakeCreateTaskEndpoint(svc),
		transport.DecodeCreateTaskRequest,
		transport.EncodeJSONResponse,
	)

	getTaskHandler := httptransport.NewServer(
		transport.MakeGetTaskEndpoint(svc),
		transport.DecodeGetTaskRequest,
		transport.EncodeJSONResponse,
	)

	r := mux.NewRouter()

	r.Handle("/task/", createTaskHandler).Methods("POST")
	r.Handle("/task/{id}", getTaskHandler).Methods("GET")

	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m",
	)
	flag.Parse()

	srv := &http.Server{
		Addr: "0.0.0.0:80",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 150,
		ReadTimeout:  time.Second * 150,
		IdleTimeout:  time.Second * 600,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint // We received an interrupt signal, shut down.
		logger.Info("shutting down")

		ctx, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Error(fmt.Sprintf("HTTP server shutdown %v", err))
		}
		<-ctx.Done()

		close(idleConnsClosed)
	}()

	logger.Info("Starting listening")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error(fmt.Sprintf("HTTP server ListenAndServe: %v", err))
	}

	<-idleConnsClosed
}
