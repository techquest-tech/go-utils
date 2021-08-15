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

type RSADecrypt struct {
	Key    string
	priKey *rsa.PrivateKey
}

func (d *RSADecrypt) PostInit() error {
	raw := []byte(d.Key)
	// log.Println("lens key ", len(d.Key))
	if strings.HasPrefix(d.Key, "@") {
		file := d.Key[1:]
		log.Println("read pem file ", file)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		raw = data
	}
	block, _ := pem.Decode(raw)
	if block == nil {
		msg := fmt.Sprintf("decode private key failed. raw = %s", d.Key)
		// d.GetLogger().Error(msg)
		return fmt.Errorf(msg)
	}
	prikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// d.GetLogger().Error("parse private PEM failed.", err)
		return err
	}
	d.priKey = prikey
	log.Println("parse private PEM done.")
	return nil

}

func (d RSADecrypt) Process(raw []byte) ([]byte, error) {
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, d.priKey, raw)
	if err != nil {
		return nil, err
	}
	log.Println("decrypted by private key done. body lens = ", len(decrypted))
	return decrypted, nil
}
