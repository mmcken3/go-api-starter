package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mmcken3/go-api-starter/cmd/api/handler"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type cfg struct {
	Port  int  `envconfig:"PORT" default:"8002"`
	Debug bool `envconfig:"DEBUG" default:"false"`
}

var config cfg
var log *logrus.Logger

func init() {
	log = logrus.New()

	if err := godotenv.Load(); err != nil {
		log.Println(errors.Wrap(err, "no .env file loaded"))
	}

	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(errors.Wrap(err, "envconfig process"))
	}

	if !config.Debug {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
}

func main() {
	log.Println("Go API Starter is now running!")

	cert, err := loadCert()
	if err != nil {
		log.Fatal(errors.Wrap(err, "loading cert"))
	}

	// set up our global handler
	h, err := handler.New(handler.Config{
		Logger: log,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "handler new"))
	}

	server := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf(":%d", config.Port),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	// do graceful server shutdown
	go gracefulShutdown(server, time.Second*30)

	log.Infof("listening on port %d", config.Port)
	if err := server.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
		log.WithError(err).Fatal("cannot start a server")
	}
}

// gracefulShutdown shuts down our server in a graceful way.
func gracefulShutdown(server *http.Server, timeout time.Duration) {
	sigStop := make(chan os.Signal)

	signal.Notify(sigStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	<-sigStop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("graceful shutdown failed")
	}

	log.Info("graceful shutdown complete")
}
