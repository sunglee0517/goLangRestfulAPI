// controllers/dataroom_controller.go

package controllers

import (
	"company/models"
	"company/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DataroomController struct {
	DataroomService *services.DataroomService
}

func (c *DataroomController) GetDataroomList(w http.ResponseWriter, r *http.Request) {
	datarooms, err := c.DataroomService.GetDataroomList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datarooms)
}

func (c *DataroomController) GetDataroomDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dno, err := strconv.Atoi(vars["dno"])
	if err != nil {
		http.Error(w, "Invalid Dataroom ID", http.StatusBadRequest)
		return
	}

	dataroom, err := c.DataroomService.GetDataroomDetail(dno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataroom)
}

func (c *DataroomController) InsertDataroom(w http.ResponseWriter, r *http.Request) {
	var dataroom models.Dataroom
	if err := json.NewDecoder(r.Body).Decode(&dataroom); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.DataroomService.InsertDataroom(&dataroom); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *DataroomController) UpdateDataroom(w http.ResponseWriter, r *http.Request) {
	var dataroom models.Dataroom
	if err := json.NewDecoder(r.Body).Decode(&dataroom); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.DataroomService.UpdateDataroom(&dataroom); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *DataroomController) DeleteDataroom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dno, err := strconv.Atoi(vars["dno"])
	if err != nil {
		http.Error(w, "Invalid Dataroom ID", http.StatusBadRequest)
		return
	}

	if err := c.DataroomService.DeleteDataroom(dno); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
