package yaml

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAesEncrypt(t *testing.T) {

	rawpassword := "guest"

	encrypt := AESEncrypt(rawpassword)
	logrus.Info("encrypt password:", encrypt)

	origin := AESDecrypt(encrypt)
	logrus.Info("decrypt password:", origin)

	if rawpassword != origin {
		t.Fail()
	}

}
