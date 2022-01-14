package api

import (
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"net/http"
	"strings"
)

// getPersonById router handler function to get Person by ID
func (app *Application) getPersonById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	uuidStr := params["id"]
	id, err := uuid.FromString(uuidStr)
	if err != nil {
		app.logger.Print(uuidStr, err)
		app.errorJSON(w, http.StatusBadRequest, err)
		return
	}

	app.logger.Println("id is", id)

	person, err := models.FindPersonByID(id)
	if err != nil {
		app.errorJSON(w, http.StatusNotFound, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, person, "people")
	if err != nil {
		app.logger.Println(err)
	}

}

func (app *Application) getAllPeople(w http.ResponseWriter, r *http.Request) {

	people := models.AllPeople()
	err := app.writeJSON(w, http.StatusOK, people, "people")
	if err != nil {
		app.logger.Println(err)
	}

}

func (app *Application) getPersonByFullName(w http.ResponseWriter, r *http.Request) {

	fName := r.URL.Query()["first_name"]
	lName := r.URL.Query()["last_name"]
	app.logger.Println(fName, lName)

	person := models.FindPeopleByName(fName[0], lName[0])
	err := app.writeJSON(w, http.StatusOK, person, "people")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *Application) getPersonByPhone(w http.ResponseWriter, r *http.Request) {

	phone := r.URL.Query()["phone_number"]
	app.logger.Println("phone", phone)
	people := models.FindPeopleByPhoneNumber(strings.Join(phone, ""))

	err := app.writeJSON(w, http.StatusOK, people, "people")
	if err != nil {
		app.logger.Println(err)
	}
}
