package main

import (
	"github.com/gorilla/mux"
	"github.com/grameh/GoDevEnvironment/controller"
	"github.com/grameh/GoDevEnvironment/dataModel"
	"github.com/grameh/GoDevEnvironment/storage"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	reader := storage.GetReader()
	writer := storage.GetWriter()

	subjectDataModel := dataModel.SubjectDataModel{
		Reader: reader,
		Writer: writer,
	}

	subjectController := controller.SubjectController{
		DataModel: subjectDataModel,
	}

	router.HandleFunc("/subject", subjectController.ReadAll).Methods("GET")
	router.HandleFunc("/subject", subjectController.Write).Methods("PUT")
	router.HandleFunc("/subject/{subjectId}", subjectController.Read).Methods("GET")
	router.HandleFunc("/subject/{subjectId}", subjectController.Update).Methods("PATCH")
	router.HandleFunc("/subject/{subjectId}", subjectController.Delete).Methods("DELETE")

	router.PathPrefix("/static").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
