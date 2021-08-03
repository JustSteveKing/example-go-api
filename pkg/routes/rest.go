package routes

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/JustSteveKing/example-go-api/pkg/database"
	"github.com/JustSteveKing/example-go-api/pkg/kernel"
	"github.com/JustSteveKing/example-go-api/pkg/services/photos"
)

func Load(app *kernel.Application) {
	serviceRouter := app.Router.Methods(http.MethodGet).Subrouter()

	serviceRouter.HandleFunc("/", HandleApiRoot(app)).Name("api:root")

	serviceRouter.HandleFunc("/clients", HandleGetClients(app)).Name("api:clients:index")

	serviceRouter.HandleFunc("/photos", HandlePhotosExternal(app)).Name("api:external:photos")
}

func HandlePhotosExternal(app *kernel.Application) http.HandlerFunc {
	service := photos.NewService()

	return func(wr http.ResponseWriter, request *http.Request) {
		var (
			waitGroup sync.WaitGroup
			response  *http.Response
			apiError  error
		)

		waitGroup.Add(1)
		go func() {
			response, apiError = service.Get(
				"https://jsonplaceholder.typicode.com/photos",
				&waitGroup,
			)
		}()
		waitGroup.Wait()

		if apiError != nil {
			app.Logger.Fatal(apiError.Error())
			panic(apiError)
		}

		fmt.Println(response)
	}
}

func HandleGetClients(app *kernel.Application) http.HandlerFunc {
	db := database.Connect(app)

	return func(wr http.ResponseWriter, request *http.Request) {
		var clients []database.Client

		db.Find(&clients)

		app.Respond(
			wr,
			request,
			clients,
			http.StatusOK,
		)
	}
}

func HandleApiRoot(app *kernel.Application) http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(wr http.ResponseWriter, request *http.Request) {
		app.Respond(
			wr,
			request,
			&response{
				Message: "Welcome to " + app.Config.App.Name,
			},
			http.StatusOK,
		)
	}
}
