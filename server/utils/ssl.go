// utils/ssl.go

package utils

import (
	"crypto/tls"
	"fmt"
	"time"
)

func GetCertificateExpiration(domain string) (time.Time, error) {
	fmt.Println("domain", domain)
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}
	defer conn.Close()

	err = conn.VerifyHostname(domain)
	if err != nil {
		panic("Hostname doesn't match with certificate: " + err.Error())
	}
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return time.Time{}, fmt.Errorf("no certificates found for %s", domain)
	}

	return certs[0].NotAfter, nil
}
