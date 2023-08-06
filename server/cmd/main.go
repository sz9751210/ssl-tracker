// cmd/main.go

package main

import (
	"fmt"
	"net/http"
	// "github.com/alandev/ssl-monitoring/config"
	"log"

	store "github.com/alandev/ssl-monitoring/db"
	"github.com/alandev/ssl-monitoring/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // 引入 SQLite 驅動程式
)

func main() {
	store.InitDB()
	// db, err := config.InitDatabase()
	// if err != nil {
	// 	log.Fatal("無法初始化db", err)
	// }
	// defer db.Close()
	// init gin router
	router := gin.Default()
	// 配置 CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	router.GET("/certificates", handlers.HandleGetAllCertificates)
	router.GET("/certificate/", handlers.HandleGetSingleCertificate)
	router.POST("/add-certificate", handlers.HandleAddCertificate)

	// start web server
	port := 8080
	address := fmt.Sprintf(":%d", port)
	log.Printf("server running at port %s", address)
	err := router.Run(address)
	if err != nil {
		log.Fatal("Web server start failed", err)
	}
}
