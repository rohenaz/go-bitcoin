package bitcoin

import (
	"errors"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/bitcoinsv/bsvd/chaincfg"
	"github.com/bitcoinsv/bsvutil/hdkeychain"
)

const (
	// RecommendedSeedLength is the recommended length in bytes for a seed to a master node.
	RecommendedSeedLength = 32 // 256 bits

	// SecureSeedLength is the max size of a seed length (most secure
	SecureSeedLength = 64 // 512 bits
)

// GenerateHDKey will create a new master node for use in creating a hierarchical deterministic key chain
func GenerateHDKey(seedLength uint8) (hdKey *hdkeychain.ExtendedKey, err error) {

	// Missing or invalid seed length
	if seedLength == 0 {
		seedLength = RecommendedSeedLength
	}

	// Generate a new seed (added extra security from 256 to 512 bits for seed length)
	var seed []byte
	if seed, err = hdkeychain.GenerateSeed(seedLength); err != nil {
		return
	}

	// Generate a new master key
	return hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
}

// GenerateHDKeyPair will generate a new xPub HD master node (xPrivateKey & xPublicKey)
func GenerateHDKeyPair(seedLength uint8) (xPrivateKey, xPublicKey string, err error) {

	// Generate an HD master key
	var masterKey *hdkeychain.ExtendedKey
	if masterKey, err = GenerateHDKey(seedLength); err != nil {
		return
	}

	// Set the xPriv
	xPrivateKey = masterKey.String()

	// Create the extended public key
	var pubKey *hdkeychain.ExtendedKey
	if pubKey, err = masterKey.Neuter(); err != nil {
		// Error should nearly never occur since it's using a safely derived masterKey
		return
	}

	// Set the actual xPub
	xPublicKey = pubKey.String()

	return
}

// GetHDKeyByPath gets the corresponding HD key from a chain/num path
func GetHDKeyByPath(hdKey *hdkeychain.ExtendedKey, chain, num uint32) (*hdkeychain.ExtendedKey, error) {

	// Make sure we have a valid key
	if hdKey == nil {
		return nil, errors.New("hdKey is nil")
	}

	// Derive the child key from the chain path
	childKeyChain, err := hdKey.Child(chain)
	if err != nil {
		return nil, err
	}

	// Get the child key from the num path
	return childKeyChain.Child(num)
}

// GetPrivateKeyByPath gets the key for a given derivation path (chain/num)
func GetPrivateKeyByPath(hdKey *hdkeychain.ExtendedKey, chain, num uint32) (*bsvec.PrivateKey, error) {

	// Get the child key from the num & chain
	childKeyNum, err := GetHDKeyByPath(hdKey, chain, num)
	if err != nil {
		return nil, err
	}

	// Get the private key
	return childKeyNum.ECPrivKey()
}