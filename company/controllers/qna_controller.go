// controllers/qna_controller.go

package controllers

import (
	"company/models"
	"company/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type QnaController struct {
	QnaService *services.QnaService
}

func (c *QnaController) GetQnaList(w http.ResponseWriter, r *http.Request) {
	qnas, err := c.QnaService.GetQnaList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qnas)
}

func (c *QnaController) GetQnaDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qno, err := strconv.Atoi(vars["qno"])
	if err != nil {
		http.Error(w, "Invalid Qna ID", http.StatusBadRequest)
		return
	}

	qna, err := c.QnaService.GetQnaDetail(qno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qna)
}

func (c *QnaController) InsertQna(w http.ResponseWriter, r *http.Request) {
	var qna models.Qna
	if err := json.NewDecoder(r.Body).Decode(&qna); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.QnaService.InsertQna(&qna); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *QnaController) UpdateQna(w http.ResponseWriter, r *http.Request) {
	var qna models.Qna
	if err := json.NewDecoder(r.Body).Decode(&qna); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.QnaService.UpdateQna(&qna); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *QnaController) DeleteQna(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qno, err := strconv.Atoi(vars["qno"])
	if err != nil {
		http.Error(w, "Invalid Qna ID", http.StatusBadRequest)
		return
	}

	if err := c.QnaService.DeleteQna(qno); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
