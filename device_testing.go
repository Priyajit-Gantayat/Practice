package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup mock database
func setupTestDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=yourpassword dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	// Migrate the schema
	db.AutoMigrate(&Device{})
	// Clear existing data
	db.Exec("TRUNCATE TABLE devices RESTART IDENTITY CASCADE")
	return db
}

// Test helper to create a test router
func setupTestRouter() *gin.Engine {
	db = setupTestDB()
	r := setupRouter()
	return r
}

// Test registering a new device
func TestRegisterDevice(t *testing.T) {
	r := setupTestRouter()

	payload := Device{
		DeviceName:   "Test Device",
		DeviceType:   "Mobile",
		Brand:        "TestBrand",
		Model:        "ModelX",
		Os:           "Android",
		OsVersion:    "11",
		PurchaseDate: "2023-01-01",
		WarrantyEnd:  "2025-01-01",
		Status:       "Active",
		Price:        500,
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/device", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	response := Device{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, payload.DeviceName, response.DeviceName)
}

// Test listing devices
func TestListDevices(t *testing.T) {
	r := setupTestRouter()

	// Seed data
	db.Create(&Device{
		DeviceName: "Device1",
		DeviceType: "Mobile",
		Brand:      "Brand1",
		Status:     "Active",
		Price:      200,
	})
	db.Create(&Device{
		DeviceName: "Device2",
		DeviceType: "Laptop",
		Brand:      "Brand2",
		Status:     "Inactive",
		Price:      1000,
	})

	req, _ := http.NewRequest("GET", "/device", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var devices []Device
	_ = json.Unmarshal(w.Body.Bytes(), &devices)
	assert.Len(t, devices, 2)
}

// Test getting a device by ID
func TestGetDeviceByID(t *testing.T) {
	r := setupTestRouter()

	device := Device{
		DeviceName: "Device1",
		DeviceType: "Mobile",
		Brand:      "Brand1",
		Status:     "Active",
		Price:      200,
	}
	db.Create(&device)

	req, _ := http.NewRequest("GET", "/device/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	response := Device{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, device.DeviceName, response.DeviceName)
}

// Test updating a device
func TestUpdateDevice(t *testing.T) {
	r := setupTestRouter()

	device := Device{
		DeviceName: "Device1",
		DeviceType: "Mobile",
		Brand:      "Brand1",
		Status:     "Active",
		Price:      200,
	}
	db.Create(&device)

	updatePayload := map[string]interface{}{
		"device_name": "Updated Device",
		"price":       300,
	}
	jsonPayload, _ := json.Marshal(updatePayload)

	req, _ := http.NewRequest("PUT", "/device/1", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var updated Device
	db.First(&updated, 1)
	assert.Equal(t, "Updated Device", updated.DeviceName)
	assert.Equal(t, uint(300), updated.Price)
}

// Test deleting a device
func TestDeleteDevice(t *testing.T) {
	r := setupTestRouter()

	device := Device{
		DeviceName: "Device1",
		DeviceType: "Mobile",
		Brand:      "Brand1",
		Status:     "Active",
		Price:      200,
	}
	db.Create(&device)

	req, _ := http.NewRequest("DELETE", "/device/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var deleted Device
	result := db.First(&deleted, 1)
	assert.Error(t, result.Error) // Record should not exist
}

// Test CSV upload
func TestUploadCSV(t *testing.T) {
	r := setupTestRouter()

	csvData := `Device1,Mobile,Brand1,Model1,Android,11,2023-01-01,2025-01-01,Active,500
Device2,Laptop,Brand2,Model2,Windows,10,2022-01-01,2024-01-01,Inactive,1000`
	reqBody := strings.NewReader(csvData)

	req, _ := http.NewRequest("POST", "/upload", reqBody)
	req.Header.Set("Content-Type", "text/csv")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var devices []Device
	db.Find(&devices)
	assert.Len(t, devices, 2)
	assert.Equal(t, "Device1", devices[0].DeviceName)
	assert.Equal(t, "Device2", devices[1].DeviceName)
}
