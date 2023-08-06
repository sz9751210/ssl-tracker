package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	conn, err := tls.Dial("tcp", "www.google.com:443", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}

	err = conn.VerifyHostname("www.google.com")
	if err != nil {
		panic("Hostname doesn't match with certificate: " + err.Error())
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))
}
