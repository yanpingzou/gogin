package models

type User struct {
	ID       int64  `json:"id" database:"id" form:"name"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	//Timestamp time.Time `json:"timestamp" form:"timestamp"`
}

func (user *User) Insert() (int64, error) {
	r := rds.Get()
	r.Do("SET", "key1", "value1")
	defer r.Close()
	stmt, err := db.Prepare("INSERT INTO user(`name`,email,password) VALUES (?, ?, ?)") // ? = placeholder
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}
