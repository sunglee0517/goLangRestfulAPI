// services/member_service.go

package services

import (
	"company/models"
	"database/sql"
)

type MemberService struct {
	DB *sql.DB
}

func (s *MemberService) GetMemberList() ([]models.Member, error) {
	var members []models.Member
	rows, err := s.DB.Query("SELECT id, pw, name, birth, email, tel, addr1, addr2, postcode, regdate FROM member")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member models.Member
		if err := rows.Scan(&member.ID, &member.PW, &member.Name, &member.Birth, &member.Email, &member.Tel, &member.Addr1, &member.Addr2, &member.Postcode, &member.RegDate); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (s *MemberService) GetMember(id string) (*models.Member, error) {
	var member models.Member
	row := s.DB.QueryRow("SELECT id, pw, name, birth, email, tel, addr1, addr2, postcode, regdate FROM member WHERE id = ?", id)
	if err := row.Scan(&member.ID, &member.PW, &member.Name, &member.Birth, &member.Email, &member.Tel, &member.Addr1, &member.Addr2, &member.Postcode, &member.RegDate); err != nil {
		return nil, err
	}
	return &member, nil
}

func (s *MemberService) InsertMember(member *models.Member) error {
	_, err := s.DB.Exec("INSERT INTO member (id, pw, name, birth, email, tel, addr1, addr2, postcode, regdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		member.ID, member.PW, member.Name, member.Birth, member.Email, member.Tel, member.Addr1, member.Addr2, member.Postcode, member.RegDate)
	return err
}

func (s *MemberService) UpdateMember(member *models.Member) error {
	_, err := s.DB.Exec("UPDATE member SET pw = ?, name = ?, birth = ?, email = ?, tel = ?, addr1 = ?, addr2 = ?, postcode = ? WHERE id = ?",
		member.PW, member.Name, member.Birth, member.Email, member.Tel, member.Addr1, member.Addr2, member.Postcode, member.ID)
	return err
}

func (s *MemberService) DeleteMember(id string) error {
	_, err := s.DB.Exec("DELETE FROM member WHERE id = ?", id)
	return err
}
