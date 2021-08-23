package routes

import (
	"os"
	"testing"

	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/edge-api/pkg/db"
	"github.com/redhatinsights/edge-api/pkg/models"
)

var (
	testImage models.Image
	testRepo  models.Repo

	updateDevices = []models.Device{
		{UUID: "1", DesiredHash: "11"},
		{UUID: "2", DesiredHash: "11"},
		{UUID: "3", DesiredHash: "22"},
		{UUID: "4", DesiredHash: "12"},
	}

	updateTrans = []models.UpdateTransaction{
		{
			Account: "0000000",
			Devices: []models.Device{updateDevices[0], updateDevices[1]},
		},
		{
			Account: "0000001",
			Devices: []models.Device{updateDevices[2], updateDevices[3]},
		},
	}
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	config.Init()
	config.Get().Debug = true
	db.InitDB()
	err := db.DB.AutoMigrate(
		&models.Commit{},
		&models.UpdateTransaction{},
		&models.Package{},
		&models.Image{},
		&models.Repo{},
		&models.Device{},
		&models.DispatchRecord{},
	)
	if err != nil {
		panic(err)
	}
	testImage = models.Image{
		Account: "0000000",
		Status:  models.ImageStatusBuilding,
		Commit: &models.Commit{
			Status: models.ImageStatusBuilding,
		},
	}
	db.DB.Create(&testImage.Commit)
	db.DB.Create(&testImage)
	testRepo = models.Repo{
		Commit: testImage.Commit,
	}
	db.DB.Create(&testRepo)
	db.DB.Create(&updateTrans)

}

func tearDown() {
	db.DB.Exec("DELETE FROM commits")
	db.DB.Exec("DELETE FROM repos")
	db.DB.Exec("DELETE FROM images")
	db.DB.Exec("DELETE FROM update_transactions")
}