package main

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app fiber.Router) {
	app.Get("/pedagang", GetPedagang)
	app.Get("/pedagang/:id", GetPedagangById)
	app.Get("/jenis-dagangan", GetJenisDagangan)
}
