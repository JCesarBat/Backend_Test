package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type DeleteUserResponse struct {
	Valid bool `json:"valid"`
}

func (server *Server) DeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("user_id")
	idInt, err := strconv.Atoi(userID)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(response("error", err))
	}

	err = server.GormStore.DeleteUser(idInt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response("error", fmt.Errorf("error deleting the user : %w", err).Error()))
	}

	return ctx.JSON(DeleteUserResponse{Valid: true})
}
