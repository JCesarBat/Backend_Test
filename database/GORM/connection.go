package GORM

import database "Goolang_backend_Project/database/sqlc"

func (store *GormStore) CreateConnection(connection database.Connection) (database.Connection, error) {
	result := store.Db.Select(
		"user_id",
		"account_id",
		"active",
	).Create(&connection)
	if result.Error != nil {
		return database.Connection{}, result.Error
	}
	return connection, nil

}
