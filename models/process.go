package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Kendaraan struct {
	ID             int       `json:"id"`
	JenisID        int       `json:"id_jenis"`
	NamaPemilik    string    `json:"nama_pemilik"`
	NamaKendaraan  string    `json:"nama_kendaraan"`
	NomorKendaraan int32     `json:"nomor_kendaraan"`
	DetailServis   string    `json:"detail_servis"`
	TanggalServis  time.Time `json:"tanggal_servis"`
	StatusServis   string    `json:"status_servis"`
}

type KendaraanHandler struct {
	DB *sql.DB
}

func (h *KendaraanHandler) CreateKendaraan(c *gin.Context) {
	var kendaraan Kendaraan
	err := c.ShouldBindJSON(&kendaraan)
	if err != nil {
		fmt.Println("Received JSON payload:", c.Request.Body)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate Tanggal Sekarang
	kendaraan.TanggalServis = time.Now()

	// Insert Data Ke Database
	stmt := ` INSERT INTO data_servis_kendaraan (id_jenis, nama_pemilik, nama_kendaraan, nomor_kendaraan, detail_servis, tanggal_servis, status_servis) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
	`
	var id int
	err = h.DB.QueryRow(stmt, kendaraan.JenisID, kendaraan.NamaPemilik, kendaraan.NamaKendaraan, int32(kendaraan.NomorKendaraan), kendaraan.DetailServis, kendaraan.TanggalServis, kendaraan.StatusServis).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{
		"message": "Data telah dibuat!",
		"id":      id,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *KendaraanHandler) GetKendaraan(c *gin.Context) {
	id := c.Param("id")
	kendaraanID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Ambil Satu Data Servis + Ambil Jenis kendaraan berdasarkan Relasi dari id_jenis
	stmt := `
		SELECT data_servis_kendaraan.id_jenis, nama_kendaraan, jenis_kendaraan.jenis_kendaraan nama_pemilik, nomor_kendaraan, detail_servis, tanggal_servis, status_servis
		FROM data_servis_kendaraan INNER JOIN jenis_kendaraan ON jenis_kendaraan.id_jenis = data_servis_kendaraan.id_jenis
		WHERE id = $1
	`
	row := h.DB.QueryRow(stmt, kendaraanID)

	var kendaraan Kendaraan
	err = row.Scan(&kendaraan.JenisID, &kendaraan.NamaPemilik, &kendaraan.NamaKendaraan, &kendaraan.NomorKendaraan, &kendaraan.DetailServis, &kendaraan.TanggalServis, &kendaraan.StatusServis)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kendaraan not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, kendaraan)
}

func (h *KendaraanHandler) GetAllKendaraan(c *gin.Context) {
	stmt := `
		SELECT id, id_jenis, nama_pemilik, nama_kendaraan, nomor_kendaraan, detail_servis, tanggal_servis, status_servis
		FROM data_servis_kendaraan
	`
	rows, err := h.DB.Query(stmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var kendaraans []Kendaraan
	for rows.Next() {
		var kendaraan Kendaraan
		err := rows.Scan(&kendaraan.ID, &kendaraan.JenisID, &kendaraan.NamaPemilik, &kendaraan.NamaKendaraan, &kendaraan.NomorKendaraan, &kendaraan.DetailServis, &kendaraan.TanggalServis, &kendaraan.StatusServis)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		kendaraans = append(kendaraans, kendaraan)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kendaraans)
}

func (h *KendaraanHandler) UpdateKendaraan(c *gin.Context) {
	id := c.Param("id")
	kendaraanID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var kendaraan Kendaraan
	err = c.ShouldBindJSON(&kendaraan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Satu Data Servis
	stmt := `
		UPDATE data_servis_kendaraan
		SET id_jenis = $1, nama_pemilik = $2 nama_kendaraan = $3, nomor_kendaraan = $4, detail_servis = $5, tanggal_servis = $6, status_servis = $7
		WHERE id = $8
	`
	_, err = h.DB.Exec(stmt, kendaraan.JenisID, kendaraan.NamaPemilik, kendaraan.NamaKendaraan, kendaraan.NomorKendaraan, kendaraan.DetailServis, kendaraan.TanggalServis, kendaraan.StatusServis, kendaraanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"message": "Data telah diperbarui!",
	}
	c.JSON(http.StatusOK, response)
}

func (h *KendaraanHandler) DeleteKendaraan(c *gin.Context) {
	id := c.Param("id")
	kendaraanID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hapus Satu Data Servis
	stmt := `
		DELETE FROM data_servis_kendaraan
		WHERE id = $1
	`
	_, err = h.DB.Exec(stmt, kendaraanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Telah Di hapus!"})
}
