package server

import (
	"Goolang_backend_Project/database/GORM"
	database "Goolang_backend_Project/database/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

type Server struct {
	App       *fiber.App
	GormStore *GORM.GormStore
	SqlStore  *database.SQLStore
}

func StartServer(GormStore *GORM.GormStore, store *database.SQLStore) *Server {

	server := Server{
		GormStore: GormStore,
		SqlStore:  store,
	}
	app := fiber.New()

	grupoConMiddleware := app.Group("/", cache.New())
	grupoConMiddleware.Get("auth/login", server.GetLogin)
	grupoConMiddleware.Post("auth/login", server.Login)

	app.Post("/auth", server.CreateUser)
	app.Get("/auth/:user_id", server.GetUser)
	server.App = app
	return &server
}

func response(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}
