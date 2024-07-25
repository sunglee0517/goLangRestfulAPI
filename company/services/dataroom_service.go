// services/dataroom_service.go

package services

import (
	"company/models"
	"database/sql"
)

type DataroomService struct {
	DB *sql.DB
}

func (s *DataroomService) GetDataroomList() ([]models.Dataroom, error) {
	var datarooms []models.Dataroom
	rows, err := s.DB.Query("SELECT dno, title, content, author, datafile, resdate, hits FROM dataroom")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dataroom models.Dataroom
		if err := rows.Scan(&dataroom.Dno, &dataroom.Title, &dataroom.Content, &dataroom.Author, &dataroom.Datafile, &dataroom.ResDate, &dataroom.Hits); err != nil {
			return nil, err
		}
		datarooms = append(datarooms, dataroom)
	}
	return datarooms, nil
}

func (s *DataroomService) GetDataroomDetail(dno int) (*models.Dataroom, error) {
	var dataroom models.Dataroom
	row := s.DB.QueryRow("SELECT dno, title, content, author, datafile, resdate, hits FROM dataroom WHERE dno = ?", dno)
	if err := row.Scan(&dataroom.Dno, &dataroom.Title, &dataroom.Content, &dataroom.Author, &dataroom.Datafile, &dataroom.ResDate, &dataroom.Hits); err != nil {
		return nil, err
	}
	return &dataroom, nil
}

func (s *DataroomService) InsertDataroom(dataroom *models.Dataroom) error {
	_, err := s.DB.Exec("INSERT INTO dataroom (title, content, author, datafile, resdate, hits) VALUES (?, ?, ?, ?, ?, ?)",
		dataroom.Title, dataroom.Content, dataroom.Author, dataroom.Datafile, dataroom.ResDate, dataroom.Hits)
	return err
}

func (s *DataroomService) UpdateDataroom(dataroom *models.Dataroom) error {
	_, err := s.DB.Exec("UPDATE dataroom SET title = ?, content = ?, author = ?, datafile = ?, resdate = ?, hits = ? WHERE dno = ?",
		dataroom.Title, dataroom.Content, dataroom.Author, dataroom.Datafile, dataroom.ResDate, dataroom.Hits, dataroom.Dno)
	return err
}

func (s *DataroomService) DeleteDataroom(dno int) error {
	_, err := s.DB.Exec("DELETE FROM dataroom WHERE dno = ?", dno)
	return err
}
