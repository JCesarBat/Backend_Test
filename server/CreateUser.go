package server

import (
	database "Goolang_backend_Project/database/sqlc"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type CreateUserRequest struct {
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Identification string `json:"identification" binding:"required" `
	Password       string `json:"password" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required,min=8"`
}
type CreateUserResponse struct {
	User             database.User             `json:"user"`
	StatusConnection database.NullRecordStatus `json:"statusConnection"`
}

func (server *Server) CreateUser(ctx *fiber.Ctx) error {
	var request CreateUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response("error", fmt.Errorf("the request is invalid :%w", err).Error()))
	}
	user := database.User{
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		Identification: sql.NullString{String: request.Identification, Valid: true},
		Password:       request.Password,
		Email:          request.Email,
		Phone:          sql.NullString{String: request.Phone, Valid: true},
	}
	user, err := server.GormStore.CreateUser(user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ctx.Status(http.StatusConflict).JSON(response("error", "a field is duplicated in the database"))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response("error", fmt.Errorf("error %w", err).Error()))
	}
	connection := database.Connection{
		UserID:    user.UserID,
		AccountID: user.AccountID,
		Active:    database.RecordStatusAppstatusInactive,
	}

	connection, err = server.GormStore.CreateConnection(connection)
	if err != nil {

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ctx.Status(http.StatusConflict).JSON(response("error", "a field is duplicated in the database"))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response("error", fmt.Errorf("error %w", err).Error()))
	}
	res := CreateUserResponse{
		User:             user,
		StatusConnection: database.NullRecordStatus{RecordStatus: connection.RecordStatus.RecordStatus, Valid: true},
	}
	return ctx.JSON(res)
}
