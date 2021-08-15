package rsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	privateFile = "../private.pem"
	pubFile     = "../public.pem"
)

func TestGenerator(t *testing.T) {
	err := Generate(2048, privateFile, pubFile)
	assert.Nil(t, err)
}

func TestRSA(t *testing.T) {

	encrypted := &RSAEncrypt{
		Key: "@" + pubFile,
		// Headers: true,
	}

	err := encrypted.PostInit()
	assert.Nil(t, err)

	decrypted := &RSADecrypt{
		Key: "@" + privateFile,
	}
	err = decrypted.PostInit()
	assert.Nil(t, err)

	expected := ("14p5l8zxs7hajbou0eawuiyhogvogoccu546b5077zuidwmk3nlddg9qwr8adid6")

	out, err := encrypted.Process([]byte(expected))
	assert.Nil(t, err)
	assert.NotNil(t, out)

	d, err := decrypted.Process(out)
	assert.Nil(t, err)
	assert.NotNil(t, d)

	result := string(d)
	assert.Equal(t, expected, result)
}
