package routes

import (
	"net/http"

	"github.com/JustSteveKing/example-go-api/pkg/kernel"
)

func Load(app *kernel.Application) {
	serviceRouter := app.Router.Methods(http.MethodGet).Subrouter()

	serviceRouter.HandleFunc("/", HandleApiRoot(app)).Name("api:root")
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
