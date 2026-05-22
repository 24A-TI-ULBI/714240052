package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

const AbsensiCollection = "absensi"

// GetAbsensiByNPM menangani GET /absensi/:npm
func GetAbsensiByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"npm": npm}
	
	docs, err := helper.GetAllDoc[model.Absensi](db, AbsensiCollection, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data absensi: "+err.Error())
	}

	// Mengembalikan array objek absensi langsung sesuai ekspektasi frontend
	return c.JSON(docs)
}

// InsertAbsensi menangani POST /absensi
func InsertAbsensi(c *fiber.Ctx) error {
	var req model.Absensi
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.ID == "" || req.NPM == "" || req.MatkulCode == "" || req.Tanggal == "" || req.Status == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Data absensi tidak lengkap")
	}

	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, AbsensiCollection, req)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan data absensi: "+err.Error())
	}

	return helper.SuccessResponse(c, req)
}

// UpdateAbsensi menangani PUT /absensi/:id
func UpdateAbsensi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak boleh kosong")
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Status == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Status tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": req.Status}}

	_, err := helper.UpdateDoc(db, AbsensiCollection, filter, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui status absensi: "+err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{"message": "Status absensi berhasil diperbarui"})
}

// DeleteAbsensi menangani DELETE /absensi/:id
func DeleteAbsensi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"_id": id}

	_, err := helper.DeleteDoc(db, AbsensiCollection, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data absensi: "+err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{"message": "Data absensi berhasil dihapus"})
}
