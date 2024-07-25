// services/qna_service.go

package services

import (
	"company/models"
	"database/sql"
)

type QnaService struct {
	DB *sql.DB
}

func (service *QnaService) GetQnaList() ([]models.Qna, error) {
	rows, err := service.DB.Query("SELECT qno, lev, parno, title, content, author, resdate, hits FROM qna")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	qnas := []models.Qna{}
	for rows.Next() {
		var qna models.Qna
		err := rows.Scan(&qna.Qno, &qna.Lev, &qna.Parno, &qna.Title, &qna.Content, &qna.Author, &qna.Resdate, &qna.Hits)
		if err != nil {
			return nil, err
		}
		qnas = append(qnas, qna)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return qnas, nil
}

func (s *QnaService) GetQnaDetail(qno int) (*models.Qna, error) {
	var qna models.Qna
	row := s.DB.QueryRow("SELECT qno, lev, parno, title, content, author, resdate, hits FROM qna WHERE qno = ?", qno)
	if err := row.Scan(&qna.Qno, &qna.Lev, &qna.Parno, &qna.Title, &qna.Content, &qna.Author, &qna.Resdate, &qna.Hits); err != nil {
		return nil, err
	}
	return &qna, nil
}

func (s *QnaService) InsertQna(qna *models.Qna) error {
	_, err := s.DB.Exec("INSERT INTO qna (lev, parno, title, content, author, resdate, hits) VALUES (?, ?, ?, ?, ?, ?, ?)",
		qna.Lev, qna.Parno, qna.Title, qna.Content, qna.Author, qna.Resdate, qna.Hits)
	return err
}

func (s *QnaService) UpdateQna(qna *models.Qna) error {
	_, err := s.DB.Exec("UPDATE qna SET lev = ?, parno = ?, title = ?, content = ?, author = ?, resdate = ?, hits = ? WHERE qno = ?",
		qna.Lev, qna.Parno, qna.Title, qna.Content, qna.Author, qna.Resdate, qna.Hits, qna.Qno)
	return err
}

func (s *QnaService) DeleteQna(qno int) error {
	_, err := s.DB.Exec("DELETE FROM qna WHERE qno = ?", qno)
	return err
}
