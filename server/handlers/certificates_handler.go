// handlers/certificates.go

package handlers

import (
	"net/http"
	"time"

	"github.com/alandev/ssl-monitoring/db"
	"github.com/alandev/ssl-monitoring/types"
	"github.com/alandev/ssl-monitoring/utils"
	"github.com/gin-gonic/gin"
)

var certificates []types.Certificate

func HandleGetAllCertificates(c *gin.Context) {
	certificates := store.GetCertificates()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, certificates)
}

func HandleGetSingleCertificate(c *gin.Context) {
	domain := c.Query("domain")
	cert, err := store.GetCertificateByDomain(domain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, cert)
}

func HandleAddCertificate(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
		return
	}

	var requestData struct {
		Domain string `json:"domain"`
	}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	expirationDate, err := utils.GetCertificateExpiration(requestData.Domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving certificate expiration date"})
		return
	}

	daysUntilExpiration := int(expirationDate.Sub(time.Now()).Hours() / 24)

	err = store.AddCertificate(types.Certificate{
		Domain:         requestData.Domain,
		ExpirationDate: expirationDate,
		DaysUntilExp:   daysUntilExpiration,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Certificate added successfully"})
}
