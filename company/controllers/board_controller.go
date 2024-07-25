// controllers/board_controller.go

package controllers

import (
	"company/models"
	"company/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BoardController struct {
	BoardService *services.BoardService
}

func (c *BoardController) GetBoardList(w http.ResponseWriter, r *http.Request) {
	boards, err := c.BoardService.GetBoardList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boards)
}

func (c *BoardController) GetBoardDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	no, err := strconv.Atoi(vars["no"])
	if err != nil {
		http.Error(w, "Invalid board ID", http.StatusBadRequest)
		return
	}

	board, err := c.BoardService.GetBoardDetail(no)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(board)
}

func (c *BoardController) InsertBoard(w http.ResponseWriter, r *http.Request) {
	var board models.Board
	if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.BoardService.InsertBoard(&board); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *BoardController) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var board models.Board
	if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.BoardService.UpdateBoard(&board); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *BoardController) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	no, err := strconv.Atoi(vars["no"])
	if err != nil {
		http.Error(w, "Invalid board ID", http.StatusBadRequest)
		return
	}

	if err := c.BoardService.DeleteBoard(no); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
