package GORM

import (
	database "Goolang_backend_Project/database/sqlc"
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
