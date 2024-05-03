package server

import (
	database "Goolang_backend_Project/database/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User             database.User             `json:"user"`
	StatusConnection database.NullRecordStatus `json:"statusConnection"`
}

func (server *Server) Login(ctx *fiber.Ctx) error {
	var request LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response("error", fmt.Errorf("the params are invalid : %w", err).Error()))

	}
	user, err := server.GormStore.GetUserEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(response("error", fmt.Errorf("not exist any user whuit this email").Error()))

		}
		return ctx.Status(http.StatusInternalServerError).JSON(response("error", fmt.Errorf("internal server error ").Error()))
	}
	if user.Password != request.Password {
		return ctx.Status(http.StatusUnauthorized).JSON(response("error", fmt.Errorf("the password dont match").Error()))
	}
	arg := database.UpdateConnectionParams{
		UserID: user.UserID,
		Active: database.RecordStatusAppstatusActive,
	}
	connection, err := server.SqlStore.UpdateConnection(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(http.StatusNotFound).JSON(response("error", fmt.Errorf("you user have conflict whit the connection table plese create the connection table for this user %w", err).Error()))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response("error", fmt.Errorf("internal server error connection :%w", err).Error()))
	}
	ctx.Locals("Cache", connection)

	response := LoginResponse{
		User: user,
		StatusConnection: database.NullRecordStatus{
			RecordStatus: connection.Active,
			Valid:        true},
	}

	return ctx.JSON(response)
}
