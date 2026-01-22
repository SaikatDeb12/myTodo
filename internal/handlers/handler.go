package handlers

import (
	"encoding/json"
	"myTodo/internal/models"
	"myTodo/internal/storage"
	"myTodo/internal/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetTodos(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(storage.Todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){
	var todo models.Todo
	err :=json.NewDecoder(r.Body).Decode(&todo)
	if(err != nil){
		utils.RespondError(w, http.StatusBadRequest, err, "Enable to decode the body")
	}

	len := len(storage.Todos)
	todo.ID=len+1
	storage.Todos=append(storage.Todos, todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request){
	id, _ :=strconv.Atoi(chi.URLParam(r, "id"))

	var updated models.Todo
	err := json.NewDecoder(r.Body).Decode(&updated)
	if(err !=nil){
		utils.RespondError(w, http.StatusNotModified, err, "Decode Error")
	}

	for i,todo := range storage.Todos{
		if todo.ID==id{
			storage.Todos[i].Title=updated.Title
			storage.Todos[i].Completed=updated.Completed
			storage.Todos[i].Details=updated.Details
			json.NewEncoder(w).Encode(storage.Todos[i])
			return
		}
	}
	
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	id,_ := strconv.Atoi(chi.URLParam(r, "id"))
	for i, todo := range storage.Todos{
		if (todo.ID==id){
			storage.Todos=append(storage.Todos[:i], storage.Todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	utils.RespondError(w, http.StatusNotFound, nil, "Todo not found!")
}
