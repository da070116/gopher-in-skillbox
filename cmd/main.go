package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	gopherinskillbox "skillbox-test"
	"skillbox-test/pkg"
	"skillbox-test/pkg/handler"
	"skillbox-test/pkg/repository"
	"skillbox-test/pkg/service"
	"syscall"
)

// main - entry point
func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error in config load: %s", err.Error())
	}

	// get args from console to define port and host
	port := flag.Int("port", viper.GetInt("port"), "port to launch server")
	host := flag.String("host", "localhost", "host to launch server")
	flag.Parse()

	db, err := repository.NewSqliteDB(repository.Config{
		Name: viper.GetString("db_name"),
	})

	if err != nil {
		logrus.Fatalf("Failed to init database: %s", err.Error())
	}
	defer pkg.CloseDB(db)

	server := new(gopherinskillbox.Server)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go func() {
		err = server.Run(handlers.InitRoutes(), *host, *port)
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error on server start:%s", err.Error())
			return
		}
	}()

	logrus.Printf("Server on %s and port %d started", *host, *port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit

	err = server.Shutdown(context.Background())
	if err != nil {
		logrus.Fatalf("Error during shutdown: %s", err.Error())
	}
}

// initConfig - use configuration files
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
