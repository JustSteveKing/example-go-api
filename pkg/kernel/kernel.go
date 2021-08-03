package kernel

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JustSteveKing/example-go-api/pkg/config"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Application is our main application struct.
type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *config.Config
}

// Boot - Lets boot our Application with stuff.
func Boot() *Application {
	config := config.Load()
	router := mux.NewRouter()
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	logger, err := zap.NewDevelopment()

	if err != nil {
		panic(err)
	}

	return &Application{
		Server: &http.Server{
			Addr:         ":" + config.App.Port,
			Handler:      corsHandler(router),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Router: router,
		Logger: logger,
		Config: config,
	}
}

// Run - Lets run our Application server or panic.
func (app *Application) Run() {
	if err := app.Server.ListenAndServe(); err != nil {
		app.Logger.Fatal(err.Error())
		panic(err)
	}
}

// WaitForShutdown - Lets wait for a shutdown signal and shutdown gracefully
func (app *Application) WaitForShutdown() {
	// Create a channel to listen for OS signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal through our channel
	<-interruptChan

	app.Logger.Info("Received shutdown signal, gracefully terminating")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.Server.Shutdown(ctx)
	os.Exit(1)
}
