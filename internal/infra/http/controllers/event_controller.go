package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" { // if the request method is get, json will be returned with the fields to be filled
			data := event.Event{}
			encoder := json.NewEncoder(w)
			err := encoder.Encode(&data)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.FindOne(): %s", err)
				}
				return
			}
		}

		if r.Method == "POST" {
			var data *event.Event
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.FindOne(): %s", err)
				}
				return
			}

			id := data.Id
			name := data.Name

			err2 := (*c.service).Create(id, name)

			if err2 != nil {
				fmt.Printf("EventController.Create(): %s", err)
				err2 = internalServerError(w, err)
				if err2 != nil {
					fmt.Printf("EventController.Create(): %s", err)
				}
				return
			}

			err = success(w, "Created")
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := event.Event{}
			encoder := json.NewEncoder(w)
			err := encoder.Encode(&data)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.Update(): %s", err)
				}
				return
			}
		}

		if r.Method == "PUT" {
			var data *event.Event
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.Update(): %s", err)
				}
				return
			}

			id := data.Id
			name := data.Name

			err2 := (*c.service).Update(id, name)

			if err2 != nil {
				fmt.Printf("EventController.Update(): %s", err)
				err2 = internalServerError(w, err)
				if err2 != nil {
					fmt.Printf("EventController.Update(): %s", err)
				}
				return
			}

			err = success(w, "Updated")
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
		}

	}
}

func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := event.Event{}
			encoder := json.NewEncoder(w)
			err := encoder.Encode(&data)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.Delete(): %s", err)
				}
				return
			}
		}

		if r.Method == "DELETE" {
			var data *event.Event
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.Delete(): %s", err)
				}
				return
			}

			id := data.Id
			name := data.Name

			err2 := (*c.service).Delete(id, name)
			if err2 != nil {
				fmt.Printf("EventController.Delete(): %s", err)
				err2 = internalServerError(w, err)
				if err2 != nil {
					fmt.Printf("EventController.Delete(): %s", err)
				}
				return
			}

			err2 = success(w, "Deleted")
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
		}
	}
}
