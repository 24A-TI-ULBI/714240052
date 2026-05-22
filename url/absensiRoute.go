package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func AbsensiRoute(app *fiber.App) {
	app.Get("/absensi/:npm", controller.GetAbsensiByNPM)
	app.Post("/absensi", controller.InsertAbsensi)
	app.Put("/absensi/:id", controller.UpdateAbsensi)
	app.Delete("/absensi/:id", controller.DeleteAbsensi)
}
