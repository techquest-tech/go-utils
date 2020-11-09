package str

import (
	"os"
)

//ReplaceByEnv replace string with Env value if exists.
func ReplaceByEnv(envkey string, arg *string) {
	if envVar := os.Getenv(envkey); envVar != "" {
		*arg = envVar
		// logrus.Infof("update value with ENV.%s to %s", envkey, envVar)
	}
}
