package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	HeaderKey = "headerEncrypted"
)

type RSAEncrypt struct {
	// core.DataPineline
	Key    string
	pubKey *rsa.PublicKey
}

func (c *RSAEncrypt) PostInit() error {
	raw := []byte(c.Key)
	if strings.HasPrefix(c.Key, "@") {
		file := c.Key[1:]
		log.Println("read pem file ", file)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			// c.GetLogger().Error("read pub key failed. err = ", err)
			return err
		}
		raw = data
	}
	block, _ := pem.Decode(raw)
	if block == nil {
		msg := fmt.Sprintf("decode pub key failed. raw = %s", raw)
		// c.GetLogger().Error(msg)
		return fmt.Errorf(msg)
	}
	pubkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		// c.GetLogger().Error("parse pub PEM failed, ", err)
		return err
	}

	c.pubKey = pubkey.(*rsa.PublicKey)
	log.Println("parse pub PEM done.")
	return nil
}

func (c RSAEncrypt) Process(raw []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, c.pubKey, raw)
}
