package repository

import (
	"context"
	"errors"
	"time"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (u *UserRepository) RegisMentor(ctx context.Context, user *model.MentorRegis) error {

	query := "INSERT INTO mentor (skill,bio,name,address,phone,email,created_at,image) VALUES (?,?,?,?,?,?,?,?)"
	tx, _ := u.Db.Begin()

	_, err := tx.ExecContext(ctx, query, user.Skill, user.Bio, user.Name, user.Address, user.Phone, user.Email, time.Now(), user.Image)
	if err != nil {
		tx.Rollback()
		return errors.New("Error Server")
	}
	tx.Commit()

	return nil

}

func (u *UserRepository) GetMentorByskill(ctx context.Context, skill string) ([]*model.MentorSkill, error) {
	var mentors []*model.MentorSkill

	query := "SELECT id,name,bio,skill,image FROM mentor WHERE skill =?"
	rows, err := u.Db.QueryContext(ctx, query, skill)
	if err != nil {
		return nil, errors.New("Error getting mentor")
	}
	defer rows.Close()
	for rows.Next() {
		var mentor model.MentorSkill
		err := rows.Scan(&mentor.ID, &mentor.Name, &mentor.Bio, &mentor.Skill, &mentor.Image)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, &mentor)
	}
	return mentors, nil

}

func (u *UserRepository) CheckEmailMentor(ctx context.Context, email string) error {
	rows, err := u.Db.QueryContext(ctx, "SELECT id from mentor where email = ?", email)

	if err != nil {
		return errors.New("Error Server")
	}
	defer rows.Close()

	if rows.Next() {
		return errors.New("Email already exist")
	}
	return nil

}
func (u *UserRepository) GetMentorById(ctx context.Context, id int) (*model.MentorDetail, error) {
	query := "SELECT id,name,skill,bio,image FROM mentor WHERE id = ?"

	var mentor model.MentorDetail
	rows, err := u.Db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&mentor.ID, &mentor.Name, &mentor.Skill, &mentor.Bio, &mentor.Image)
		if err != nil {
			return nil, err
		}
		return &mentor, nil
	}
	return nil, errors.New("Mentor not found")

}

func (u *UserRepository) GetMentorEmailById(ctx context.Context, id int) (string, error) {
	query := "SELECT email FROM mentor WHERE id = ?"
	var email string
	err := u.Db.QueryRowContext(ctx, query, id).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (u *UserRepository) CheckMentorBySkill(ctx context.Context, skill string) error {
	rows, err := u.Db.QueryContext(ctx, "SELECT id from mentor where skill = ?", skill)

	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("Category not found")

}

func (u *UserRepository) MentorList(ctx context.Context) ([]*model.MentorDetail, error) {
	var mentors []*model.MentorDetail
	query := "SELECT id,name,skill,bio,image FROM mentor"
	rows, err := u.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var mentor model.MentorDetail
		err := rows.Scan(&mentor.ID, &mentor.Name, &mentor.Skill, &mentor.Bio, &mentor.Image)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, &mentor)
	}
	return mentors, nil
}

func (u *UserRepository) GetDataMentorByNoBooking(ctx context.Context, noorder string) (*model.MentorKontak, error) {
	query := "SELECT m.name,m.phone,m.email,m.address FROM bookmentor u INNER JOIN mentor m ON u.mentor_id = m.id WHERE u.bookid = ? AND u.status = ?"
	rows, err := u.Db.QueryContext(ctx, query, noorder, "Accepted")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var mentor model.MentorKontak
	if rows.Next() {
		err := rows.Scan(&mentor.Name, &mentor.Phone, &mentor.Email, &mentor.Address)
		if err != nil {
			return nil, err
		}
		return &mentor, nil
	}
	return nil, errors.New("Detail Mentor Tidak ditemukan")
}
