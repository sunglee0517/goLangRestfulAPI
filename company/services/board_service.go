// services/board_service.go

package services

import (
	"company/models"
	"database/sql"
)

type BoardService struct {
	DB *sql.DB
}

func (s *BoardService) GetBoardList() ([]models.Board, error) {
	var boards []models.Board
	rows, err := s.DB.Query("SELECT no, title, content, author, resdate, hits FROM board")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var board models.Board
		if err := rows.Scan(&board.No, &board.Title, &board.Content, &board.Author, &board.ResDate, &board.Hits); err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}
	return boards, nil
}

func (s *BoardService) GetBoardDetail(no int) (*models.Board, error) {
	var board models.Board
	row := s.DB.QueryRow("SELECT no, title, content, author, resdate, hits FROM board WHERE no = ?", no)
	if err := row.Scan(&board.No, &board.Title, &board.Content, &board.Author, &board.ResDate, &board.Hits); err != nil {
		return nil, err
	}
	return &board, nil
}

func (s *BoardService) InsertBoard(board *models.Board) error {
	_, err := s.DB.Exec("INSERT INTO board (title, content, author, resdate, hits) VALUES (?, ?, ?, ?, ?)",
		board.Title, board.Content, board.Author, board.ResDate, board.Hits)
	return err
}

func (s *BoardService) UpdateBoard(board *models.Board) error {
	_, err := s.DB.Exec("UPDATE board SET title = ?, content = ?, author = ?, resdate = ?, hits = ? WHERE no = ?",
		board.Title, board.Content, board.Author, board.ResDate, board.Hits, board.No)
	return err
}

func (s *BoardService) DeleteBoard(no int) error {
	_, err := s.DB.Exec("DELETE FROM board WHERE no = ?", no)
	return err
}
