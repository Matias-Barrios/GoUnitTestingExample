package models

type User struct {
	ID       int    `db:"int" json:"id"`
	Name     string `db:"string" json:"name"`
	Lastname string `db:"string" json:"lastname"`
	Email    string `db:"string" json:"email"`
	Age      int    `db:"int" json:"age"`
}
