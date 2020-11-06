package yaml

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasttemplate"
	"gopkg.in/yaml.v3"
)

//AesKey key for AES
var AesKey = "B0X128M4CO524D18"

//StartTag default start tag
var StartTag = "${"

//EndTag End tag
var EndTag = "}"

//Version version
var Version = "development"

func replaceByEnv(envkey string, arg *string) {
	if envVar := os.Getenv(envkey); envVar != "" {
		*arg = envVar
		logrus.Info("update value with ENV.YAML_AES_KEY")
	}
}

func init() {
	replaceByEnv("YAML_AES_KEY", &AesKey)
	replaceByEnv("YAML_START_TAG", &StartTag)
	replaceByEnv("YAML_END_TAG", &EndTag)
	replaceByEnv("YAML_VESION", &Version)
}

func extendYaml(w io.Writer, tag string) (int, error) {
	tags := strings.Split(tag, ".")
	prefix := strings.TrimSpace(tags[0])
	value := ""
	if len(tags) > 1 {
		value = strings.TrimSpace(tags[1])
	}

	switch prefix {
	case "env":
		envString := os.Getenv(value)
		return w.Write([]byte(envString))
	case "now":
		now := time.Now()
		strNow := now.Format(time.RFC3339)
		return w.Write([]byte(strNow))
	case "version":
		return w.Write([]byte(Version))
	case "base64":
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			logrus.Error("setting error, base64 decode failed.", err)
			return 0, err
		}
		return w.Write([]byte(decoded))
	case "aes":
		decryptCode := AesDecrypt(value, AesKey)
		return w.Write([]byte(decryptCode))
	default:
		return w.Write([]byte(tag))
	}
}

// LoadYaml load yaml file, Unmarshal it to interface{}
func LoadYaml(file string, out interface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.WithField("file", file).Error("read file failed.", err)
		return err
	}
	template, err := fasttemplate.NewTemplate(string(content), StartTag, EndTag)
	if err != nil {
		logrus.Error("valid template file failed.", err)
	}
	decoded := template.ExecuteFuncString(extendYaml)

	err = yaml.Unmarshal([]byte(decoded), out)
	if err != nil {
		logrus.Error("yaml unmarshal failed.", err)
		return err
	}

	return nil
}
