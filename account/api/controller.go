package api

import (
	"encoding/json"
	"github/four-servings/meonzi/account/app/command"
	"github/four-servings/meonzi/account/app/query"
	"net/http"
)

type (
	// Controller http controller
	Controller interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}

	controllerImplement struct {
		commandBus command.Bus
		queryBus   query.Bus
	}

	createAccountBody struct {
		Name string
	}
)

// NewController create controller instance
func NewController(commandBus command.Bus, queryBus query.Bus) Controller {
	return &controllerImplement{commandBus, queryBus}
}

// Handle handle http request
func (c *controllerImplement) Handle(w http.ResponseWriter, r *http.Request) {
	c.branchByMethod(w, r)
}

func (c *controllerImplement) branchByMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.handlePOST(w, r)
		return
	case "GET":
		c.handleGET(w, r)
		return
	default:
		c.handleNotAllowedMethod(w, r)
		return
	}
}

func (c *controllerImplement) handlePOST(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	body := createAccountBody{}
	if decoder.Decode(&body) != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	c.commandBus.Handle(&command.CreateAccount{Name: body.Name})
	w.WriteHeader(http.StatusCreated)
}

func (c *controllerImplement) handleGET(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		result := c.queryBus.Handle(&query.Find{Name: r.URL.Query().Get("name")})
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (c *controllerImplement) handleNotAllowedMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
