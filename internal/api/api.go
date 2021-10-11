package api

import (
	"awesomeProject/internal/api/task"
	"awesomeProject/internal/common"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Api interface {
	Run()
}

type api struct {
	taskController common.Controller

	config Config
	router *mux.Router
}

func NewApi(config Config) Api {
	return &api{
		config: config,
	}
}

func (api *api) Run() {
	api.router = mux.NewRouter()

	api.configureServices()
	api.configurePipeline()

	api.router.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Some Result")
	})

	log.Printf("Listening at localhost[%s]", api.config.Address)
	if err := http.ListenAndServe(api.config.Address, api.router); err != nil {
		log.Fatal("ListenAndServe failed", err)
	}
}

func (api *api) configureServices() {
	taskRepository := task.NewRepository()
	api.taskController = task.NewController(taskRepository)
}

func (api *api) configurePipeline() {
	api.taskController.RegisterRoutes(api.router)
}
