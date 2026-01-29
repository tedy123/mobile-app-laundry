package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Laundry App Development"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Add localhost to SANs
	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"))
	template.IPAddresses = append(template.IPAddresses, net.ParseIP("10.0.2.2")) // Android Emulator access
	template.DNSNames = append(template.DNSNames, "localhost")

	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Create certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	// Save certificate to file
	certOut, err := os.Create("server.crt")
	if err != nil {
		log.Fatalf("Failed to open server.crt for writing: %v", err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		log.Fatalf("Failed to write data to server.crt: %v", err)
	}
	if err := certOut.Close(); err != nil {
		log.Fatalf("Error closing server.crt: %v", err)
	}
	log.Print("wrote server.crt\n")

	// Save private key to file
	keyOut, err := os.Create("server.key")
	if err != nil {
		log.Fatalf("Failed to open server.key for writing: %v", err)
	}
	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	if err := pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to server.key: %v", err)
	}
	if err := keyOut.Close(); err != nil {
		log.Fatalf("Error closing server.key: %v", err)
	}
	log.Print("wrote server.key\n")
}
