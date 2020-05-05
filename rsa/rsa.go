package rsa

import(
	"encoding/json"
	"crypto/rand"
	"math/big"
	"log"

	"github.com/SayedAlesawy/seekur/utils/errors"
)

var logSign = "[RSA]"

// PublicKey Represents the public key
type PublicKey struct{
	N *big.Int `json:"N"`
	E *big.Int `json:"E"`
}

// PrivateKey Represents the private key
type PrivateKey struct{
	N *big.Int `json:"N"`
	D *big.Int `json:"D"`
}

// Encode A function to encode a public key
func(pubKeyObj *PublicKey) Encode() ([]byte, bool){
	encoded, err := json.Marshal(pubKeyObj)

	return encoded, errors.IsError(err)
}

// DecodePubKey A function to decode a public key
func DecodePubKey(encoded []byte) (PublicKey, bool){
	var pubKey PublicKey

	err := json.Unmarshal(encoded, &pubKey)

	return pubKey, errors.IsError(err)
}

// Encode A function to encode a private key
func(privKeyObj *PrivateKey) Encode() ([]byte, bool){
	encoded, err := json.Marshal(privKeyObj)

	return encoded, errors.IsError(err)
}

// DecodePrivKey A function to decode a private key
func DecodePrivKey(encoded []byte) (PrivateKey, bool){
	var privKey PrivateKey

	err := json.Unmarshal(encoded, &privKey)

	return privKey, errors.IsError(err)
}

// GenerateKeyPair A function to generate a public/private key pair
func GenerateKeyPair(len int) (*PublicKey, *PrivateKey, bool) {
	// Return if key length is odd
	if len % 2 != 0 {
		return nil, nil, true
	}

	// Generate p
	p, err := rand.Prime(rand.Reader, len/2)
	if err != nil {
		log.Println(logSign, "Error while generating prime number p")

		return nil, nil, true
	}

	// Generate q
	q, err := rand.Prime(rand.Reader, len/2)
	if err != nil {
		log.Println(logSign, "Error while generating prime number q")

		return nil, nil, true
	}

	// Calculate n as p*q
	n := new(big.Int).Set(p)
	n.Mul(n, q)

	// Calculate totient(n) as (p-1)*(q-1)
	p.Sub(p, big.NewInt(1))
	q.Sub(q, big.NewInt(1))
	totient := new(big.Int).Set(p)
	totient.Mul(totient, q)

	// Set e as recommended by RFC 2313
	e := big.NewInt(65537)

	// Calculate d as the modular inverse of e
	d := new(big.Int).ModInverse(e, totient)
	if d == nil {
		log.Println(logSign, "Error while calculating d")

		return nil, nil, true
	}

	publicKey := &PublicKey{N: n, E: e}
	privateKey := &PrivateKey{N: n, D: d}

	return publicKey, privateKey, false
}

// Encrypt A function to encrypt a message using the public key
func Encrypt(pubKey *PublicKey, msg []byte) big.Int {
	msgBits := new(big.Int).SetBytes(msg)

	var encrypted big.Int
	encrypted.Exp(msgBits, pubKey.E, pubKey.N)

	return encrypted
}

// Decrypt A function to encrypt a message using the private key
func Decrypt(privKey *PrivateKey, msg []byte) big.Int {
	encrypted := new(big.Int).SetBytes(msg)

	var decrypted big.Int

	decrypted.Exp(encrypted, privKey.D, privKey.N)

	return decrypted
}
