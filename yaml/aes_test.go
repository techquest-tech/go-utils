package yaml_test

import (
	"log"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/techquest-tech/go-utils/yaml"
)

func TestAesEncrypt(t *testing.T) {

	rawpassword := "guest"

	encrypt := yaml.AESEncrypt(rawpassword)
	log.Println("encrypt password:", encrypt)

	origin := yaml.AESDecrypt(encrypt)
	log.Println("decrypt password:", origin)

	assert.Equal(t, rawpassword, origin)

	// if rawpassword != origin {
	// 	t.Fail()
	// }

}
