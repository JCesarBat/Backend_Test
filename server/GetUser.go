package server

import (
	database "Goolang_backend_Project/database/sqlc"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type GetUserResponse struct {
	User database.User `json:"user"`
}

func (server *Server) GetUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("user_id")
	idInt, err := strconv.Atoi(userID)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(response("error", err))
	}

	user, err := server.GormStore.GetUserID(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(response("error", fmt.Errorf("not exist any user whuit this id").Error()))

		}
		return ctx.Status(http.StatusInternalServerError).JSON(response("error", fmt.Errorf("internal server error ").Error()))

	}
	res := GetUserResponse{User: user}

	return ctx.JSON(res)
}
