package yaml

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAesEncrypt(t *testing.T) {

	rawpassword := "guest"

	encrypt := AesEncrypt(rawpassword, AesKey)
	logrus.Info("encrypt password:", encrypt)

	origin := AesDecrypt(encrypt, AesKey)
	logrus.Info("decrypt password:", origin)

	if rawpassword != origin {
		t.Fail()
	}

}
