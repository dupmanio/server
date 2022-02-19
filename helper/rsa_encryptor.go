/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>, February 2022
 */

package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/dupman/server/resources"
)

type RSAEncryptor struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

const keySize = 2048

var errUnableToDecodeKey = errors.New(resources.UnableToDecodeKey)

func NewRSAEncryptor() *RSAEncryptor {
	return &RSAEncryptor{}
}

// SetPrivateKey sets the Private Key value.
func (e *RSAEncryptor) SetPrivateKey(privateKey string) (err error) {
	privateKeyDecoded, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(privateKeyDecoded)

	e.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	return nil
}

// PrivateKey gets the Private Key value.
func (e *RSAEncryptor) PrivateKey() (privateKey string) {
	privateKeyBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(e.privateKey),
		},
	)

	return base64.StdEncoding.EncodeToString(privateKeyBytes)
}

// SetPublicKey sets the Public Key value.
func (e *RSAEncryptor) SetPublicKey(publicKey string) (err error) {
	publicKeyDecoded, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(publicKeyDecoded)

	ifc, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	if key, ok := ifc.(*rsa.PublicKey); ok {
		e.publicKey = key

		return nil
	}

	return errUnableToDecodeKey
}

// PublicKey gets the Public Key value.
func (e *RSAEncryptor) PublicKey() (publicKey string, err error) {
	ASN1, err := x509.MarshalPKIXPublicKey(e.publicKey)
	if err != nil {
		return publicKey, err
	}

	publicKeyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: ASN1,
	})

	return base64.StdEncoding.EncodeToString(publicKeyBytes), err
}

// GenerateKeyPair generates a new key pair.
func (e *RSAEncryptor) GenerateKeyPair() (err error) {
	keyPair, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}

	e.privateKey = keyPair
	e.publicKey = &keyPair.PublicKey

	return nil
}

// Encrypt encrypts data with given public key.
func (e *RSAEncryptor) Encrypt(text string) (encrypted string, err error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, e.publicKey, []byte(text), nil)
	if err != nil {
		return encrypted, err
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), err
}

// Decrypt decrypts data with given private key.
func (e *RSAEncryptor) Decrypt(encrypted string) (decrypted string, err error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	decryptedBytes, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, e.privateKey, encryptedBytes, nil)
	if err != nil {
		return decrypted, err
	}

	return string(decryptedBytes), err
}
