package main

import (
	"fmt"

	"autobros_baru/connection"
	"autobros_baru/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := connection.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	r := gin.Default()

	kendaraanHandler := models.KendaraanHandler{
		DB: db,
	}

	r.POST("/autobros/tambahservis", kendaraanHandler.CreateKendaraan)
	r.GET("/autobros/detailservis/:id", kendaraanHandler.GetKendaraan)
	r.GET("/autobros/tabelservis", kendaraanHandler.GetAllKendaraan)
	r.PUT("/autobros/updateservis/:id", kendaraanHandler.UpdateKendaraan)
	r.DELETE("/autobros/hapusservis/:id", kendaraanHandler.DeleteKendaraan)
	r.LoadHTMLGlob("autobros/web")
	r.Static("/autobros/dashboard", "./web")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Autobros"})
	})
	r.Run(":8000")
}
