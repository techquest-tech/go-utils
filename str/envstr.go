package str

import (
	"os"

	"github.com/sirupsen/logrus"
)

//ReplaceByEnv replace string with Env value if exists.
func ReplaceByEnv(envkey string, arg *string) {
	if envVar := os.Getenv(envkey); envVar != "" {
		*arg = envVar
		logrus.Info("update value with ENV.", envkey)
	}
}
