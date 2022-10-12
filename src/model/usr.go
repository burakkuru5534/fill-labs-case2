package model

import (
	"github.com/burakkuru5534/src/helper"
)

type Usr struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (u *Usr) Create() error {

	sq := "INSERT INTO usr (email) VALUES ($1) RETURNING id"
	err := helper.App.DB.QueryRow(sq, u.Email).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usr) Update(id int64) error {

	sq := "UPDATE usr SET  email = $1 WHERE id = $2"
	_, err := helper.App.DB.Exec(sq, u.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Delete(id int64) error {

	sq := "DELETE FROM usr WHERE id = $1"
	_, err := helper.App.DB.Exec(sq, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Get(id int64) error {

	sq := "SELECT email FROM usr WHERE id = $1"
	err := helper.App.DB.QueryRow(sq, id).Scan(&u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) GetAll() ([]Usr, error) {

	rows, err := helper.App.DB.Query("SELECT email FROM usr")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var usrs []Usr

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var usr Usr
		if err := rows.Scan(&usr.Email); err != nil {
			return usrs, err
		}
		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return usrs, err
	}
	return usrs, nil
}
