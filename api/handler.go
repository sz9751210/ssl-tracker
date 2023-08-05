package api

import (
	"net/http"
	"time"

	"github.com/alandev/go-ssl-tracker/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouters(router *gin.Engine, db *gorm.DB) {
	router.POST("/add-certificate", func(c *gin.Context) {
		var req struct {
			Domain     string `json:"domain" binding:"required"`
			Expiration string `json:"expiration" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		expiration, err := time.Parse("2006-01-02", req.Expiration)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "無效的時間格式"})
			return
		}

		certificate := types.Certificate{
			Domain:     req.Domain,
			Expiration: expiration,
		}
		if err := db.Create(&certificate).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "無法添加憑證"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "憑證添加成功"})
	})
	router.GET("/certificates", func(c *gin.Context) {
		var certificates []types.Certificate
		if err := db.Find(&certificates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "無法獲取憑證"})
		}
		c.JSON(http.StatusOK, certificates)
	})
}
