package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func AbsensiRoute(app *fiber.App) {
	// Attendance
	app.Get("/absensi/hari-ini", controller.GetAbsensiHariIni)
	app.Get("/absensi/:npm", controller.GetAbsensiByNPM)
	app.Post("/absensi", controller.InsertAbsensi)
	app.Put("/absensi/:id", controller.UpdateAbsensi)
	app.Delete("/absensi/:id", controller.DeleteAbsensi)

	// Rekap Absensi
	app.Get("/rekap-absensi/matkul/:kode", controller.GetRekapAbsensiByMatkul)
	app.Get("/rekap-absensi/:npm", controller.GetRekapAbsensiByNPM)
}
