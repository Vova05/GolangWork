package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	time2 "time"
)



func (ts *taskServer) taskHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/task/" {
		// Запрос направлен к "/task/", без идущего в конце ID.
		if req.Method == http.MethodPost {
			//ts.createTaskHandler(w, req)
		} else if req.Method == http.MethodGet {
			ts.getAllTasksHandler(w, req)
		} else if req.Method == http.MethodDelete {
			//ts.deleteAllTasksHandler(w, req)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}
func (ts *taskServer) getAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all tasks at %s\n", req.URL.Path)




	allTasks := ts.store.GetAllTasks()
	js, err := json.Marshal(allTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func (ts *taskServer) createTaskHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling create tasks at %s\n", req.URL.Path)

	name := req.URL.Query().Get("name")
	if name == "" {
		http.NotFound(w, req)
		return
	}


	arr := []string{"new","tag"}
	time := time2.Now()
	createTasks := ts.store.CreateTask(name,arr,time)
	js, err := json.Marshal(createTasks )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func (ts *taskServer) getTaskHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get  task at %s\n", req.URL.Path)
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, req)
		return
	}
	getTasks, _ := ts.store.GetTask(id)
	js, err := json.Marshal(getTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}