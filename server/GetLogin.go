package server

import (
	database "Goolang_backend_Project/database/sqlc"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type GetLoginResponse struct {
	StatusConnection database.RecordStatus `json:"statusConnection"`
	Connection       database.Connection   `json:"connection"`
}

func (Server *Server) GetLogin(ctx *fiber.Ctx) error {
	cache := ctx.Locals("Cache")
	connection, ok := cache.(database.Connection)

	if !ok {
		return ctx.Status(http.StatusNotAcceptable).JSON(response("error", fmt.Errorf("the cahe dont have a connection ").Error()))
	}
	res := GetLoginResponse{
		StatusConnection: connection.Active,
		Connection:       connection,
	}

	return ctx.JSON(res)

}
