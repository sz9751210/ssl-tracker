package store

import (
	"log"
	"time"

	// "github.com/alandev/ssl-monitoring/config"
	"github.com/alandev/ssl-monitoring/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("certificates.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	db.AutoMigrate(&types.Certificate{})
}

func GetCertificates() []types.Certificate {
	var certs []types.Certificate
	db.Find(&certs)
	for i := range certs {
		setStatus(&certs[i])
	}

	return certs
}

func GetCertificateByDomain(domain string) (types.Certificate, error) {
	var cert types.Certificate
	err := db.Where("domain = ?", domain).First(&cert).Error
	setStatus(&cert)
	return cert, err
}

func AddCertificate(cert types.Certificate) error {
	return db.Create(&cert).Error
}

func setStatus(cert *types.Certificate) {
	daysUntilExp := cert.ExpirationDate.Sub(time.Now()).Hours() / 24

	if daysUntilExp <= 0 {
		cert.Status = 0 // 已到期，紅色或數字0
	} else if daysUntilExp <= 7 {
		cert.Status = 1 // 一週內到期，黃色或數字1
	} else if daysUntilExp <= 30 {
		cert.Status = 2 // 一個月內到期，藍色或數字2
	} else {
		cert.Status = 3 // 一個月以上到期，綠色或數字3
	}
}
