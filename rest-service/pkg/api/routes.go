// Package api
package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

// App is global variable for the Application
var App *Application

// Application struct contains all stuffs for the application
type Application struct {
	logger *log.Logger
	port   int
}

// InitAndRun function initialize and run Rest Api
func InitAndRun() {
	App = &Application{
		logger: log.New(os.Stdout, "App", log.Ldate|log.Ltime),
		port:   8080,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", App.port),
		Handler:      App.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	App.logger.Println("Starting server on port:", App.port)

	if err := server.ListenAndServe(); err != nil {
		App.logger.Println(err)
	}

}

func (app *Application) routes() *mux.Router {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).Path("/people/{id}").HandlerFunc(app.getPersonById)
	router.Methods(http.MethodGet).Path("/people").Queries("first_name", "{first_name}").
		Queries("last_name", "{last_name}").HandlerFunc(app.getPersonByFullName)
	router.Methods(http.MethodGet).Path("/people").Queries("phone_number", "{phone_number}").
		HandlerFunc(app.getPersonByPhone)
	router.Methods(http.MethodGet).Path("/people").HandlerFunc(app.getAllPeople)

	return router
}
