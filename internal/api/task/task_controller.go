package task

import (
	"awesomeProject/internal/common"
	"awesomeProject/internal/common/response"
	"awesomeProject/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type controller struct {
	taskRepository domain.Repository
}

func NewController(taskRepository domain.Repository) common.Controller {
	return &controller{
		taskRepository: taskRepository,
	}
}

func (c *controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/task", c.GetAll).Methods("GET")
	router.HandleFunc("/task/{id}", c.Get).Methods("GET")
	router.HandleFunc("/task", c.Add).Methods("POST")
	router.HandleFunc("/task", c.RemoveAll).Methods("DELETE")
	router.HandleFunc("/task/{id}", c.Remove).Methods("DELETE")
}

func (c *controller) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
		response.WriteJson(w, response.Error{
			Message: "Invalid task id",
		}, http.StatusBadRequest)
		return
	}

	task := c.taskRepository.Get(id)

	if task == nil {
		response.WriteJson(w, response.Error{
			Message: "Task not found",
		}, http.StatusBadRequest)
		return
	}

	response.WriteJson(w, struct {
		Task *domain.Task `json:"task"`
	}{
		Task: task,
	}, http.StatusOK)
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, struct {
		Task []*domain.Task `json:"tasks"`
	}{
		Task: c.taskRepository.GetAll(),
	}, http.StatusOK)
}

func (c *controller) Add(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		response.WriteJson(w, response.Error{
			Message: fmt.Sprintf("Got invalid task %s", err),
		}, http.StatusBadRequest)
		return
	}

	if err := c.taskRepository.Add(&task); err != nil {
		response.WriteJson(w, response.Error{
			Message: "Failed to add task",
		}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
		response.WriteJson(w, response.Error{
			Message: "Invalid task id",
		}, http.StatusBadRequest)
		return
	}

	if err := c.taskRepository.Remove(id); err != nil {
		response.WriteJson(w, response.Error{
			Message: "Failed to remove task",
		}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) RemoveAll(w http.ResponseWriter, r *http.Request) {
	if err := c.taskRepository.RemoveAll(); err != nil {
		response.WriteJson(w, response.Error{
			Message: "Failed to remove all tasks",
		}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
