package main

import (
	"log"

	"github.com/bitcoinschema/go-bitcoin"
)

func main() {

	// Start with a private key  (we will make one for this example)
	privateKey, err := bitcoin.CreatePrivateKeyString()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Create a wif
	var wif string
	if wif, err = bitcoin.PrivateKeyToWifString(privateKey); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("private key: %s converted to wif: %s", privateKey, wif)
}
