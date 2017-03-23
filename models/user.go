package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" database:"id" form:"name"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Timestamp time.Time `json:"timestamp" form:"timestamp"`
}

func (user User) Insert() (int64, error) {
	stmtIns, err := db.Prepare("insert into user(name,email,password) VALUES(?, ?, ?)") // ? = placeholder
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastRowID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastRowID, nil
}
