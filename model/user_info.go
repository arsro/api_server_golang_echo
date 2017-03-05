package model

import (
	"github.com/gocraft/dbr"
)

type (
	User struct {
		Id 		int 	`json:"id"`
		Name 	string 	`json:"name"`
	}
)

/**
 * initialize User struct
 */
func NewUser(user_name string) *User {
	u := new(User)
	u.Name = user_name

	return u
}


/**
 * Select User
 */
func (u *User) Select(tx *dbr.Tx, user_id int) error {
	//https://github.com/gocraft/dbr/blob/10fa49bf1ee63a0f822283ef9805d818f9e84742/select_builder.go
	err := tx.Select("*").From("user_info").Where("id = ?", user_id).LoadStruct(u)
	
	return err
}

/**
 * Insert User
 */
func (u *User) Insert(tx *dbr.Tx) error {
	//https://github.com/gocraft/dbr/blob/10fa49bf1ee63a0f822283ef9805d818f9e84742/insert_builder.go
 	_, err := tx.InsertInto("user_info").Columns("name").Record(u).Exec()
	return err
}


/**
 * Update user
 */
func (u *User) Update(tx *dbr.Tx) error {
	_, err := tx.Update("user_info").Set("name",u.Name).Where("id = ?", u.Id).Exec()
	return err
}

/**
 * Delete user
 */
func (u *User) Delete(tx *dbr.Tx, user_id int) error {
	_, err := tx.DeleteFrom("user_info").Where("id = ?", user_id).Exec()
	return err
}
