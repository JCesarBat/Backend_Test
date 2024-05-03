package server

import (
	"Goolang_backend_Project/database/GORM"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type PutUserRequest struct {
	Id             int    `json:"id"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Identification string `json:"identification" binding:"required" `
	Password       string `json:"password" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required,min=8"`
}

type PutUserResponse struct {
	FirstName      string `json:"firstName" `
	LastName       string `json:"lastName" `
	Identification string `json:"identification" `
	Password       string `json:"password"`
	Email          string `json:"email" `
	Phone          string `json:"phone" `
}

func (server *Server) PutUser(ctx *fiber.Ctx) error {
	var req PutUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response("error", fmt.Errorf("invalid argument :%w", err).Error()))
	}

	params := GORM.UpdateUserParams{
		Id:             req.Id,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Identification: req.Identification,
		Password:       req.Password,
		Email:          req.Email,
		Phone:          req.Phone,
	}
	user, err := server.GormStore.UpdateUser(params)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response("error", fmt.Errorf("error updating the user:%w", err).Error()))
	}

	resp := PutUserResponse{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Identification: user.Identification.String,
		Password:       user.Password,
		Email:          user.Email,
		Phone:          user.Phone.String,
	}

	return ctx.JSON(resp)
}
