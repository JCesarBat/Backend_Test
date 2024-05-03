package GORM

import (
	database "Goolang_backend_Project/database/sqlc"
	"database/sql"
	"fmt"
)

func (store *GormStore) CreateUser(user database.User) (database.User, error) {
	result := store.Db.Select(
		"FirstName",
		"LastName",
		"Identification",
		"Password",
		"Email",
		"Phone").Create(&user)
	if result.Error != nil {
		return database.User{}, result.Error
	}
	user, err := store.GetUserEmail(user.Email)
	if err != nil {
		return database.User{}, err
	}
	return user, nil

}

func (store *GormStore) GetUserEmail(email string) (database.User, error) {
	var user database.User
	result := store.Db.Where("email = ? ", email).First(&user)
	if result.Error != nil {
		return database.User{}, result.Error
	}
	return user, nil
}

func (store *GormStore) GetUserID(id int) (database.User, error) {
	var user database.User
	result := store.Db.Where("user_id = ? ", id).First(&user)
	if result.Error != nil {
		return database.User{}, result.Error
	}
	return user, nil
}

type UpdateUserParams struct {
	Id             int    `json:"id	"`
	FirstName      string `json:"firstName" `
	LastName       string `json:"lastName" `
	Identification string `json:"identification" `
	Password       string `json:"password"`
	Email          string `json:"email" `
	Phone          string `json:"phone" `
}

func (store *GormStore) UpdateUser(arg UpdateUserParams) (database.User, error) {
	var user database.User
	result := store.Db.Model(&user).Where("user_id = ?", arg.Id).Updates(database.User{
		FirstName:      arg.FirstName,
		LastName:       arg.LastName,
		Identification: sql.NullString{String: arg.Identification, Valid: true},
		Password:       arg.Password,
		Email:          arg.Email,
		Phone:          sql.NullString{String: arg.Phone, Valid: true},
	})
	if result.Error != nil {
		return database.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return database.User{}, fmt.Errorf("not found any user to update")
	}

	return user, nil
}

func (store *GormStore) DeleteUser(user_id int) error {

	result := store.Db.Delete(&database.User{}, user_id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("not found any user to delete ")
	}
	return nil
}
